package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/emicklei/go-restful/v3"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

// Parse parses and merges path/query/form/body parameters into `parsedDataMap` attribute.
func Parse(r *restful.Request, out interface{}) error {
	var (
		err  error
		body []byte
		data = make(map[string]interface{})
	)

	// Merge path parameters.
	for k, v := range r.PathParameters() {
		data[k] = v
	}
	// Merge query parameters.
	for k, v := range r.Request.URL.Query() {
		data[k] = v[0]
	}
	// Merge form parameters.
	if err = r.Request.ParseForm(); err != nil {
		return err
	}
	for k, v := range r.Request.PostForm {
		if !gregex.IsMatchString(`^[\w\-\[\]]+$`, k) && len(r.Request.PostForm) == 1 {
			// It might be JSON/XML content.
			if s := gstr.Trim(k + strings.Join(v, " ")); len(s) > 0 {
				if s[0] == '{' && s[len(s)-1] == '}' || s[0] == '<' && s[len(s)-1] == '>' {
					body = []byte(s)
					break
				}
			}
		} else {
			data[k] = v[0]
		}
	}
	// Merge body parameters.
	if len(body) == 0 {
		body, err = ioutil.ReadAll(r.Request.Body)
		if err != nil {
			return gerror.Wrap(err, `read request body failed`)
		}
	}
	body = bytes.TrimSpace(body)
	// Treat the body content as JSON, and it should be JSON object.
	if len(body) > 1 && body[0] == '{' && body[len(body)-1] == '}' {
		if !json.Valid(body) {
			return gerror.Wrap(err, `request body content is not JSON type`)
		}
		m := make(map[string]interface{})
		if err = json.Unmarshal(body, &m); err != nil {
			return gerror.Wrap(err, `json.Unmarshal failed`)
		}
		gutil.MapMerge(data, m)
	}

	return gconv.Scan(data, out)
}
