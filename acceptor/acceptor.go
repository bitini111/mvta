package acceptor

import "net"

// PlayerConn iface
type PlayerConn interface {
	GetNextMessage() (b []byte, err error)
	RemoteAddr() net.Addr
	net.Conn
}

// Acceptor type interface
type Acceptor interface {
	ListenAndServe()
	Stop()
	GetAddr() string
	GetConnChan() chan PlayerConn
	EnableProxyProtocol()
}
