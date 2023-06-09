package mvta

import (
	"fmt"
	"github.com/bitini111/mvta/component"
	"github.com/bitini111/mvta/config"
	"github.com/bitini111/mvta/eventbus"
	"github.com/bitini111/mvta/log"
	"github.com/bitini111/mvta/task"
	"runtime"

	"os"
	"os/signal"
	"syscall"
)

type Container struct {
	sig        chan os.Signal
	components []component.Component
}

// NewContainer 创建一个容器
func NewContainer() *Container {
	return &Container{sig: make(chan os.Signal)}
}

// Add 添加组件
func (c *Container) Add(components ...component.Component) {
	c.components = append(c.components, components...)
}

// Serve 启动容器
func (c *Container) Serve() {
	log.Debug(fmt.Sprintf("Welcome to the mvta framework %s, Learn more at %s", Version, Website))

	for _, comp := range c.components {
		comp.Init()
	}

	for _, comp := range c.components {
		comp.Start()
	}

	switch runtime.GOOS {
	case `windows`:
		signal.Notify(c.sig, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	default:
		signal.Notify(c.sig, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGKILL, syscall.SIGTERM)
	}

	sig := <-c.sig

	log.Warnf("process got signal %v, container will close", sig)

	signal.Stop(c.sig)

	for _, comp := range c.components {
		comp.Destroy()
	}

	if err := eventbus.Close(); err != nil {
		log.Errorf("eventbus close failed: %v", err)
	}

	task.Release()

	config.Close()
}
