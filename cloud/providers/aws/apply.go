package aws

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"
	// "github.com/appscode/pharmer/templates"
	proto "github.com/appscode/api/kubernetes/v1beta1"
	"github.com/appscode/go/errors"
	. "github.com/appscode/go/types"
	"github.com/appscode/pharmer/api"
	"github.com/appscode/pharmer/cloud"
	"github.com/appscode/pharmer/phid"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "github.com/appscode/pharmer/templates"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	_ec2 "github.com/aws/aws-sdk-go/service/ec2"
	_iam "github.com/aws/aws-sdk-go/service/iam"
)

const (
	preTagDelay = 5 * time.Second
)

func (cm *ClusterManager) Apply(cluster string, dryRun bool) error {
	var err error

	if cm.cluster, err = cloud.Store(cm.ctx).Clusters().Get(cluster); err != nil {
		return err
	}
	if cm.conn, err = NewConnector(cm.ctx, cm.cluster); err != nil {
		return err
	}

	defer func(releaseReservedIp bool) {
		if cm.cluster.Status.Phase == api.ClusterPhasePending {
			cm.cluster.Status.Phase = api.ClusterPhaseFailing
		}
		cloud.Store(cm.ctx).Clusters().UpdateStatus(cm.cluster)
		cloud.Logger(cm.ctx).Infof("Cluster %v is %v", cm.cluster.Name, cm.cluster.Status.Phase)
		if cm.cluster.Status.Phase != api.ClusterPhaseReady {
			cloud.Logger(cm.ctx).Infof("Cluster %v is deleting", cm.cluster.Name)
			cm.Delete(&proto.ClusterDeleteRequest{
				Name:              cm.cluster.Name,
				ReleaseReservedIp: releaseReservedIp,
			})
		}
	}(cm.cluster.Spec.MasterReservedIP == "auto")

	if err = cm.conn.detectUbuntuImage(); err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cm.cluster.Spec.Cloud.InstanceImage = cm.conn.cluster.Spec.Cloud.InstanceImage
	// TODO: FixIt!
	//cm.cluster.Spec.RootDeviceName = cm.conn.cluster.Spec.RootDeviceName
	//fmt.Println(cm.cluster.Spec.Cloud.InstanceImage, cm.cluster.Spec.RootDeviceName, "---------------*********")

	if err = cm.ensureIAMProfile(); err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	if err = cm.importPublicKey(); err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	if err = cm.setupVpc(); err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	if err = cm.createDHCPOptionSet(); err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	if err = cm.setupSubnet(); err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	if err = cm.setupInternetGateway(); err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	if err = cm.setupRouteTable(); err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	if err = cm.setupSecurityGroups(); err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	masterInstance, err := cm.startMaster()
	if err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	//for _, ng := range req.NodeGroups {
	//	igm := &InstanceGroupManager{
	//		cm: cm,
	//		instance: cloud.Instance{
	//			Type: cloud.InstanceType{
	//				ContextVersion: cm.cluster.Generation,
	//				Sku:            ng.Sku,
	//
	//				Master:       false,
	//				SpotInstance: false,
	//			},
	//			Stats: cloud.GroupStats{
	//				Count: ng.Count,
	//			},
	//		},
	//	}
	//	igm.AdjustInstanceGroup()
	//}

	cloud.Logger(cm.ctx).Info("Waiting for cluster initialization")

	// Wait for master A record to propagate
	if err := cloud.EnsureDnsIPLookup(cm.ctx, cm.cluster); err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	if err := cloud.EnsureDnsIPLookup(cm.ctx, cm.cluster); err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	// wait for nodes to start
	if err := cloud.WaitForReadyMaster(cm.ctx, cm.cluster); err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	// -------------------------------------------------------------------------------------------------------------
	cloud.Logger(cm.ctx).Info("Listing autoscaling groups")
	groups := make([]*string, 0)
	//for _, ng := range req.NodeGroups {
	//	groups = append(groups, StringP(cm.namer.AutoScalingGroupName(ng.Sku)))
	//}
	r2, err := cm.conn.autoscale.DescribeAutoScalingGroups(&autoscaling.DescribeAutoScalingGroupsInput{
		AutoScalingGroupNames: groups,
	})
	if err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	fmt.Println(r2)
	cloud.Store(cm.ctx).Instances(cm.cluster.Name).Create(masterInstance)
	for _, group := range r2.AutoScalingGroups {
		for _, instance := range group.Instances {
			ki, err := cm.newKubeInstance(*instance.InstanceId)
			ki.Spec.Role = api.RoleKubernetesPool
			cloud.Store(cm.ctx).Instances(cm.cluster.Name).Create(ki)
			if err != nil {
				return errors.FromErr(err).WithContext(cm.ctx).Err()
			}
		}
	}

	// detect-master
	// wait-master: via curl call polling
	// build-config

	//  # KUBE_SHARE_MASTER is used to add nodes to an existing master
	//  if [[ "${KUBE_SHARE_MASTER:-}" == "true" ]]; then
	//    detect-master
	//    start-nodes
	//    wait-nodes
	//  else
	//    start-master
	//    start-nodes
	//    wait-nodes
	//    wait-master
	//
	//    # Build ~/.kube/config
	//    build-config
	//  fi
	// check-cluster
	cm.cluster.Status.Phase = api.ClusterPhaseReady
	return nil
}

func (cm *ClusterManager) ensureIAMProfile() error {
	r1, _ := cm.conn.iam.GetInstanceProfile(&_iam.GetInstanceProfileInput{InstanceProfileName: &cm.cluster.Spec.Cloud.AWS.IAMProfileMaster})
	if r1.InstanceProfile == nil {
		err := cm.createIAMProfile(cm.cluster.Spec.Cloud.AWS.IAMProfileMaster)
		if err != nil {
			return errors.FromErr(err).WithContext(cm.ctx).Err()
		}
		cloud.Logger(cm.ctx).Infof("Master instance profile %v created", cm.cluster.Spec.Cloud.AWS.IAMProfileMaster)
	}
	r2, _ := cm.conn.iam.GetInstanceProfile(&_iam.GetInstanceProfileInput{InstanceProfileName: &cm.cluster.Spec.Cloud.AWS.IAMProfileNode})
	if r2.InstanceProfile == nil {
		err := cm.createIAMProfile(cm.cluster.Spec.Cloud.AWS.IAMProfileNode)
		if err != nil {
			return errors.FromErr(err).WithContext(cm.ctx).Err()
		}
		cloud.Logger(cm.ctx).Infof("Node instance profile %v created", cm.cluster.Spec.Cloud.AWS.IAMProfileNode)
	}
	return nil
}

