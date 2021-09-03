package module

type StaticTcpRouteAddReq struct {
	InPort  int    `json:"in_port"`
	InSsl   bool   `json:"in_ssl"`
	OutHost string `json:"out_host"`
	OutPort int    `json:"out_port"`
}
