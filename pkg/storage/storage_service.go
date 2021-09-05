package storage

import (
	"github.com/paashzj/gutil"
	"nginx_mate_go/pkg/constant"
	"nginx_mate_go/pkg/path"
)

var storage = gutil.NewFsStorage(path.NginxStorage)

func init() {
	storage.AddNamespace(constant.StorageNsStaticTcpRoute)
}

func Acquire() *gutil.FsStorage {
	return storage
}