func (cm *ClusterManager) createIAMProfile(key string) error {
	//rootDir := "kubernetes/aws/iam/"
	role := "" // TODO(tamal); FixIt!  templates.AssetText(rootDir + key + "-role.json")
	r1, err := cm.conn.iam.CreateRole(&_iam.CreateRoleInput{
		RoleName:                 &key,
		AssumeRolePolicyDocument: &role,
	})
	cloud.Logger(cm.ctx).Debug("Created IAM role", r1, err)
	cloud.Logger(cm.ctx).Infof("IAM role %v created", key)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	policy := "" // TODO(tamal); FixIt!  templates.AssetText(rootDir + key + "-policy.json")
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	r2, err := cm.conn.iam.PutRolePolicy(&_iam.PutRolePolicyInput{
		RoleName:       &key,
		PolicyName:     &key,
		PolicyDocument: &policy,
	})
	cloud.Logger(cm.ctx).Debug("Created IAM role-policy", r2, err)
	cloud.Logger(cm.ctx).Infof("IAM role-policy %v created", key)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	r3, err := cm.conn.iam.CreateInstanceProfile(&_iam.CreateInstanceProfileInput{
		InstanceProfileName: &key,
	})
	cloud.Logger(cm.ctx).Debug("Created IAM instance-policy", r3, err)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cloud.Logger(cm.ctx).Infof("IAM instance-policy %v created", key)

	r4, err := cm.conn.iam.AddRoleToInstanceProfile(&_iam.AddRoleToInstanceProfileInput{
		InstanceProfileName: &key,
		RoleName:            &key,
	})
	cloud.Logger(cm.ctx).Debug("Added IAM role to instance-policy", r4, err)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cloud.Logger(cm.ctx).Infof("IAM role %v added to instance-policy %v", key, key)
	return nil
}

func (cm *ClusterManager) importPublicKey() error {
	resp, err := cm.conn.ec2.ImportKeyPair(&_ec2.ImportKeyPairInput{
		KeyName:           StringP(cm.cluster.Status.SSHKeyExternalID),
		PublicKeyMaterial: cloud.SSHKey(cm.ctx).PublicKey,
	})
	cloud.Logger(cm.ctx).Debug("Imported SSH key", resp, err)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	// TODO ignore "InvalidKeyPair.Duplicate" error
	if err != nil {
		cloud.Logger(cm.ctx).Info("Error importing public key", resp, err)
		//os.Exit(1)
		return errors.FromErr(err).WithContext(cm.ctx).Err()

	}
	cloud.Logger(cm.ctx).Infof("SSH key with (AWS) fingerprint %v imported", cloud.SSHKey(cm.ctx).AwsFingerprint)

	return nil
}

func (cm *ClusterManager) setupVpc() error {
	cloud.Logger(cm.ctx).Infof("Checking VPC tagged with %v", cm.cluster.Name)
	r1, err := cm.conn.ec2.DescribeVpcs(&_ec2.DescribeVpcsInput{
		Filters: []*_ec2.Filter{
			{
				Name: StringP("tag:Name"),
				Values: []*string{
					StringP(cm.namer.VPCName()),
				},
			},
			{
				Name: StringP("tag:KubernetesCluster"),
				Values: []*string{
					StringP(cm.cluster.Name), // Tag by Name or PHID?
				},
			},
		},
	})
	cloud.Logger(cm.ctx).Debug("VPC described", r1, err)
	if len(r1.Vpcs) > 1 {
		cm.cluster.Status.Cloud.AWS.VpcId = *r1.Vpcs[0].VpcId
		cloud.Logger(cm.ctx).Infof("VPC %v found", cm.cluster.Status.Cloud.AWS.VpcId)
	}

	cloud.Logger(cm.ctx).Info("No VPC found, creating new VPC")
	r2, err := cm.conn.ec2.CreateVpc(&_ec2.CreateVpcInput{
		CidrBlock: StringP(cm.cluster.Spec.Cloud.AWS.VpcCIDR),
	})
	cloud.Logger(cm.ctx).Debug("VPC created", r2, err)
	//errorutil.EOE(err)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cloud.Logger(cm.ctx).Infof("VPC %v created", *r2.Vpc.VpcId)
	cm.cluster.Status.Cloud.AWS.VpcId = *r2.Vpc.VpcId

	r3, err := cm.conn.ec2.ModifyVpcAttribute(&_ec2.ModifyVpcAttributeInput{
		VpcId: StringP(cm.cluster.Status.Cloud.AWS.VpcId),
		EnableDnsSupport: &_ec2.AttributeBooleanValue{
			Value: TrueP(),
		},
	})
	cloud.Logger(cm.ctx).Debug("DNS support enabled", r3, err)
	cloud.Logger(cm.ctx).Infof("Enabled DNS support for VPCID %v", cm.cluster.Status.Cloud.AWS.VpcId)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	r4, err := cm.conn.ec2.ModifyVpcAttribute(&_ec2.ModifyVpcAttributeInput{
		VpcId: StringP(cm.cluster.Status.Cloud.AWS.VpcId),
		EnableDnsHostnames: &_ec2.AttributeBooleanValue{
			Value: TrueP(),
		},
	})
	cloud.Logger(cm.ctx).Debug("DNS hostnames enabled", r4, err)
	cloud.Logger(cm.ctx).Infof("Enabled DNS hostnames for VPCID %v", cm.cluster.Status.Cloud.AWS.VpcId)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	time.Sleep(preTagDelay)
	cm.addTag(cm.cluster.Status.Cloud.AWS.VpcId, "Name", cm.namer.VPCName())
	cm.addTag(cm.cluster.Status.Cloud.AWS.VpcId, "KubernetesCluster", cm.cluster.Name)
	return nil
}

func (cm *ClusterManager) addTag(id string, key string, value string) error {
	resp, err := cm.conn.ec2.CreateTags(&_ec2.CreateTagsInput{
		Resources: []*string{
			StringP(id),
		},
		Tags: []*_ec2.Tag{
			{
				Key:   StringP(key),
				Value: StringP(value),
			},
		},
	})
	cloud.Logger(cm.ctx).Debug("Added tag ", resp, err)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cloud.Logger(cm.ctx).Infof("Added tag %v:%v to id %v", key, value, id)
	return nil
}

