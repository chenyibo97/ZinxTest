package ziface

/*
连接管理模块抽象层
*/

type IConnManager interface {
	Add(connection IConnection)
	Remove(connection IConnection)
	Get(connId uint32) (IConnection, error)
	Len() int
	ClearConn()
}
