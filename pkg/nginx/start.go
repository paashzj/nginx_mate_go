package nginx

import (
	"encoding/json"
	"github.com/paashzj/gutil"
	"go.uber.org/zap"
	"nginx_mate_go/pkg/constant"
	"nginx_mate_go/pkg/module"
	"nginx_mate_go/pkg/path"
	"nginx_mate_go/pkg/storage"
	"nginx_mate_go/pkg/util"
	"os"
	"path/filepath"
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
	err = startNginxPlatform()
	if err != nil {
		util.Logger().Error("run start ngx scripts failed ", zap.Error(err))
		return
	}
}

func restartNgx() {
	err := generateNgxConf()
	if err != nil {
		util.Logger().Error("generate ngx config file failed ", zap.Error(err))
		return
	}
	err = restartNginxPlatform()
	if err != nil {
		util.Logger().Error("run restart ngx scripts failed ", zap.Error(err))
		return
	}
}

func generateNgxConf() (err error) {
	err = generateStaticTcpRoute()
	if err != nil {
		return
	}
	// todo
	return nil
}

func generateStaticTcpRoute() (err error) {
	err = os.RemoveAll(filepath.FromSlash(path.NginxStaticTcpRouteDir))
	if err != nil {
		return
	}
	err = os.Mkdir(filepath.FromSlash(path.NginxStaticTcpRouteDir), os.FileMode(0777))
	if err != nil {
		util.Logger().Error("clean directory error")
		return
	}
	keyList, err := storage.Acquire().GetKeyList(constant.StorageNsStaticTcpRoute)
	if err != nil {
		return
	}
	for _, key := range keyList {
		bytes, err := storage.Acquire().Get(constant.StorageNsStaticTcpRoute, key)
		if err != nil {
			return err
		}
		var req module.StaticTcpRouteAddReq
		err = json.Unmarshal(bytes, &req)
		if err != nil {
			return err
		}
		err = writeStaticTcpRouteConfig(req)
		if err != nil {
			return err
		}
	}
	return nil
}