func (cm *ClusterManager) createDHCPOptionSet() error {
	optionSetDomain := fmt.Sprintf("%v.compute.internal", cm.cluster.Spec.Cloud.Region)
	if cm.cluster.Spec.Cloud.Region == "us-east-1" {
		optionSetDomain = "ec2.internal"
	}
	r1, err := cm.conn.ec2.CreateDhcpOptions(&_ec2.CreateDhcpOptionsInput{
		DhcpConfigurations: []*_ec2.NewDhcpConfiguration{
			{
				Key:    StringP("domain-name"),
				Values: []*string{StringP(optionSetDomain)},
			},
			{
				Key:    StringP("domain-name-servers"),
				Values: []*string{StringP("AmazonProvidedDNS")},
			},
		},
	})
	cloud.Logger(cm.ctx).Debug("Created DHCP options ", r1, err)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cloud.Logger(cm.ctx).Infof("DHCP options created with id %v", *r1.DhcpOptions.DhcpOptionsId)
	cm.cluster.Status.Cloud.AWS.DHCPOptionsId = *r1.DhcpOptions.DhcpOptionsId

	time.Sleep(preTagDelay)
	cm.addTag(cm.cluster.Status.Cloud.AWS.DHCPOptionsId, "Name", cm.namer.DHCPOptionsName())
	cm.addTag(cm.cluster.Status.Cloud.AWS.DHCPOptionsId, "KubernetesCluster", cm.cluster.Name)

	r2, err := cm.conn.ec2.AssociateDhcpOptions(&_ec2.AssociateDhcpOptionsInput{
		DhcpOptionsId: StringP(cm.cluster.Status.Cloud.AWS.DHCPOptionsId),
		VpcId:         StringP(cm.cluster.Status.Cloud.AWS.VpcId),
	})
	cloud.Logger(cm.ctx).Debug("Associated DHCP options ", r2, err)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cloud.Logger(cm.ctx).Infof("DHCP options %v associated with %v", cm.cluster.Status.Cloud.AWS.DHCPOptionsId, cm.cluster.Status.Cloud.AWS.VpcId)

	return nil
}

func (cm *ClusterManager) setupSubnet() error {
	cloud.Logger(cm.ctx).Info("Checking for existing subnet")
	r1, err := cm.conn.ec2.DescribeSubnets(&_ec2.DescribeSubnetsInput{
		Filters: []*_ec2.Filter{
			{
				Name: StringP("tag:KubernetesCluster"),
				Values: []*string{
					StringP(cm.cluster.Name),
				},
			},
			{
				Name: StringP("availabilityZone"),
				Values: []*string{
					StringP(cm.cluster.Spec.Cloud.Zone),
				},
			},
			{
				Name: StringP("vpc-id"),
				Values: []*string{
					StringP(cm.cluster.Status.Cloud.AWS.VpcId),
				},
			},
		},
	})
	cloud.Logger(cm.ctx).Debug("Retrieved subnet", r1, err)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	if len(r1.Subnets) == 0 {
		cloud.Logger(cm.ctx).Info("No subnet found, creating new subnet")
		r2, err := cm.conn.ec2.CreateSubnet(&_ec2.CreateSubnetInput{
			CidrBlock:        StringP(cm.cluster.Spec.Cloud.AWS.SubnetCIDR),
			VpcId:            StringP(cm.cluster.Status.Cloud.AWS.VpcId),
			AvailabilityZone: StringP(cm.cluster.Spec.Cloud.Zone),
		})
		cloud.Logger(cm.ctx).Debug("Created subnet", r2, err)
		if err != nil {
			return errors.FromErr(err).WithContext(cm.ctx).Err()
		}
		cloud.Logger(cm.ctx).Infof("Subnet %v created", *r2.Subnet.SubnetId)
		cm.cluster.Status.Cloud.AWS.SubnetId = *r2.Subnet.SubnetId

		time.Sleep(preTagDelay)
		cm.addTag(cm.cluster.Status.Cloud.AWS.SubnetId, "KubernetesCluster", cm.cluster.Name)

	} else {
		cm.cluster.Status.Cloud.AWS.SubnetId = *r1.Subnets[0].SubnetId
		existingCIDR := *r1.Subnets[0].CidrBlock
		cloud.Logger(cm.ctx).Infof("Subnet %v found with CIDR %v", cm.cluster.Status.Cloud.AWS.SubnetId, existingCIDR)

		cloud.Logger(cm.ctx).Infof("Retrieving VPC %v", cm.cluster.Status.Cloud.AWS.VpcId)
		r3, err := cm.conn.ec2.DescribeVpcs(&_ec2.DescribeVpcsInput{
			VpcIds: []*string{StringP(cm.cluster.Status.Cloud.AWS.VpcId)},
		})
		cloud.Logger(cm.ctx).Debug("Retrieved VPC", r3, err)
		if err != nil {
			return errors.FromErr(err).WithContext(cm.ctx).Err()
		}

		octets := strings.Split(*r3.Vpcs[0].CidrBlock, ".")
		cm.cluster.Spec.Cloud.AWS.VpcCIDRBase = octets[0] + "." + octets[1]
		cm.cluster.Spec.MasterInternalIP = cm.cluster.Spec.Cloud.AWS.VpcCIDRBase + ".0" + cm.cluster.Spec.Cloud.AWS.MasterIPSuffix
		cloud.Logger(cm.ctx).Infof("Assuming MASTER_INTERNAL_IP=%v", cm.cluster.Spec.MasterInternalIP)
	}
	return nil
}

func (cm *ClusterManager) setupInternetGateway() error {
	cloud.Logger(cm.ctx).Infof("Checking IGW with attached VPCID %v", cm.cluster.Status.Cloud.AWS.VpcId)
	r1, err := cm.conn.ec2.DescribeInternetGateways(&_ec2.DescribeInternetGatewaysInput{
		Filters: []*_ec2.Filter{
			{
				Name: StringP("attachment.vpc-id"),
				Values: []*string{
					StringP(cm.cluster.Status.Cloud.AWS.VpcId),
				},
			},
		},
	})
	cloud.Logger(cm.ctx).Debug("Retrieved IGW", r1, err)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	if len(r1.InternetGateways) == 0 {
		cloud.Logger(cm.ctx).Info("No IGW found, creating new IGW")
		r2, err := cm.conn.ec2.CreateInternetGateway(&_ec2.CreateInternetGatewayInput{})
		cloud.Logger(cm.ctx).Debug("Created IGW", r2, err)
		if err != nil {
			return errors.FromErr(err).WithContext(cm.ctx).Err()
		}
		cm.cluster.Status.Cloud.AWS.IGWId = *r2.InternetGateway.InternetGatewayId
		time.Sleep(preTagDelay)
		cloud.Logger(cm.ctx).Infof("IGW %v created", cm.cluster.Status.Cloud.AWS.IGWId)

		r3, err := cm.conn.ec2.AttachInternetGateway(&_ec2.AttachInternetGatewayInput{
			InternetGatewayId: StringP(cm.cluster.Status.Cloud.AWS.IGWId),
			VpcId:             StringP(cm.cluster.Status.Cloud.AWS.VpcId),
		})
		cloud.Logger(cm.ctx).Debug("Attached IGW to VPC", r3, err)
		if err != nil {
			return errors.FromErr(err).WithContext(cm.ctx).Err()
		}
		cloud.Logger(cm.ctx).Infof("Attached IGW %v to VPCID %v", cm.cluster.Status.Cloud.AWS.IGWId, cm.cluster.Status.Cloud.AWS.VpcId)

		cm.addTag(cm.cluster.Status.Cloud.AWS.IGWId, "Name", cm.namer.InternetGatewayName())
		cm.addTag(cm.cluster.Status.Cloud.AWS.IGWId, "KubernetesCluster", cm.cluster.Name)
	} else {
		cm.cluster.Status.Cloud.AWS.IGWId = *r1.InternetGateways[0].InternetGatewayId
		cloud.Logger(cm.ctx).Infof("IGW %v found", cm.cluster.Status.Cloud.AWS.IGWId)
	}
	return nil
}

