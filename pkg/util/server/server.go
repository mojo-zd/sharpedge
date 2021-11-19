package server

import (
	"context"
	"fmt"
	"net/http"
	"regexp"

	"git.woa.com/mojoma/sharpedge/pkg/util/server/meta"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/protocol/goai"
	"github.com/gogf/gf/v2/util/gconv"
)

type Server struct {
	config  Config
	openapi *goai.OpenApiV3
}

func New(config ...Config) *Server {
	var cfg = defConfig
	if len(config) > 0 {
		cfg = config[0]
	}
	//fillConfigWithDefaultValues(&cfg)
	return &Server{config: cfg, openapi: new(goai.OpenApiV3)}
}

func (s *Server) Run(_ context.Context) error {
	return http.ListenAndServe(s.config.Address, nil)
}

func (s *Server) Registry(path string, instances ...interface{}) {
	service := new(restful.WebService)
	service.Path(path).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	for _, instance := range instances {
		s.doRegistry(instance, service)
	}

	restful.DefaultContainer.Add(service)

	// registry open api
	s.setOpenApiService()
}

func (s *Server) doRegistry(instance interface{}, service *restful.WebService) {
	apiInstance := NewApiInstance(instance).AcquireMethods()

	for info, params := range apiInstance.GetMethods() {
		if err := ValidateParams(ExpectInputNum, ExpectOutputNum, info.Name, params); err != nil {
			fmt.Printf("skip registry route, api obj is [%s], method is [%s]", info.Parent, info.Name)
		}

		metadata := meta.GetMetadata(params.Inputs[1])
		// add open api path parameter
		for _, param := range s.BuildParameters(metadata) {
			service.Param(service.PathParameter(param, ""))
		}

		service.Route(
			service.
				Method(gstr.ToUpper(metadata.Method)).
				Doc(metadata.Summary).
				Metadata(restfulspec.KeyOpenAPITags, metadata.Tags).
				Path(metadata.Path).
				To(s.newRouteFunction(metadata, info, params)),
		)
	}
}

func (s *Server) newRouteFunction(metadata meta.Metadata, info MethodInfo, params MethodParams) restful.RouteFunction {
	return func(request *restful.Request, response *restful.Response) {
		var start = gtime.Now()
		var err error
		defer func() {
			if exception := recover(); exception != nil {
				if v, ok := exception.(error); ok && gerror.HasStack(v) {
					err = v
				} else {
					err = gerror.New(`exception recovered:` + gconv.String(exception))
				}
			}
			cost := gtime.Now().Sub(start).Milliseconds()
			fmt.Printf("request done, cost:%dms. request path: %s, method:%s\n",
				cost, metadata.Path, gstr.ToUpper(metadata.Method))
			// TODO 通过response输出到客户端
		}()

		reqObj := NewObject(params.Inputs[1])
		if err = Parse(request, reqObj); err != nil {
			return
		}

		outputs := InvokeHttpHandler(request.Request.Context(), info.Value, reqObj)
		content, err := gjson.Encode(outputs[0].Interface())
		_, _ = response.Write(content)
		//var logContentFormat = `request done, cost: %d ms, output: %#v, code: %d, message: "%s", detail: %+v, error: %+v`
	}
}

// BuildParameters return path parameter name
func (s *Server) BuildParameters(metadata meta.Metadata) []string {
	var paramNames []string
	regx, _ := regexp.Compile(`{.[A-Za-z]+}`)
	params := regx.FindAllString(metadata.Path, -1)

	regx, _ = regexp.Compile(`{(.[A-Za-z]+)}`)
	for _, param := range params {
		matches := regx.FindStringSubmatch(param)
		if len(matches) == 0 {
			continue
		}
		paramNames = append(paramNames, matches[len(matches)-1])
	}
	return paramNames
}
