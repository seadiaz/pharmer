package digitalocean

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/appscode/errors"
	_env "github.com/appscode/go/env"
	"github.com/appscode/pharmer/api"
	"github.com/appscode/pharmer/cloud/lib"
	"github.com/appscode/pharmer/contexts"
	"github.com/appscode/pharmer/phid"
	"github.com/appscode/pharmer/storage"
	"github.com/appscode/pharmer/system"
	"github.com/cenkalti/backoff"
	"github.com/digitalocean/godo"
)

type instanceManager struct {
	ctx   *contexts.ClusterContext
	conn  *cloudConnector
	namer namer
}

func (im *instanceManager) GetInstance(md *contexts.InstanceMetadata) (*contexts.KubernetesInstance, error) {
	master := net.ParseIP(md.Name) == nil

	var instance *contexts.KubernetesInstance
	backoff.Retry(func() (err error) {
		const pageSize = 50
		curPage := 0
		for {
			var droplets []godo.Droplet
			droplets, _, err = im.conn.client.Droplets.List(context.TODO(), &godo.ListOptions{
				Page:    curPage,
				PerPage: pageSize,
			})
			if err != nil {
				return
			}
			for _, droplet := range droplets {
				var internalIP string
				internalIP, err = droplet.PrivateIPv4()
				if err != nil {
					return
				}
				if internalIP == md.InternalIP {
					instance, err = im.newKubeInstanceFromDroplet(&droplet)
					if master {
						instance.Role = system.RoleKubernetesMaster
					} else {
						instance.Role = system.RoleKubernetesPool
					}
					return
				}
			}
			curPage++
			if len(droplets) < pageSize {
				break
			}
		}
		return nil
	}, backoff.NewExponentialBackOff())

	if instance == nil {
		return nil, errors.New("No instance found with name", md.Name).WithContext(im.ctx).Err()
	}
	return instance, nil
}

func (im *instanceManager) createInstance(name, role, sku string) (*godo.Droplet, error) {
	startupScript := im.RenderStartupScript(im.ctx.NewScriptOptions(), sku, role)
	imgID, err := strconv.Atoi(im.ctx.InstanceImage)
	if err != nil {
		return nil, errors.FromErr(err).WithContext(im.ctx).Err()
	}
	req := &godo.DropletCreateRequest{
		Name:   name,
		Region: im.ctx.Zone,
		Size:   sku,
		Image:  godo.DropletCreateImage{ID: imgID},
		SSHKeys: []godo.DropletCreateSSHKey{
			{Fingerprint: im.ctx.SSHKey.OpensshFingerprint},
		},
		PrivateNetworking: true,
		IPv6:              false,
		UserData:          startupScript,
	}
	if _env.FromHost().IsPublic() {
		req.SSHKeys = []godo.DropletCreateSSHKey{
			{Fingerprint: im.ctx.SSHKey.OpensshFingerprint},
		}
	}
	droplet, resp, err := im.conn.client.Droplets.Create(context.TODO(), req)
	im.ctx.Logger().V(6).Infoln("do response", resp, " errors", err)
	im.ctx.Notifier.StoreAndNotify(api.JobStatus_Running, fmt.Sprintf("Droplet %v created", droplet.Name))
	return droplet, err
}

func (im *instanceManager) RenderStartupScript(opt *contexts.ScriptOptions, sku, role string) string {
	cmd := lib.StartupConfigFromAPI(opt, role)
	if api.UseFirebase() {
		cmd = lib.StartupConfigFromFirebase(opt, role)
	}

	if role == system.RoleKubernetesMaster {
		return lib.RenderKubeInstaller(opt, sku, role, cmd)
	}
	return lib.RenderKubeStarter(opt, sku, cmd)
}

func (im *instanceManager) applyTag(dropletID int) error {
	_, err := im.conn.client.Tags.TagResources(context.TODO(), "KubernetesCluster:"+im.ctx.Name, &godo.TagResourcesRequest{
		Resources: []godo.Resource{
			{
				ID:   strconv.Itoa(dropletID),
				Type: godo.DropletResourceType,
			},
		},
	})
	im.ctx.Notifier.StoreAndNotify(api.JobStatus_Running, fmt.Sprintf("Tag %v applied to droplet %v", "KubernetesCluster:"+im.ctx.Name, dropletID))
	return err
}

func (im *instanceManager) assignReservedIP(ip string, dropletID int) error {
	action, resp, err := im.conn.client.FloatingIPActions.Assign(context.TODO(), ip, dropletID)
	if err != nil {
		return errors.FromErr(err).WithContext(im.ctx).Err()
	}
	im.ctx.Logger().V(6).Infoln("do response", resp, " errors", err)
	im.ctx.Logger().Debug("Created droplet with name", action.String())
	im.ctx.Notifier.StoreAndNotify(api.JobStatus_Running, fmt.Sprintf("Reserved ip %v assigned to droplet %v", ip, dropletID))
	return nil
}

func (im *instanceManager) newKubeInstance(id int) (*contexts.KubernetesInstance, error) {
	droplet, _, err := im.conn.client.Droplets.Get(context.TODO(), id)
	if err != nil {
		return nil, lib.InstanceNotFound
	}
	return im.newKubeInstanceFromDroplet(droplet)
}

func (im *instanceManager) getInstanceId(name string) (int, error) {
	droplets, _, err := im.conn.client.Droplets.List(context.TODO(), &godo.ListOptions{})
	if err != nil {
		return -1, errors.FromErr(err).WithContext(im.ctx).Err()
	}

	for _, item := range droplets {
		fmt.Println(item.Name, name, "<><><><<><><><><><><><")
		if item.Name == name {
			return item.ID, nil
		}
	}
	return -1, errors.New("Instance not found").Err()
}

func (im *instanceManager) newKubeInstanceFromDroplet(droplet *godo.Droplet) (*contexts.KubernetesInstance, error) {
	var externalIP, internalIP string
	externalIP, err := droplet.PublicIPv4()
	if err != nil {
		return nil, err
	}
	internalIP, err = droplet.PrivateIPv4()
	if err != nil {
		return nil, err
	}

	return &contexts.KubernetesInstance{
		PHID:           phid.NewKubeInstance(),
		ExternalID:     strconv.Itoa(droplet.ID),
		ExternalStatus: droplet.Status,
		Name:           droplet.Name,
		ExternalIP:     externalIP,
		InternalIP:     internalIP,
		SKU:            droplet.SizeSlug,                       // 512mb // convert to SKU
		Status:         storage.KubernetesInstanceStatus_Ready, // droplet.Status == active
	}, nil
}

// reboot does not seem to run /etc/rc.local
func (im *instanceManager) reboot(id int) error {
	im.ctx.Logger().Infof("Rebooting instance %v", id)
	action, _, err := im.conn.client.DropletActions.Reboot(context.TODO(), id)
	if err != nil {
		return errors.FromErr(err).WithContext(im.ctx).Err()
	}
	im.ctx.Logger().Debugf("Instance status %v, %v", action, err)
	im.ctx.Logger().Infof("Instance %v reboot status %v", action.ResourceID, action.Status)
	return nil
}