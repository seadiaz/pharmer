package scaleway

import (
	"fmt"
	"strings"

	proto "github.com/appscode/api/kubernetes/v1beta1"
	"github.com/appscode/go/errors"
	"github.com/appscode/pharmer/api"
	"github.com/appscode/pharmer/cloud"
	"github.com/cenkalti/backoff"
	sapi "github.com/scaleway/scaleway-cli/pkg/api"
)

func (cm *ClusterManager) Delete(req *proto.ClusterDeleteRequest) error {
	defer cm.cluster.Delete()

	if cm.cluster.Status.Phase == api.ClusterPhasePending {
		cm.cluster.Status.Phase = api.ClusterPhaseFailing
	} else if cm.cluster.Status.Phase == api.ClusterPhaseReady {
		cm.cluster.Status.Phase = api.ClusterPhaseDeleting
	}
	// cloud.Store(cm.ctx).UpdateKubernetesStatus(cm.ctx.PHID, cm.ctx.Status)

	var err error
	if cm.conn == nil {
		cm.conn, err = NewConnector(cm.ctx, cm.cluster)
		if err != nil {
			cm.cluster.Status.Reason = err.Error()
			return errors.FromErr(err).WithContext(cm.ctx).Err()
		}
	}
	cm.namer = namer{cluster: cm.cluster}
	cm.ins, err = cloud.NewInstances(cm.ctx, cm.cluster)
	if err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cm.ins.Instances, err = cloud.Store(cm.ctx).Instances(cm.cluster.Name).List(api.ListOptions{})
	if err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	var errs []string
	if cm.cluster.Status.Reason != "" {
		errs = append(errs, cm.cluster.Status.Reason)
	}

	for _, i := range cm.ins.Instances {
		backoff.Retry(func() error {
			err := cm.conn.client.DeleteServerForce(i.Status.ExternalID)
			if err != nil {
				return err
			}
			return nil
		}, backoff.NewExponentialBackOff())
		cloud.Logger(cm.ctx).Infof("Droplet %v with id %v for clutser is deleted", i.Name, i.Status.ExternalID, cm.cluster.Name)
	}

	if req.ReleaseReservedIp && cm.cluster.Spec.MasterReservedIP != "" {
		backoff.Retry(func() error {
			return cm.releaseReservedIP(cm.cluster.Spec.MasterReservedIP)
		}, backoff.NewExponentialBackOff())
	}

	// Delete SSH key from DB
	if err := cm.deleteSSHKey(); err != nil {
		errs = append(errs, err.Error())
	}

	if err := cloud.DeleteARecords(cm.ctx, cm.cluster); err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		// Preserve statusCause for failed cluster
		if cm.cluster.Status.Phase == api.ClusterPhaseDeleting {
			cm.cluster.Status.Reason = strings.Join(errs, "\n")
		}
		return fmt.Errorf(strings.Join(errs, "\n"))
	}

	cloud.Logger(cm.ctx).Infof("Cluster %v is deleted successfully", cm.cluster.Name)
	return nil
}

func (cm *ClusterManager) releaseReservedIP(ip string) error {
	ips, err := cm.conn.client.GetIPS()
	if err != nil {
		return errors.FromErr(err).Err()
	}
	for _, i := range ips.IPS {
		if i.Address == ip && i.Server == nil {
			err = cm.conn.client.DeleteIP(ip)
			if err != nil {
				return errors.FromErr(err).Err()
			}
		}
	}
	cloud.Logger(cm.ctx).Infof("Floating ip %v deleted", ip)
	return nil
}

func (cm *ClusterManager) deleteSSHKey() (err error) {
	if cm.cluster.Spec.SSHKey != nil {
		backoff.Retry(func() error {
			user, err := cm.conn.client.GetUser()
			if err != nil {
				return err
			}

			sshPubKeys := make([]sapi.ScalewayKeyDefinition, 0)
			for _, k := range user.SSHPublicKeys {
				if k.Fingerprint != cm.cluster.Spec.SSHKey.OpensshFingerprint {
					sshPubKeys = append(sshPubKeys, sapi.ScalewayKeyDefinition{Key: k.Key})
				}
			}

			return cm.conn.client.PatchUserSSHKey(user.ID, sapi.ScalewayUserPatchSSHKeyDefinition{
				SSHPublicKeys: sshPubKeys,
			})
		}, backoff.NewExponentialBackOff())
		cloud.Logger(cm.ctx).Infof("SSH key for cluster %v deleted", cm.cluster.Name)
	}

	if cm.cluster.Spec.SSHKeyPHID != "" {
		//updates := &storage.SSHKey{IsDeleted: 1}
		//cond := &storage.SSHKey{PHID: cm.ctx.SSHKeyPHID}
		//_, err = cloud.Store(cm.ctx).Engine.Update(updates, cond)
	}
	return
}
