package server

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
)

type (
	apiTest       struct{}
	ConfigListReq struct {
		g.Meta `path:"/nodes" tags:"OSS Meta Node" method:"get" summary:"Retrieve Node List"`
		Name   string
	}

	ConfigListResp struct {
		Name    string
		Message string
	}

	ConfigGetReq struct {
		g.Meta `path:"/nodes/{Id}" tags:"OSS Meta Node" method:"get" summary:"Retrieve Spec Node"`
		Name   string
		Id     string
	}

	ConfigGetResp struct {
		Name     string
		Resource string
	}
)

func (apiTest) List(ctx context.Context, req *ConfigListReq) (*ConfigListResp, error) {
	fmt.Println("list method has be called...")
	return &ConfigListResp{Name: "test", Message: "message..."}, nil
}

func (apiTest) Get(ctx context.Context, req *ConfigGetReq) (*ConfigGetResp, error) {
	return &ConfigGetResp{Name: "test", Resource: req.Id}, nil
}

func TestInstance(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		in := NewApiInstance(apiTest{}).AcquireMethods()

		for m, params := range in.methods {
			var newInputs []reflect.Value
			for i, input := range params.Inputs {
				if i == 0 {
					if input.Name() == "Context" {
						newInputs = append(newInputs, reflect.ValueOf(context.Background()))
						continue
					}
					t.Fatal("first parameter must be context.Context")
				}
				if input.Kind() == reflect.Ptr {
					newInputs = append(newInputs, reflect.New(input.Elem()))
				} else {
					newInputs = append(newInputs, reflect.New(input).Elem())
				}
			}

			t.AssertNil(ValidateParams(ExpectInputNum, ExpectOutputNum, m.Name, params))
			outs, err := in.Invoke(context.Background(), m.Value, params.Inputs)
			t.AssertNil(err)
			t.Log("method:", m.Name, outs)
		}
	})
}
