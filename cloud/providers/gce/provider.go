package gce

import (
	go_ctx "context"

	proto "github.com/appscode/api/kubernetes/v1beta1"
	"github.com/appscode/pharmer/api"
	"github.com/appscode/pharmer/cloud"
	"github.com/appscode/pharmer/config"
	"github.com/appscode/pharmer/context"
)

const (
	UID = "gce"
)

func init() {
	cloud.RegisterProvider(UID, func(cfg *config.PharmerConfig) (cloud.Provider, error) { return &provider{cfg: cfg}, nil })
}

type provider struct {
	cfg *config.PharmerConfig
}

var _ cloud.Provider = &provider{}

func (p *provider) Create(ctx go_ctx.Context, req *proto.ClusterCreateRequest) error {
	return (&clusterManager{ctx: context.NewContext(ctx, p.cfg)}).create(req)
}

func (p *provider) Scale(ctx go_ctx.Context, req *proto.ClusterReconfigureRequest) error {
	return (&clusterManager{ctx: context.NewContext(ctx, p.cfg)}).scale(req)
}

func (p *provider) Delete(ctx go_ctx.Context, req *proto.ClusterDeleteRequest) error {
	return (&clusterManager{ctx: context.NewContext(ctx, p.cfg)}).delete(req)
}

func (p *provider) SetVersion(ctx go_ctx.Context, req *proto.ClusterReconfigureRequest) error {
	return (&clusterManager{ctx: context.NewContext(ctx, p.cfg)}).setVersion(req)
}

func (p *provider) UploadStartupConfig(ctx go_ctx.Context) error {
	c := context.NewContext(ctx, p.cfg)
	conn, err := NewConnector(c, nil)
	if err != nil {
		return err
	}
	cm := &clusterManager{ctx: c, conn: conn}
	return cm.UploadStartupConfig()
}

func (p *provider) GetInstance(ctx go_ctx.Context, md *api.InstanceMetadata) (*api.KubernetesInstance, error) {
	c := context.NewContext(ctx, p.cfg)
	conn, err := NewConnector(c, nil)
	if err != nil {
		return nil, err
	}
	cm := &clusterManager{ctx: c, conn: conn}
	return cm.GetInstance(md)
}

func (p *provider) MatchInstance(i *api.KubernetesInstance, md *api.InstanceMetadata) bool {
	return i.Name == md.Name
}