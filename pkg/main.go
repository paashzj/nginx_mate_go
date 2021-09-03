package main

import (
	"github.com/gin-gonic/gin"
	"nginx_mate_go/pkg/api"
	"nginx_mate_go/pkg/nginx"
	"nginx_mate_go/pkg/util"
	"os"
	"os/signal"
)

func main() {
	util.Logger().Debug("this is a debug msg")
	util.Logger().Info("this is a info msg")
	util.Logger().Error("this is a error msg")
	nginx.Start()
	router := gin.Default()
	router.POST("/v1/nginx/route/static/tcp", api.AddStaticRoute)
	router.DELETE("/v1/nginx/route/static/tcp/:port", api.DelStaticRoute)
	util.Logger().Info("nginx mate started")
	err := router.Run(":31014")
	if err != nil {
		util.Logger().Error("open http server failed")
		return
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			return
		}
	}
}
