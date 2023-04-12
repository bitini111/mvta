/**
 * @Author: sheng
 * @Email: 396039228@qq.com
 * @Date: 2022/9/9 12:03 下午
 * @Desc: TODO
 */

package aliyun_test

import (
	"github.com/bitini111/mvta/log/aliyun"
	"testing"
)

var logger *aliyun.Logger

func init() {
	logger = aliyun.NewLogger()
}

func TestNewLogger(t *testing.T) {
	defer logger.Close()

	logger.Info("info")
}
