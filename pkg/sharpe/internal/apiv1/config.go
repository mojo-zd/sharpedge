package apiv1

import (
	"context"

	v1 "git.woa.com/mojoma/sharpedge/pkg/apis/v1"
)

var (
	APIConfig = apiConfig{}
)

type (
	apiConfig struct{}
)

func (api apiConfig) List(ctx context.Context, req *v1.ConfigListReq) (resp *v1.ConfigListResp, err error) {
	resp = &v1.ConfigListResp{}
	return
}

func (api apiConfig) Get(ctx context.Context, req *v1.ConfigGetReq) (resp *v1.ConfigGetResp, err error) {
	resp = &v1.ConfigGetResp{Name: "test-api", Id: req.Id}
	return
}
