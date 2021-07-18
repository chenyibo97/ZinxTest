package ziface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetTcpConnection() *net.TCPConn
	GetConnID() uint32
	GetRemoteAddr() net.Addr
	SendMsg(msgid uint32, data []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