func (cm *ClusterManager) setupRouteTable() error {
	cloud.Logger(cm.ctx).Infof("Checking route table for VPCID %v", cm.cluster.Status.Cloud.AWS.VpcId)
	r1, err := cm.conn.ec2.DescribeRouteTables(&_ec2.DescribeRouteTablesInput{
		Filters: []*_ec2.Filter{
			{
				Name: StringP("vpc-id"),
				Values: []*string{
					StringP(cm.cluster.Status.Cloud.AWS.VpcId),
				},
			},
			{
				Name: StringP("tag:KubernetesCluster"),
				Values: []*string{
					StringP(cm.cluster.Name),
				},
			},
		},
	})
	cloud.Logger(cm.ctx).Debug("Attached IGW to VPC", r1, err)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	if len(r1.RouteTables) == 0 {
		cloud.Logger(cm.ctx).Infof("No route table found for VPCID %v, creating new route table", cm.cluster.Status.Cloud.AWS.VpcId)
		r2, err := cm.conn.ec2.CreateRouteTable(&_ec2.CreateRouteTableInput{
			VpcId: StringP(cm.cluster.Status.Cloud.AWS.VpcId),
		})
		cloud.Logger(cm.ctx).Debug("Created route table", r2, err)
		if err != nil {
			return errors.FromErr(err).WithContext(cm.ctx).Err()
		}

		cm.cluster.Status.Cloud.AWS.RouteTableId = *r2.RouteTable.RouteTableId
		cloud.Logger(cm.ctx).Infof("Route table %v created", cm.cluster.Status.Cloud.AWS.RouteTableId)
		time.Sleep(preTagDelay)
		cm.addTag(cm.cluster.Status.Cloud.AWS.RouteTableId, "KubernetesCluster", cm.cluster.Name)

	} else {
		cm.cluster.Status.Cloud.AWS.RouteTableId = *r1.RouteTables[0].RouteTableId
		cloud.Logger(cm.ctx).Infof("Route table %v found", cm.cluster.Status.Cloud.AWS.RouteTableId)
	}

	r3, err := cm.conn.ec2.AssociateRouteTable(&_ec2.AssociateRouteTableInput{
		RouteTableId: StringP(cm.cluster.Status.Cloud.AWS.RouteTableId),
		SubnetId:     StringP(cm.cluster.Status.Cloud.AWS.SubnetId),
	})
	cloud.Logger(cm.ctx).Debug("Associating route table to subnet", r3, err)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cloud.Logger(cm.ctx).Infof("Route table %v associated to subnet %v", cm.cluster.Status.Cloud.AWS.RouteTableId, cm.cluster.Status.Cloud.AWS.SubnetId)

	r4, err := cm.conn.ec2.CreateRoute(&_ec2.CreateRouteInput{
		RouteTableId:         StringP(cm.cluster.Status.Cloud.AWS.RouteTableId),
		DestinationCidrBlock: StringP("0.0.0.0/0"),
		GatewayId:            StringP(cm.cluster.Status.Cloud.AWS.IGWId),
	})
	cloud.Logger(cm.ctx).Debug("Added route to route table", r4, err)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cloud.Logger(cm.ctx).Infof("Route added to route table %v", cm.cluster.Status.Cloud.AWS.RouteTableId)
	return nil
}

func (cm *ClusterManager) setupSecurityGroups() error {
	var ok bool
	var err error
	if cm.cluster.Status.Cloud.AWS.MasterSGId, ok, err = cm.getSecurityGroupId(cm.cluster.Spec.Cloud.AWS.MasterSGName); !ok {
		if err != nil {
			return errors.FromErr(err).WithContext(cm.ctx).Err()
		}
		err = cm.createSecurityGroup(cm.cluster.Spec.Cloud.AWS.MasterSGName, "Kubernetes security group applied to master instance")
		if err != nil {
			return errors.FromErr(err).WithContext(cm.ctx).Err()
		}
		cloud.Logger(cm.ctx).Infof("Master security group %v created", cm.cluster.Spec.Cloud.AWS.MasterSGName)
	}
	if cm.cluster.Status.Cloud.AWS.NodeSGId, ok, err = cm.getSecurityGroupId(cm.cluster.Spec.Cloud.AWS.NodeSGName); !ok {
		if err != nil {
			return errors.FromErr(err).WithContext(cm.ctx).Err()
		}
		err = cm.createSecurityGroup(cm.cluster.Spec.Cloud.AWS.NodeSGName, "Kubernetes security group applied to node instances")
		if err != nil {
			return errors.FromErr(err).WithContext(cm.ctx).Err()
		}
		cloud.Logger(cm.ctx).Infof("Naster security group %v created", cm.cluster.Spec.Cloud.AWS.NodeSGName)
	}

	err = cm.detectSecurityGroups()
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	cloud.Logger(cm.ctx).Info("Masters can talk to master")
	err = cm.autohrizeIngressBySGID(cm.cluster.Status.Cloud.AWS.MasterSGId, cm.cluster.Status.Cloud.AWS.MasterSGId)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	cloud.Logger(cm.ctx).Info("Nodes can talk to nodes")
	err = cm.autohrizeIngressBySGID(cm.cluster.Status.Cloud.AWS.NodeSGId, cm.cluster.Status.Cloud.AWS.NodeSGId)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	cloud.Logger(cm.ctx).Info("Masters and nodes can talk to each other")
	err = cm.autohrizeIngressBySGID(cm.cluster.Status.Cloud.AWS.MasterSGId, cm.cluster.Status.Cloud.AWS.NodeSGId)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	err = cm.autohrizeIngressBySGID(cm.cluster.Status.Cloud.AWS.NodeSGId, cm.cluster.Status.Cloud.AWS.MasterSGId)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	// TODO(justinsb): Would be fairly easy to replace 0.0.0.0/0 in these rules

	cloud.Logger(cm.ctx).Info("SSH is opened to the world")
	err = cm.autohrizeIngressByPort(cm.cluster.Status.Cloud.AWS.MasterSGId, 22)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	err = cm.autohrizeIngressByPort(cm.cluster.Status.Cloud.AWS.NodeSGId, 22)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	cloud.Logger(cm.ctx).Info("HTTPS to the master is allowed (for API access)")
	err = cm.autohrizeIngressByPort(cm.cluster.Status.Cloud.AWS.MasterSGId, 443)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	err = cm.autohrizeIngressByPort(cm.cluster.Status.Cloud.AWS.MasterSGId, 6443)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	return nil
}

