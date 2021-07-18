package ziface

type IServer interface {
	Start()
	Stop()
	Server()
	AddRouter(msgid uint32, router IRouter)
	GetConnmgr() IConnManager
	SetOnConnStart(func(connection IConnection))
	SetOnConnStop(func(connection IConnection))
	CallOnConnStart(connection IConnection)
	CallOnConnStop(connection IConnection)
}
