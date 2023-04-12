/**
 * @Author: sheng
 * @Email: 396039228@qq.com
 * @Date: 2022/8/30 5:36 下午
 * @Desc: TODO
 */

package logrus_test

import (
	"github.com/bitini111/mvta/log/logrus"
	"testing"
)

var logger *logrus.Logger

func init() {
	logger = logrus.NewLogger()
}

func TestNewLogger(t *testing.T) {
	logger.Warn(`log: warn`)
	logger.Error(`log: error`)
}
