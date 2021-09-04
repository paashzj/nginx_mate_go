package module

type StaticTcpRouteAddReq struct {
	InPort       int           `json:"in_port"`
	InSsl        bool          `json:"in_ssl"`
	OutEndpoints []OutEndpoint `json:"out_endpoints"`
}

type OutEndpoint struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
