package xconv_test

import (
	"fmt"
	"github.com/bitini111/mvta/utils/xconv"

	"testing"
)

func TestBytesToString(t *testing.T) {
	b := []byte("abc")

	s := xconv.BytesToString(b)

	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", &s)
}
