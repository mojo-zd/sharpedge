package server

import (
	"net/http"
	"regexp"
	"testing"
)

func TestServer(t *testing.T) {
	s := New()
	var path = "/api/v1"
	s.Registry(path, &apiTest{})
	http.ListenAndServe(":8888", nil)
}

func TestParam(t *testing.T) {
	regx, _ := regexp.Compile(`{.[A-Za-z]+}`)
	//regx.FindString("/api/v1/{Id}")
	params := regx.FindAllString("/api/v1/{Id}/{Name}", -1)
	t.Log(params)

	regx, _ = regexp.Compile(`{(.[A-Za-z]+)}`)
	for _, param := range params {
		outs := regx.FindStringSubmatch(param)
		if len(outs) == 0 {
			continue
		}
		t.Log("sub:", outs[len(outs)-1])
	}
}
