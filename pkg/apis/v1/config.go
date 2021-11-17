package v1

import (
	"git.woa.com/mojoma/sharpedge/pkg/sharpe/apiref"
	"github.com/gogf/gf/v2/frame/g"
)

// ConfigListReq is the request receiving struct for querying  config list.
type ConfigListReq struct {
	g.Meta  `path:"/configs" tags:"Config Info" method:"get" summary:"Retrieve Config List"`
	Product string `dc:"产品名称"`
	Region  string `dc:"地域"`
	CommonPaginationReq
}

type ConfigListResp struct {
	apiref.ConfigResp
}

type ConfigGetReq struct {
	g.Meta `path:"/configs/{Id}" tags:"Config Info" method:"get" summary:"Get Spec Config"`
	Id     string
}

type ConfigGetResp struct {
	Name string
	Id   string
}
