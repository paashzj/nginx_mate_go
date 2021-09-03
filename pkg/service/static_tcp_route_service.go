package service

import (
	"encoding/json"
	"nginx_mate_go/pkg/module"
	"nginx_mate_go/pkg/path"
	"os"
	"path/filepath"
	"strconv"
)

func AddStaticTcpRoute(req module.StaticTcpRouteAddReq) error {
	InPortStr := strconv.Itoa(req.InPort)
	file, err := os.OpenFile(filepath.FromSlash(path.NginxStaticTcpRoute+"/route-"+InPortStr+"-v1.json"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.FileMode(0666))
	if err != nil {
		return err
	}
	defer file.Close()
	bytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
	_, err = file.Write(bytes)
	if err != nil {
		return err
	}
	return err
}

func DelStaticTcpRoute(port string) error {
	return os.Remove(filepath.FromSlash(path.NginxStaticTcpRoute + "/route-" + port + "-v1.json"))
}
