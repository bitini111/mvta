/**
 * @Author: sheng
 * @Email: 396039228@qq.com
 * @Date: 2022/9/7 11:20 上午
 * @Desc: TODO
 */

package hook

import (
	"github.com/bitini111/mvta/log"
	"github.com/bitini111/mvta/utils/xstack"
	"github.com/sirupsen/logrus"
)

type StackHook struct {
	stackLevel log.Level
	callerSkip int
}

func NewStackHook(stackLevel log.Level, callerSkip int) *StackHook {
	return &StackHook{stackLevel: stackLevel, callerSkip: callerSkip}
}

func (h *StackHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *StackHook) Fire(entry *logrus.Entry) error {
	var level log.Level
	switch entry.Level {
	case logrus.DebugLevel:
		level = log.DebugLevel
	case logrus.InfoLevel:
		level = log.InfoLevel
	case logrus.WarnLevel:
		level = log.WarnLevel
	case logrus.ErrorLevel:
		level = log.ErrorLevel
	case logrus.FatalLevel:
		level = log.FatalLevel
	case logrus.PanicLevel:
		level = log.PanicLevel
	}

	var depth xstack.Depth
	if h.stackLevel != log.NoneLevel && level >= h.stackLevel {
		depth = xstack.Full
		entry.Data["stack_out"] = struct{}{}
	} else {
		depth = xstack.First
	}

	st := xstack.Callers(8+h.callerSkip, depth)
	defer st.Free()
	entry.Data["stack_frames"] = st.Frames()

	return nil
}