func (cm *ClusterManager) getSecurityGroupId(groupName string) (string, bool, error) {
	cloud.Logger(cm.ctx).Infof("Checking security group %v", groupName)
	r1, err := cm.conn.ec2.DescribeSecurityGroups(&_ec2.DescribeSecurityGroupsInput{
		Filters: []*_ec2.Filter{
			{
				Name: StringP("vpc-id"),
				Values: []*string{
					StringP(cm.cluster.Status.Cloud.AWS.VpcId),
				},
			},
			{
				Name: StringP("group-name"),
				Values: []*string{
					StringP(groupName),
				},
			},
			{
				Name: StringP("tag:KubernetesCluster"),
				Values: []*string{
					StringP(cm.cluster.Name),
				},
			},
		},
	})
	cloud.Logger(cm.ctx).Debug("Retrieved security group", r1, err)
	if err != nil {
		return "", false, errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	if len(r1.SecurityGroups) == 0 {
		cloud.Logger(cm.ctx).Infof("No security group %v found", groupName)
		return "", false, nil
	}
	cloud.Logger(cm.ctx).Infof("Security group %v found", groupName)
	return *r1.SecurityGroups[0].GroupId, true, nil
}

func (cm *ClusterManager) createSecurityGroup(groupName string, description string) error {
	cloud.Logger(cm.ctx).Infof("Creating security group %v", groupName)
	r2, err := cm.conn.ec2.CreateSecurityGroup(&_ec2.CreateSecurityGroupInput{
		GroupName:   StringP(groupName),
		Description: StringP(description),
		VpcId:       StringP(cm.cluster.Status.Cloud.AWS.VpcId),
	})
	cloud.Logger(cm.ctx).Debug("Created security group", r2, err)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	time.Sleep(preTagDelay)
	err = cm.addTag(*r2.GroupId, "KubernetesCluster", cm.cluster.Name)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	return nil
}

func (cm *ClusterManager) detectSecurityGroups() error {
	var ok bool
	var err error
	if cm.cluster.Status.Cloud.AWS.MasterSGId == "" {
		if cm.cluster.Status.Cloud.AWS.MasterSGId, ok, err = cm.getSecurityGroupId(cm.cluster.Spec.Cloud.AWS.MasterSGName); !ok {
			return errors.New("Could not detect Kubernetes master security group.  Make sure you've launched a cluster with appctl").WithContext(cm.ctx).Err()
		} else {
			cloud.Logger(cm.ctx).Infof("Master security group %v with id %v detected", cm.cluster.Spec.Cloud.AWS.MasterSGName, cm.cluster.Status.Cloud.AWS.MasterSGId)
		}
	}
	if cm.cluster.Status.Cloud.AWS.NodeSGId == "" {
		if cm.cluster.Status.Cloud.AWS.NodeSGId, ok, err = cm.getSecurityGroupId(cm.cluster.Spec.Cloud.AWS.NodeSGName); !ok {
			return errors.New("Could not detect Kubernetes node security group.  Make sure you've launched a cluster with appctl").WithContext(cm.ctx).Err()
		} else {
			cloud.Logger(cm.ctx).Infof("Node security group %v with id %v detected", cm.cluster.Spec.Cloud.AWS.NodeSGName, cm.cluster.Status.Cloud.AWS.NodeSGId)
		}
	}
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	return nil
}

func (cm *ClusterManager) autohrizeIngressBySGID(groupID string, srcGroup string) error {
	r1, err := cm.conn.ec2.AuthorizeSecurityGroupIngress(&_ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: StringP(groupID),
		IpPermissions: []*_ec2.IpPermission{
			{
				IpProtocol: StringP("-1"),
				UserIdGroupPairs: []*_ec2.UserIdGroupPair{
					{
						GroupId: StringP(srcGroup),
					},
				},
			},
		},
	})
	cloud.Logger(cm.ctx).Debug("Authorized ingress", r1, err)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cloud.Logger(cm.ctx).Infof("Ingress authorized into SG %v from SG %v", groupID, srcGroup)
	return nil
}

func (cm *ClusterManager) autohrizeIngressByPort(groupID string, port int64) error {
	r1, err := cm.conn.ec2.AuthorizeSecurityGroupIngress(&_ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: StringP(groupID),
		IpPermissions: []*_ec2.IpPermission{
			{
				IpProtocol: StringP("tcp"),
				FromPort:   Int64P(port),
				IpRanges: []*_ec2.IpRange{
					{
						CidrIp: StringP("0.0.0.0/0"),
					},
				},
				ToPort: Int64P(port),
			},
		},
	})
	cloud.Logger(cm.ctx).Debug("Authorized ingress", r1, err)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cloud.Logger(cm.ctx).Infof("Authorized ingress into SG %v via port %v", groupID, port)
	return nil
}

