package znet

import "studygo2/zinxtest/ziface"

//实现router时，先嵌入这个基类
type BaseRouter struct {
}

//这里不实现，等子类重写
func (b *BaseRouter) PreHandle(request ziface.IRequest) {

}

//处理业务的主方法
func (b *BaseRouter) Handle(request ziface.IRequest) {

}
func (b *BaseRouter) PostHandle(request ziface.IRequest) {

}
