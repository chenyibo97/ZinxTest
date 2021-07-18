package ziface

type IRequest interface {
	//得到当前连接
	GetConnection() IConnection
	GetData() []byte
	GetMsgId() uint32
	//得到请求的消息数据
}
