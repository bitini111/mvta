package xtime_test

import (
	"github.com/bitini111/mvta/utils/xtime"
	"testing"
)

func TestNow(t *testing.T) {
	t.Log(xtime.Now().Format(xtime.DatetimeLayout))
}

func TestToday(t *testing.T) {
	t.Log(xtime.Today())
}
