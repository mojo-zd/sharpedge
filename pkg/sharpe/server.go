package sharpe

import (
	"git.woa.com/mojoma/sharpedge/pkg/sharpe/internal/apiv1"
	"git.woa.com/mojoma/sharpedge/pkg/sharpe/internal/cnt"
	"git.woa.com/mojoma/sharpedge/pkg/util/server"
)

func NewServer() *server.Server {
	s := server.New()
	s.Registry(cnt.APIRoot, apiv1.APIConfig)
	return s
}