//
// -------------------------------------
//
func (cm *ClusterManager) startMaster() (*api.Instance, error) {
	var err error
	// TODO: FixIt!
	//cm.cluster.Spec.MasterDiskId, err = cm.ensurePd(cm.namer.MasterPDName(), cm.cluster.Spec.MasterDiskType, cm.cluster.Spec.MasterDiskSize)
	//if err != nil {
	//	return nil, errors.FromErr(err).WithContext(cm.ctx).Err()
	//}
	err = cm.reserveIP()
	if err != nil {
		return nil, errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cloud.Store(cm.ctx).Clusters().UpdateStatus(cm.cluster) // needed for master start-up config

	masterInstanceID, err := cm.createMasterInstance(cm.cluster.Spec.KubernetesMasterName, api.RoleKubernetesMaster)
	if err != nil {
		return nil, errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cloud.Logger(cm.ctx).Info("Waiting for master instance to be ready")
	// We are not able to add an elastic ip, a route or volume to the instance until that instance is in "running" state.
	err = cm.waitForInstanceState(masterInstanceID, "running")
	if err != nil {
		return nil, errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cloud.Logger(cm.ctx).Info("Master instance is ready")
	if cm.cluster.Spec.MasterReservedIP != "" {
		err = cm.assignIPToInstance(masterInstanceID)
		if err != nil {
			return nil, errors.FromErr(err).WithMessage("failed to assign ip").WithContext(cm.ctx).Err()
		}
	}

	// TODO check setting master IP is set properly
	masterInstance, err := cm.newKubeInstance(masterInstanceID) // sets external IP
	if err != nil {
		return nil, errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	masterInstance.Spec.Role = api.RoleKubernetesMaster
	cm.cluster.Spec.MasterExternalIP = masterInstance.Status.PublicIP
	cloud.Store(cm.ctx).Instances(cm.cluster.Name).Create(masterInstance)

	err = cloud.EnsureARecord(cm.ctx, cm.cluster, masterInstance) // works for reserved or non-reserved mode
	if err != nil {
		return masterInstance, errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	_, err = cloud.Store(cm.ctx).Clusters().UpdateStatus(cm.cluster) // needed for node start-up config to get master_internal_ip
	// This is a race between instance start and volume attachment.
	// There appears to be no way to start an AWS instance with a volume attached.
	// To work around this, we wait for volume to be ready in setup-master-pd.sh
	if err != nil {
		return masterInstance, errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	/*
		r1, err := cm.conn.ec2.AttachVolume(&_ec2.AttachVolumeInput{
			VolumeId:   StringP(cm.cluster.Spec.MasterDiskId),
			Device:     StringP("/dev/sdb"),
			InstanceId: StringP(masterInstanceID),
		})
		cloud.Logger(cm.ctx).Debug("Attached persistent data volume to master", r1, err)
		if err != nil {
			return masterInstance, errors.FromErr(err).WithContext(cm.ctx).Err()
		}
		cloud.Logger(cm.ctx).Infof("Persistent data volume %v attatched to master", cm.cluster.Spec.MasterDiskId)
	*/

	time.Sleep(15 * time.Second)
	r2, err := cm.conn.ec2.CreateRoute(&_ec2.CreateRouteInput{
		RouteTableId:         StringP(cm.cluster.Status.Cloud.AWS.RouteTableId),
		DestinationCidrBlock: StringP(cm.cluster.Spec.Networking.MasterSubnet),
		InstanceId:           StringP(masterInstanceID),
	})
	cloud.Logger(cm.ctx).Debug("Created route to master", r2, err)
	if err != nil {
		return masterInstance, errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cloud.Logger(cm.ctx).Infof("Master route to route table %v for ip %v created", cm.cluster.Status.Cloud.AWS.RouteTableId, masterInstanceID)
	return masterInstance, nil
}

func (cm *ClusterManager) ensurePd(name, diskType string, sizeGb int64) (string, error) {
	volumeId, err := cm.findPD(name)
	if err != nil {
		return volumeId, errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	if volumeId == "" {
		// name := cluster.Spec.ctx.KubernetesMasterName + "-pd"
		r1, err := cm.conn.ec2.CreateVolume(&_ec2.CreateVolumeInput{
			AvailabilityZone: &cm.cluster.Spec.Cloud.Zone,
			VolumeType:       &diskType,
			Size:             Int64P(sizeGb),
		})
		cloud.Logger(cm.ctx).Debug("Created master pd", r1, err)
		if err != nil {
			return "", errors.FromErr(err).WithContext(cm.ctx).Err()
		}
		volumeId = *r1.VolumeId
		cloud.Logger(cm.ctx).Infof("Master disk with size %vGB, type %v created", cm.cluster.Spec.MasterDiskSize, cm.cluster.Spec.MasterDiskType)

		time.Sleep(preTagDelay)
		err = cm.addTag(volumeId, "Name", name)
		if err != nil {
			return volumeId, errors.FromErr(err).WithContext(cm.ctx).Err()
		}
		err = cm.addTag(volumeId, "KubernetesCluster", cm.cluster.Name)
		if err != nil {
			return volumeId, errors.FromErr(err).WithContext(cm.ctx).Err()
		}
	}
	return volumeId, nil
}

func (cm *ClusterManager) findPD(name string) (string, error) {
	// name := cluster.Spec.ctx.KubernetesMasterName + "-pd"
	cloud.Logger(cm.ctx).Infof("Searching master pd %v", name)
	r1, err := cm.conn.ec2.DescribeVolumes(&_ec2.DescribeVolumesInput{
		Filters: []*_ec2.Filter{
			{
				Name: StringP("availability-zone"),
				Values: []*string{
					StringP(cm.cluster.Spec.Cloud.Zone),
				},
			},
			{
				Name: StringP("tag:Name"),
				Values: []*string{
					StringP(name),
				},
			},
			{
				Name: StringP("tag:KubernetesCluster"),
				Values: []*string{
					StringP(cm.cluster.Name),
				},
			},
		},
	})
	cloud.Logger(cm.ctx).Debug("Retrieved master pd", r1, err)
	if err != nil {
		return "", errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	if len(r1.Volumes) > 0 {
		cloud.Logger(cm.ctx).Infof("Found master pd %v", name)
		return *r1.Volumes[0].VolumeId, nil
	}
	cloud.Logger(cm.ctx).Infof("Master pd %v not found", name)
	return "", nil
}

func (cm *ClusterManager) reserveIP() error {
	// Check that MASTER_RESERVED_IP looks like an IPv4 address
	// if match, _ := regexp.MatchString("^[0-9]+.[0-9]+.[0-9]+.[0-9]+$", cluster.Spec.ctx.MasterReservedIP); !match {
	if cm.cluster.Spec.MasterReservedIP == "auto" {
		r1, err := cm.conn.ec2.AllocateAddress(&_ec2.AllocateAddressInput{
			Domain: StringP("vpc"),
		})
		cloud.Logger(cm.ctx).Debug("Allocated elastic IP", r1, err)
		if err != nil {
			return errors.FromErr(err).WithContext(cm.ctx).Err()
		}
		time.Sleep(5 * time.Second)
		cm.cluster.Spec.MasterReservedIP = *r1.PublicIp
		cloud.Logger(cm.ctx).Infof("Elastic IP %v allocated", cm.cluster.Spec.MasterReservedIP)
	}
	return nil
}

func (cm *ClusterManager) createMasterInstance(instanceName string, role string) (string, error) {
	kubeStarter, err := cloud.RenderStartupScript(cm.ctx, cm.cluster, api.RoleKubernetesMaster)
	if err != nil {
		return "", err
	}
	req := &_ec2.RunInstancesInput{
		ImageId:  StringP(cm.cluster.Spec.Cloud.InstanceImage),
		MaxCount: Int64P(1),
		MinCount: Int64P(1),
		//// http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html
		//BlockDeviceMappings: []*_ec2.BlockDeviceMapping{
		//	// MASTER_BLOCK_DEVICE_MAPPINGS
		//	{
		//		// https://github.com/appscode/kubernetes/blob/55d9dec8eb5eb02e1301045b7b81bbac689c86a1/cluster/aws/util.sh#L397
		//		DeviceName: StringP(cm.cluster.Spec.RootDeviceName),
		//		Ebs: &_ec2.EbsBlockDevice{
		//			DeleteOnTermination: TrueP(),
		//			VolumeSize:          Int64P(cm.cluster.Spec.MasterDiskSize),
		//			VolumeType:          StringP(cm.cluster.Spec.MasterDiskType),
		//		},
		//	},
		//	// EPHEMERAL_BLOCK_DEVICE_MAPPINGS
		//	{
		//		DeviceName:  StringP("/dev/sdc"),
		//		VirtualName: StringP("ephemeral0"),
		//	},
		//	{
		//		DeviceName:  StringP("/dev/sdd"),
		//		VirtualName: StringP("ephemeral1"),
		//	},
		//	{
		//		DeviceName:  StringP("/dev/sde"),
		//		VirtualName: StringP("ephemeral2"),
		//	},
		//	{
		//		DeviceName:  StringP("/dev/sdf"),
		//		VirtualName: StringP("ephemeral3"),
		//	},
		//},
		IamInstanceProfile: &_ec2.IamInstanceProfileSpecification{
			Name: StringP(cm.cluster.Spec.Cloud.AWS.IAMProfileMaster),
		},
		InstanceType: StringP(cm.cluster.Spec.MasterSKU),
		KeyName:      StringP(cm.cluster.Status.SSHKeyExternalID),
		Monitoring: &_ec2.RunInstancesMonitoringEnabled{
			Enabled: TrueP(),
		},
		NetworkInterfaces: []*_ec2.InstanceNetworkInterfaceSpecification{
			{
				AssociatePublicIpAddress: TrueP(),
				DeleteOnTermination:      TrueP(),
				DeviceIndex:              Int64P(0),
				Groups: []*string{
					StringP(cm.cluster.Status.Cloud.AWS.MasterSGId),
				},
				PrivateIpAddresses: []*_ec2.PrivateIpAddressSpecification{
					{
						PrivateIpAddress: StringP(cm.cluster.Spec.MasterInternalIP),
						Primary:          TrueP(),
					},
				},
				SubnetId: StringP(cm.cluster.Status.Cloud.AWS.SubnetId),
			},
		},
		UserData: StringP(base64.StdEncoding.EncodeToString([]byte(kubeStarter))),
	}
	r1, err := cm.conn.ec2.RunInstances(req)
	cloud.Logger(cm.ctx).Debug("Created instance", r1, err)
	if err != nil {
		return "", errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cloud.Logger(cm.ctx).Infof("Instance %v created with role %v", instanceName, role)
	instanceID := *r1.Instances[0].InstanceId
	time.Sleep(preTagDelay)

	err = cm.addTag(instanceID, "Name", cm.cluster.Spec.KubernetesMasterName)
	if err != nil {
		return instanceID, errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	err = cm.addTag(instanceID, "Role", role)
	if err != nil {
		return "", errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	err = cm.addTag(instanceID, "KubernetesCluster", cm.cluster.Name)
	if err != nil {
		return "", errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	return instanceID, nil
}

func (cm *ClusterManager) getInstancePublicIP(instanceID string) (string, bool, error) {
	r1, err := cm.conn.ec2.DescribeInstances(&_ec2.DescribeInstancesInput{
		InstanceIds: []*string{StringP(instanceID)},
	})
	cloud.Logger(cm.ctx).Debug("Retrieved Public IP for Instance", r1, err)
	if err != nil {
		return "", false, errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	if r1.Reservations != nil && r1.Reservations[0].Instances != nil && r1.Reservations[0].Instances[0].NetworkInterfaces != nil {
		cloud.Logger(cm.ctx).Infof("Public ip for instance id %v retrieved", instanceID)
		return *r1.Reservations[0].Instances[0].NetworkInterfaces[0].Association.PublicIp, true, nil
	}
	return "", false, nil
}

func (cm *ClusterManager) listInstances(groupName string) ([]*api.Instance, error) {
	r2, err := cm.conn.autoscale.DescribeAutoScalingGroups(&autoscaling.DescribeAutoScalingGroupsInput{
		AutoScalingGroupNames: []*string{
			StringP(groupName),
		},
	})
	if err != nil {
		cm.cluster.Status.Reason = err.Error()
		return nil, errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	instances := make([]*api.Instance, 0)
	for _, group := range r2.AutoScalingGroups {
		for _, instance := range group.Instances {
			ki, err := cm.newKubeInstance(*instance.InstanceId)
			if err != nil {
				return nil, errors.FromErr(err).WithContext(cm.ctx).Err()
			}
			ki.Spec.Role = api.RoleKubernetesPool
			instances = append(instances, ki)
		}
	}
	return instances, nil
}
func (cm *ClusterManager) newKubeInstance(instanceID string) (*api.Instance, error) {
	r1, err := cm.conn.ec2.DescribeInstances(&_ec2.DescribeInstancesInput{
		InstanceIds: []*string{StringP(instanceID)},
	})
	cloud.Logger(cm.ctx).Debug("Retrieved instance ", r1, err)
	if err != nil {
		return nil, cloud.InstanceNotFound
	}

	// Don't reassign internal_ip for AWS to keep the fixed 172.20.0.9 for master_internal_ip
	i := api.Instance{
		ObjectMeta: metav1.ObjectMeta{
			UID:  phid.NewKubeInstance(),
			Name: *r1.Reservations[0].Instances[0].PrivateDnsName,
		},
		Spec: api.InstanceSpec{
			SKU: *r1.Reservations[0].Instances[0].InstanceType,
		},
		Status: api.InstanceStatus{
			ExternalID:    instanceID,
			ExternalPhase: *r1.Reservations[0].Instances[0].State.Name,
			PublicIP:      *r1.Reservations[0].Instances[0].PublicIpAddress,
			PrivateIP:     *r1.Reservations[0].Instances[0].PrivateIpAddress,
		},
	}
	/*
		// The low byte represents the state. The high byte is an opaque internal value
		// and should be ignored.
		//
		//    0 : pending
		//    16 : running
		//    32 : shutting-down
		//    48 : terminated
		//    64 : stopping
		//    80 : stopped
	*/
	if i.Status.ExternalPhase == "terminated" {
		i.Status.Phase = api.InstancePhaseDeleted
	} else {
		i.Status.Phase = api.InstancePhaseReady
	}
	return &i, nil
}

func (cm *ClusterManager) allocateElasticIp() (string, error) {
	r1, err := cm.conn.ec2.AllocateAddress(&_ec2.AllocateAddressInput{
		Domain: StringP("vpc"),
	})
	cloud.Logger(cm.ctx).Debug("Allocated elastic IP", r1, err)
	if err != nil {
		return "", errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cloud.Logger(cm.ctx).Infof("Elastic IP %v allocated", *r1.PublicIp)
	time.Sleep(5 * time.Second)
	return *r1.PublicIp, nil
}

func (cm *ClusterManager) assignIPToInstance(instanceID string) error {
	r1, err := cm.conn.ec2.DescribeAddresses(&_ec2.DescribeAddressesInput{
		PublicIps: []*string{StringP(cm.cluster.Spec.MasterReservedIP)},
	})
	cloud.Logger(cm.ctx).Debug("Retrieved allocation ID for elastic IP", r1, err)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cloud.Logger(cm.ctx).Infof("Found allocation id %v for elastic IP %v", r1.Addresses[0].AllocationId, cm.cluster.Spec.MasterReservedIP)
	time.Sleep(1 * time.Minute)

	r2, err := cm.conn.ec2.AssociateAddress(&_ec2.AssociateAddressInput{
		InstanceId:   StringP(instanceID),
		AllocationId: r1.Addresses[0].AllocationId,
	})
	cloud.Logger(cm.ctx).Debug("Attached IP to instance", r2, err)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cloud.Logger(cm.ctx).Infof("IP %v attached to instance %v", cm.cluster.Spec.MasterReservedIP, instanceID)
	return nil
}

func (cm *ClusterManager) createLaunchConfiguration(name, sku string) error {
	// script := cm.RenderStartupScript(cm.cluster, sku, api.RoleKubernetesPool)
	script, err := cloud.RenderStartupScript(cm.ctx, cm.cluster, api.RoleKubernetesPool)
	if err != nil {
		return err
	}
	configuration := &autoscaling.CreateLaunchConfigurationInput{
		LaunchConfigurationName:  StringP(name),
		AssociatePublicIpAddress: BoolP(cm.cluster.Spec.EnableNodePublicIP),
		// http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html
		BlockDeviceMappings: []*autoscaling.BlockDeviceMapping{
			// NODE_BLOCK_DEVICE_MAPPINGS
			{
				// https://github.com/appscode/kubernetes/blob/55d9dec8eb5eb02e1301045b7b81bbac689c86a1/cluster/aws/util.sh#L397
				DeviceName: StringP(cm.cluster.Status.Cloud.AWS.RootDeviceName),
				Ebs: &autoscaling.Ebs{
					DeleteOnTermination: TrueP(),
					VolumeSize:          Int64P(cm.cluster.Spec.NodeDiskSize),
					VolumeType:          StringP(cm.cluster.Spec.NodeDiskType),
				},
			},
			// EPHEMERAL_BLOCK_DEVICE_MAPPINGS
			{
				DeviceName:  StringP("/dev/sdc"),
				VirtualName: StringP("ephemeral0"),
			},
			{
				DeviceName:  StringP("/dev/sdd"),
				VirtualName: StringP("ephemeral1"),
			},
			{
				DeviceName:  StringP("/dev/sde"),
				VirtualName: StringP("ephemeral2"),
			},
			{
				DeviceName:  StringP("/dev/sdf"),
				VirtualName: StringP("ephemeral3"),
			},
		},
		IamInstanceProfile: StringP(cm.cluster.Spec.Cloud.AWS.IAMProfileNode),
		ImageId:            StringP(cm.cluster.Spec.Cloud.InstanceImage),
		InstanceType:       StringP(sku),
		KeyName:            StringP(cm.cluster.Status.SSHKeyExternalID),
		SecurityGroups: []*string{
			StringP(cm.cluster.Status.Cloud.AWS.NodeSGId),
		},
		UserData: StringP(base64.StdEncoding.EncodeToString([]byte(script))),
	}
	r1, err := cm.conn.autoscale.CreateLaunchConfiguration(configuration)
	cloud.Logger(cm.ctx).Debug("Created node configuration", r1, err)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cloud.Logger(cm.ctx).Info("Node configuration created assuming node public ip is enabled")
	return nil
}

func (cm *ClusterManager) createAutoScalingGroup(name, launchConfig string, count int64) error {
	r2, err := cm.conn.autoscale.CreateAutoScalingGroup(&autoscaling.CreateAutoScalingGroupInput{
		AutoScalingGroupName: StringP(name),
		MaxSize:              Int64P(count),
		MinSize:              Int64P(count),
		DesiredCapacity:      Int64P(count),
		AvailabilityZones: []*string{
			StringP(cm.cluster.Spec.Cloud.Zone),
		},
		LaunchConfigurationName: StringP(launchConfig),
		Tags: []*autoscaling.Tag{
			{
				Key:          StringP("Name"),
				ResourceId:   StringP(name),
				ResourceType: StringP("auto-scaling-group"),
				Value:        StringP(name), // node instance prefix LN_1042
			},
			{
				Key:          StringP("Role"),
				ResourceId:   StringP(name),
				ResourceType: StringP("auto-scaling-group"),
				Value:        StringP(cm.cluster.Name + "-node"),
			},
			{
				Key:          StringP("KubernetesCluster"),
				ResourceId:   StringP(name),
				ResourceType: StringP("auto-scaling-group"),
				Value:        StringP(cm.cluster.Name),
			},
		},
		VPCZoneIdentifier: StringP(cm.cluster.Status.Cloud.AWS.SubnetId),
	})
	cloud.Logger(cm.ctx).Debug("Created autoscaling group", r2, err)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cloud.Logger(cm.ctx).Infof("Autoscaling group %v created", name)
	return nil
}

func (cm *ClusterManager) detectMaster() error {
	masterID, err := cm.getInstanceIDFromName(cm.cluster.Spec.KubernetesMasterName)
	if masterID == "" {
		cloud.Logger(cm.ctx).Info("Could not detect Kubernetes master node.  Make sure you've launched a cluster with appctl.")
		//os.Exit(0)
	}
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	masterIP, _, err := cm.getInstancePublicIP(masterID)
	if masterIP == "" {
		cloud.Logger(cm.ctx).Info("Could not detect Kubernetes master node IP.  Make sure you've launched a cluster with appctl")
		os.Exit(0)
	}
	cloud.Logger(cm.ctx).Infof("Using master: %v (external IP: %v)", cm.cluster.Spec.KubernetesMasterName, masterIP)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	return nil
}

func (cm *ClusterManager) getInstanceIDFromName(tagName string) (string, error) {
	r1, err := cm.conn.ec2.DescribeInstances(&_ec2.DescribeInstancesInput{
		Filters: []*_ec2.Filter{
			{
				Name: StringP("tag:Name"),
				Values: []*string{
					StringP(tagName),
				},
			},
			{
				Name: StringP("instance-state-name"),
				Values: []*string{
					StringP("running"),
				},
			},
			{
				Name: StringP("tag:KubernetesCluster"),
				Values: []*string{
					StringP(cm.cluster.Name),
				},
			},
		},
	})
	cloud.Logger(cm.ctx).Debug("Retrieved instace via name", r1, err)
	if err != nil {
		return "", errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	if r1.Reservations != nil && r1.Reservations[0].Instances != nil {
		return *r1.Reservations[0].Instances[0].InstanceId, nil
	}
	return "", nil
}