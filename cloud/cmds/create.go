package cmds

import (
	"context"
	"errors"
	"strings"

	proto "github.com/appscode/api/kubernetes/v1beta1"
	"github.com/appscode/go/flags"
	"github.com/appscode/go/log"
	"github.com/appscode/pharmer/cloud"
	"github.com/appscode/pharmer/config"
	"github.com/spf13/cobra"
	"github.com/tamalsaha/go-oneliners"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewCmdCreate() *cobra.Command {
	var req proto.ClusterCreateRequest
	nodes := map[string]int{}

	cmd := &cobra.Command{
		Use:               "create",
		Short:             "Create a Kubernetes cluster for a given cloud provider",
		Example:           "create --provider=(aws|gce|cc) --nodes=t1=1,t2=2 --zone=us-central1-f demo-cluster",
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			flags.EnsureRequiredFlags(cmd, "provider", "zone", "nodes")

			if len(args) > 0 {
				req.Name = args[0]
			} else {
				return errors.New("missing cluster name")
			}
			req.NodeGroups = make([]*proto.InstanceGroup, len(nodes))
			ng := 0
			for sku, count := range nodes {
				req.NodeGroups[ng] = &proto.InstanceGroup{
					Sku:   sku,
					Count: int64(count),
				}
				ng++
			}

			cfgFile, _ := config.GetConfigFile(cmd.Flags())
			cfg, err := config.LoadConfig(cfgFile)
			if err != nil {
				log.Fatalln(err)
			}
			ctx := cloud.NewContext(context.TODO(), cfg)

			clusters, err := cloud.Store(ctx).Clusters().List(metav1.ListOptions{})
			if err != nil {
				log.Fatalln(err)
			}
			for _, cluster := range clusters {
				oneliners.FILE(cluster.Name)
				if strings.EqualFold(cluster.Name, req.Name) {
					log.Fatalf("Cluster exists with name %s.", req.Name)
				}
			}
			cm, err := cloud.GetCloudManager(req.Provider, ctx)
			if err != nil {
				log.Fatalln(err)
			}
			oneliners.FILE(cm, req)
			return cm.Create(&req)
		},
	}

	cmd.Flags().StringVar(&req.Provider, "provider", "", "Provider name")
	cmd.Flags().StringVar(&req.Zone, "zone", "", "Cloud provider zone name")
	cmd.Flags().StringVar(&req.GceProject, "gce-project", "", "GCE project name(only applicable to `gce` provider)")
	cmd.Flags().StringToIntVar(&nodes, "nodes", map[string]int{}, "Node set configuration")
	cmd.Flags().StringVar(&req.CredentialUid, "credential-uid", "", "Use preconfigured cloud credential uid")
	cmd.Flags().StringVar(&req.KubernetesVersion, "version", "", "Kubernetes version")
	cmd.Flags().BoolVar(&req.DoNotDelete, "do-not-delete", false, "Set do not delete flag")

	return cmd
}