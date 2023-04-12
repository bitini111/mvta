package mode_test

import (
	"flag"
	"github.com/bitini111/mvta/mode"
	"testing"
)

func TestGetMode(t *testing.T) {
	flag.Parse()

	t.Log(mode.GetMode())
}
