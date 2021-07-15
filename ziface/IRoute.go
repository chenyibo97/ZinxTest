package ziface

type IRouter interface {
	//处理业务之前的hook'
	PreHandle(request IRequest)
	//处理业务的主方法
	Handle(request IRequest)
	PostHandle(request IRequest)
}
