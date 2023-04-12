/**
 * @Author: sheng
 * @Email: 396039228@qq.com
 * @Date: 2022/9/9 6:30 下午
 * @Desc: TODO
 */

package tencent_test

import (
	"testing"

	"github.com/bitini111/mvta/log/tencent"
)

var logger *tencent.Logger

func init() {
	logger = tencent.NewLogger()
}

func TestNewLogger(t *testing.T) {
	defer logger.Close()

	logger.Error("error")
}
