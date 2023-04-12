package acceptorwrapper

import "github.com/bitini111/mvta/acceptor"

type BaseWrapper struct {
	acceptor.Acceptor
	connChan chan acceptor.PlayerConn
	wrapConn func(conn acceptor.PlayerConn) acceptor.PlayerConn
}

func NewBaseWrapper(wrapConn func(acceptor.PlayerConn) acceptor.PlayerConn) BaseWrapper {
	return BaseWrapper{
		connChan: make(chan acceptor.PlayerConn),
		wrapConn: wrapConn,
	}
}

// ListenAndServe starts a goroutine that wraps acceptor's conn
// and calls acceptor's listenAndServe
func (b *BaseWrapper) ListenAndServe() {
	go b.pipe()
	b.Acceptor.ListenAndServe()
}

// GetConnChan returns the wrapper conn chan
func (b *BaseWrapper) GetConnChan() chan acceptor.PlayerConn {
	return b.connChan
}

func (b *BaseWrapper) pipe() {
	for conn := range b.Acceptor.GetConnChan() {
		b.connChan <- b.wrapConn(conn)
	}
}
