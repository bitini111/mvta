package xnet_test

import (
	"github.com/bitini111/mvta/utils/xnet"
	"testing"
)

func TestInternalIP(t *testing.T) {
	ip, err := xnet.InternalIP()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(ip)
}

func TestExternalIP(t *testing.T) {
	ip, err := xnet.ExternalIP()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(ip)
}

func TestFulfillAddr(t *testing.T) {
	addr := xnet.FulfillAddr(":3553")

	t.Log(addr)
}
