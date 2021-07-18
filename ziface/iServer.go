package ziface

type IServer interface {
	Start()
	Stop()
	Server()
	AddRouter(msgid uint32, router IRouter)
}
