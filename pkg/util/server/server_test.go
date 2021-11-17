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
	regx, _ := regexp.Compile(`({.+})`)
	//regx.FindString("/api/v1/{Id}")
	t.Log(regx.FindAllString("/api/v1/{Id}/{Name}", -1))

	regx, _ = regexp.Compile(`{(.+)}`)

	t.Log("sub:", regx.FindStringSubmatch("{Id}"))
}
