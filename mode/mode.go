package mode

import (
	"flag"
	"github.com/bitini111/mvta/env"
)

const (
	MvtaModeEnvName = "MVTA_MODE"
)

const (
	// DebugMode indicates mvta mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates mvta mode is release.
	ReleaseMode = "release"
	// TestMode indicates mvta mode is test.
	TestMode = "test"
)

var mvataMode string

func init() {
	def := flag.String("mode", DebugMode, "Specify the project run mode")
	mode := env.Get(MvtaModeEnvName, *def).String()
	SetMode(mode)
}

// SetMode 设置运行模式
func SetMode(m string) {
	if m == "" {
		m = DebugMode
	}

	switch m {
	case DebugMode, TestMode, ReleaseMode:
		mvataMode = m
	default:
		panic("mvta mode unknown: " + m + " (available mode: debug test release)")
	}
}

// GetMode 获取运行模式
func GetMode() string {
	return mvataMode
}

// IsDebugMode 是否Debug模式
func IsDebugMode() bool {
	return mvataMode == DebugMode
}

// IsTestMode 是否Test模式
func IsTestMode() bool {
	return mvataMode == TestMode
}

// IsReleaseMode 是否Release模式
func IsReleaseMode() bool {
	return mvataMode == ReleaseMode
}
