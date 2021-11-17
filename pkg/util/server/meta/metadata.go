package meta

import (
	"reflect"

	"github.com/gogf/gf/util/gmeta"
)

const (
	TagPath    = `path`
	TagTags    = `tags`
	TagMethod  = `method`
	TagSummary = `summary`
)

type Metadata struct {
	Path    string
	Tags    []string
	Method  string
	Summary string
}

func GetMetadata(obj reflect.Type) Metadata {
	var value reflect.Value
	if obj.Kind() == reflect.Ptr {
		value = reflect.New(obj.Elem())
	} else {
		value = reflect.New(obj).Elem()
	}

	target := value.Interface()
	return Metadata{
		Path:    gmeta.Get(target, TagPath).String(),
		Tags:    gmeta.Get(target, TagTags).Strings(),
		Method:  gmeta.Get(target, TagMethod).String(),
		Summary: gmeta.Get(target, TagSummary).String(),
	}
}
