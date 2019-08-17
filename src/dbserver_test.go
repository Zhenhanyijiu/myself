package main

import (
	"net/http"
	"testing"
)

const (
	urlstr = "http://192.168.6.95:7777/db/connect"
)

func TestDbServer_ConnDb(t *testing.T) {
	rsp, err := http.Get(urlstr + "/db/connect")
}
