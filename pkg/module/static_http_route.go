package module

type StaticHttpRouteAddReq struct {
	InPort       int           `json:"in_port"`
	InUrl        string        `json:"in_url"`
	InSsl        bool          `json:"in_ssl"`
	OutEndpoints []OutEndpoint `json:"out_endpoints"`
}
