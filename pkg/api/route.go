package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nginx_mate_go/pkg/module"
	"nginx_mate_go/pkg/nginx"
	"nginx_mate_go/pkg/service"
)

func AddStaticRoute(c *gin.Context) {
	req := module.StaticTcpRouteAddReq{}
	err := c.BindJSON(&req)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	err = service.AddStaticTcpRoute(req)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	nginx.ReloadChannel <- struct{}{}
	c.Status(http.StatusCreated)
}

func DelStaticRoute(c *gin.Context) {
	port := c.Param("port")
	err := service.DelStaticTcpRoute(port)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	nginx.ReloadChannel <- struct{}{}
	c.Status(http.StatusNoContent)
}
