package service

import (
	"nginx_mate_go/pkg/constant"
	"nginx_mate_go/pkg/module"
	"nginx_mate_go/pkg/storage"
	"strconv"
)

func AddStaticTcpRoute(req module.StaticTcpRouteAddReq) error {
	InPortStr := strconv.Itoa(req.InPort)
	return storage.Acquire().Add(constant.StorageNsStaticTcpRoute, "route-"+InPortStr, req)
}

func DelStaticTcpRoute(port string) error {
	return storage.Acquire().Del(constant.StorageNsStaticTcpRoute, "route-"+port)
}
