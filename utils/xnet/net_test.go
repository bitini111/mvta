package xnet_test

import (
	"github.com/bitini111/mvta/utils/xnet"
	"testing"
)

func TestParseAddr(t *testing.T) {
	listenAddr, exposeAddr, err := xnet.ParseAddr(":0")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(listenAddr, exposeAddr)
}
