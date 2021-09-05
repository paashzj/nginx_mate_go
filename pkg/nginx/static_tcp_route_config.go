package nginx

import (
	"bufio"
	"nginx_mate_go/pkg/module"
	"nginx_mate_go/pkg/path"
	"os"
	"path/filepath"
	"strconv"
)

func writeStaticTcpRouteConfig(req module.StaticTcpRouteAddReq) error {
	fileName := filepath.FromSlash(path.NginxStaticTcpRouteDir + "/" + strconv.Itoa(req.InPort) + ".conf")
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.FileMode(0666))
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	// upstream
	writer.WriteString("upstream ")
	writer.WriteString(strconv.Itoa(req.InPort))
	writer.WriteString(" {\n")

	// write server
	for _, endpoint := range req.OutEndpoints {
		writer.WriteString("    server ")
		writer.WriteString(endpoint.Host)
		writer.WriteString(":")
		writer.WriteString(strconv.Itoa(endpoint.Port))
		writer.WriteString(";\n")
	}
	writer.WriteString("}\n")

	writer.WriteString("\n")

	writer.WriteString("server {\n")
	writer.WriteString("    listen ")
	writer.WriteString(strconv.Itoa(req.InPort))
	if req.InSsl {
		writer.WriteString(" ssl")
	}
	writer.WriteString(";\n")
	if req.InSsl {
		writer.WriteString("    ssl_certificate /opt/sh/openresty/nginx/cert/server.crt;\n")
		writer.WriteString("    ssl_certificate_key /opt/sh/openresty/nginx/cert/server.key;\n")
		writer.WriteString("    ssl_verify_client off;\n")
	}
	writer.WriteString("    proxy_connect_timeout 5s;\n")
	writer.WriteString("    proxy_pass ")
	writer.WriteString(strconv.Itoa(req.InPort))
	writer.WriteString(";\n")
	writer.WriteString("}")

	return writer.Flush()
}
