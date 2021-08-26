package config

import "github.com/paashzj/gutil"

// nginx config
var (
	ClusterName   string
)

func init() {
	ClusterName = gutil.GetEnvStr("CLUSTER_NAME", "nginx")
}
