package digitalocean

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/appscode/errors"
	"github.com/appscode/pharmer/api"
	"github.com/appscode/pharmer/cloud/common"
	"github.com/appscode/pharmer/contexts"
	"github.com/appscode/pharmer/system"
	"github.com/digitalocean/godo"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type InstanceGroupManager struct {
	cm       *clusterManager
	instance common.Instance
	im       *instanceManager
}

func (igm *InstanceGroupManager) AdjustInstanceGroup() error {
	instanceGroupName := igm.cm.namer.GetInstanceGroupName(igm.instance.Type.Sku) //igm.cm.ctx.Name + "-" + strings.Replace(igm.instance.Type.Sku, "_", "-", -1) + "-node"
	found, _, err := igm.GetInstanceGroup(instanceGroupName)
	if err != nil {
		igm.cm.ctx.StatusCause = err.Error()
		return errors.FromErr(err).WithContext(igm.cm.ctx).Err()
	}
	fmt.Println(found)

	igm.cm.ctx.ContextVersion = igm.instance.Type.ContextVersion
	igm.cm.ctx.Load()
	var nodeAdjust int64 = 0
	if found {
		nodeAdjust, _ = common.Mutator(igm.cm.ctx, igm.instance)
	}
	if !found {
		err = igm.createInstanceGroup(igm.instance.Stats.Count)
	} else if igm.instance.Stats.Count == 0 {
		if nodeAdjust < 0 {
			nodeAdjust = -nodeAdjust
		}
		err := igm.deleteInstanceGroup(igm.instance.Type.Sku, nodeAdjust)
		if err != nil {
			igm.cm.ctx.StatusCause = err.Error()
			return errors.FromErr(err).WithContext(igm.cm.ctx).Err()
		}
	} else {
		if nodeAdjust < 0 {
			err := igm.deleteInstanceGroup(igm.instance.Type.Sku, -nodeAdjust)
			if err != nil {
				igm.cm.ctx.StatusCause = err.Error()
				return errors.FromErr(err).WithContext(igm.cm.ctx).Err()
			}
		} else {
			err := igm.createInstanceGroup(nodeAdjust)
			if err != nil {
				igm.cm.ctx.StatusCause = err.Error()
				return errors.FromErr(err).WithContext(igm.cm.ctx).Err()
			}
		}
	}
	return nil
}

func (igm *InstanceGroupManager) GetInstanceGroup(instanceGroup string) (bool, map[string]*contexts.KubernetesInstance, error) {
	var flag bool = false
	igm.im.conn.client.Droplets.List(context.TODO(), &godo.ListOptions{})
	existingNGs := make(map[string]*contexts.KubernetesInstance)
	droplets, _, err := igm.cm.conn.client.Droplets.List(context.TODO(), &godo.ListOptions{})
	if err != nil {
		return flag, existingNGs, errors.FromErr(err).WithContext(igm.cm.ctx).Err()
	}

	for _, item := range droplets {
		if strings.HasPrefix(item.Name, instanceGroup) {
			flag = true
			instance, err := igm.im.newKubeInstance(item.ID)
			if err != nil {
				return flag, existingNGs, errors.FromErr(err).WithContext(igm.cm.ctx).Err()
			}
			instance.Role = system.RoleKubernetesPool
			internalIP, err := item.PrivateIPv4()
			existingNGs[internalIP] = instance
		}

	}
	return flag, existingNGs, nil
}

func (igm *InstanceGroupManager) createInstanceGroup(count int64) error {
	for i := int64(0); i < count; i++ {
		_, err := igm.StartNode()
		if err != nil {
			return errors.FromErr(err).WithContext(igm.cm.ctx).Err()

		}
	}
	return nil
}

func (igm *InstanceGroupManager) deleteInstanceGroup(sku string, count int64) error {
	instances, err := igm.listInstances(sku)
	if err != nil {
		return errors.FromErr(err).WithContext(igm.cm.ctx).Err()
	}
	for _, instance := range instances {
		count--
		dropletID, err := strconv.Atoi(instance.ExternalID)
		if err != nil {
			return errors.FromErr(err).WithContext(igm.cm.ctx).Err()
		}
		err = igm.cm.deleteDroplet(dropletID, instance.InternalIP)
		if err != nil {
			return errors.FromErr(err).WithContext(igm.cm.ctx).Err()
		}
		if count <= 0 {
			break
		}
	}
	return nil
}

func (igm *InstanceGroupManager) listInstances(sku string) ([]*contexts.KubernetesInstance, error) {
	instances := make([]*contexts.KubernetesInstance, 0)
	kc, err := igm.cm.ctx.NewKubeClient()
	if err != nil {
		igm.cm.ctx.StatusCause = err.Error()
		return instances, errors.FromErr(err).WithContext(igm.cm.ctx).Err()

	}
	_, droplets, err := igm.GetInstanceGroup(igm.cm.namer.GetInstanceGroupName(sku))
	if err != nil {
		return instances, errors.FromErr(err).WithContext(igm.cm.ctx).Err()
	}
	nodes, err := kc.Client.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		return instances, errors.FromErr(err).WithContext(igm.cm.ctx).Err()
	}
	for _, n := range nodes.Items {
		nl := api.FromMap(n.GetLabels())
		if nl.GetString(api.NodeLabelKey_SKU) == sku && nl.GetString(api.NodeLabelKey_Role) == "node" {
			if _, found := droplets[n.Name]; found {
				instances = append(instances, droplets[n.Name])
			}
		}
	}
	return instances, nil
}

func (igm *InstanceGroupManager) StartNode() (*contexts.KubernetesInstance, error) {
	droplet, err := igm.im.createInstance(igm.cm.namer.GenNodeName(igm.instance.Type.Sku), system.RoleKubernetesPool, igm.instance.Type.Sku)
	if err != nil {
		igm.cm.ctx.StatusCause = err.Error()
		return nil, errors.FromErr(err).WithContext(igm.cm.ctx).Err()
	}
	// record nodes
	igm.cm.conn.waitForInstance(droplet.ID, "active")
	node, err := igm.im.newKubeInstance(droplet.ID)
	if err != nil {
		igm.cm.ctx.StatusCause = err.Error()
		return nil, errors.FromErr(err).WithContext(igm.cm.ctx).Err()
	}
	igm.im.applyTag(droplet.ID)
	node.Role = system.RoleKubernetesPool
	igm.cm.ins.Instances = append(igm.cm.ins.Instances, node)
	return node, nil
}
