package nginx

import (
	"github.com/paashzj/gutil"
	"go.uber.org/zap"
	"nginx_mate_go/pkg/path"
	"nginx_mate_go/pkg/util"
)

var ReloadChannel = make(chan struct{})

func Start() {
	startNgx()
	go func() {
		for {
			<-ReloadChannel
			startOrReloadNgx()
		}
	}()
}

func startOrReloadNgx() {
	exists, err := gutil.ProcessExists("master process nginx")
	if err != nil {
		util.Logger().Error("unknown ngx exists ", zap.Error(err))
		return
	}
	if exists {
		restartNgx()
	} else {
		startNgx()
	}
}

func startNgx() {
	err := generateNgxConf()
	if err != nil {
		util.Logger().Error("generate ngx config file failed ", zap.Error(err))
		return
	}
	stdout, stderr, err := gutil.CallScript("bash -x " + path.NginxStartScript)
	if err != nil {
		util.Logger().Error("run start ngx scripts failed ", zap.Error(err))
		return
	}
	util.Logger().Info("shell result ", zap.String("stdout", stdout), zap.String("stderr", stderr))
}

func restartNgx() {
	err := generateNgxConf()
	if err != nil {
		util.Logger().Error("generate ngx config file failed ", zap.Error(err))
		return
	}
	stdout, stderr, err := gutil.CallScript("bash -x " + path.NginxRestartScript)
	if err != nil {
		util.Logger().Error("run restart ngx scripts failed ", zap.Error(err))
		return
	}
	util.Logger().Info("shell result ", zap.String("stdout", stdout), zap.String("stderr", stderr))
}

func generateNgxConf() (err error) {
	err = generateStaticTcpRoute()
	if err != nil {
		return
	}
	// todo
	return nil
}

func generateStaticTcpRoute() error {
	return nil
}
