package server

import (
	"context"
	"reflect"

	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	ExpectInputNum  = 2
	ExpectOutputNum = 2
)

type apiInstance struct {
	value   reflect.Value
	typ     reflect.Type
	methods map[MethodInfo]MethodParams
}

type MethodParams struct {
	Inputs  []reflect.Type
	Outputs []reflect.Type
}

type MethodInfo struct {
	Parent string
	Name   string
	Value  reflect.Value
}

func NewApiInstance(obj interface{}) *apiInstance {
	return &apiInstance{
		value:   reflect.ValueOf(obj),
		typ:     reflect.TypeOf(obj),
		methods: make(map[MethodInfo]MethodParams),
	}
}

func (in *apiInstance) AcquireMethods() *apiInstance {
	parentName := in.typ.Name()
	for i := 0; i < in.value.NumMethod(); i++ {
		method := in.value.Method(i)
		name := in.typ.Method(i).Name
		inputs, outs := in.InAndOuts(method)
		in.methods[MethodInfo{Name: name, Value: method, Parent: parentName}] = MethodParams{Inputs: inputs, Outputs: outs}
	}
	return in
}

func (in *apiInstance) GetMethods() map[MethodInfo]MethodParams {
	return in.methods
}

func (in *apiInstance) InAndOuts(method reflect.Value) (inputs, outs []reflect.Type) {
	for i := 0; i < method.Type().NumIn(); i++ {
		inputs = append(inputs, method.Type().In(i))
	}

	for i := 0; i < method.Type().NumOut(); i++ {
		outs = append(outs, method.Type().Out(i))
	}
	return
}

func (in *apiInstance) Invoke(ctx context.Context, method reflect.Value, inputs []reflect.Type) ([]reflect.Value, error) {
	var (
		newInputs []reflect.Value
		err       error
	)
	for i, input := range inputs {
		if i == 0 {
			if input.Name() == "Context" {
				newInputs = append(newInputs, reflect.ValueOf(ctx))
				continue
			}
			err = gerror.Newf("first parameter must be context.Context")
			break
		}
		if input.Kind() == reflect.Ptr {
			newInputs = append(newInputs, reflect.New(input.Elem()))
		} else {
			newInputs = append(newInputs, reflect.New(input).Elem())
		}
	}

	return method.Call(newInputs), err
}

func InvokeHttpHandler(ctx context.Context, method, input reflect.Value) []reflect.Value {
	return method.Call([]reflect.Value{reflect.ValueOf(ctx), input})
}

func NewObject(input reflect.Type) (value reflect.Value) {
	if input.Kind() == reflect.Ptr {
		value = reflect.New(input.Elem())
	} else {
		value = reflect.New(input).Elem()
	}
	return
}

func ValidateParams(expectInputNum, expectOutputNum int, name string, params MethodParams) error {
	if len(params.Inputs) != expectInputNum {
		return gerror.Newf("actual method[%s] input number is %d, but expect %d input parameters", name, len(params.Inputs), expectInputNum)
	}

	if len(params.Outputs) != expectOutputNum {
		return gerror.Newf("actual method[%s] output number is %d, but expect %d input parameters", name, len(params.Outputs), expectOutputNum)
	}
	return nil
}
