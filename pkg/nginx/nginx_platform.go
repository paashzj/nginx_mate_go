package nginx

import (
	"github.com/paashzj/gutil"
	"go.uber.org/zap"
	"nginx_mate_go/pkg/path"
	"nginx_mate_go/pkg/util"
)

func startNginxPlatform() error {
	stdout, stderr, err := gutil.CallScript(path.NginxStartScript)
	util.Logger().Info("shell result ", zap.String("stdout", stdout), zap.String("stderr", stderr))
	return err
}

func restartNginxPlatform() error {
	stdout, stderr, err := gutil.CallScript(path.NginxRestartScript)
	util.Logger().Info("shell result ", zap.String("stdout", stdout), zap.String("stderr", stderr))
	return err
}
