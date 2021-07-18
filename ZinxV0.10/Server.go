package main

import (
	"fmt"
	"studygo2/zinxtest/ziface"
	"studygo2/zinxtest/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (b *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("call preHandle")
	data := request.GetData()
	fmt.Println(string(data))

}
func (b *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call Handle")
	err := request.GetConnection().SendMsg(3, []byte("ping ping ping"))
	if err != nil {
		fmt.Println("handle fail error")
	}
}
func (b *PingRouter) PostHandle(request ziface.IRequest) {
	/*fmt.Println("call postHandle")
	_, err := request.GetConnection().GetTcpConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("call back  after ping error")
	}*/
}

type HelloRouter struct {
	znet.BaseRouter
}

func (b *HelloRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("call preHandle")
	data := request.GetData()
	fmt.Println(string(data))

}
func (b *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Println("call Handle")
	err := request.GetConnection().SendMsg(3, []byte("hello"))
	if err != nil {
		fmt.Println("handle fail error")
	}
}
func (b *HelloRouter) PostHandle(request ziface.IRequest) {
	/*fmt.Println("call postHandle")
	_, err := request.GetConnection().GetTcpConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("call back  after ping error")
	}*/
}

func main() {
	server := znet.NewServer("test")
	server.SetOnConnStart(DoConectionBegin)
	server.SetOnConnStop(DoConectionEnd)
	server.AddRouter(0, &PingRouter{})
	server.AddRouter(1, &HelloRouter{})
	server.Server()
}
func DoConectionBegin(conn ziface.IConnection) {

	fmt.Println("-----> call DoConectionBegin")
	fmt.Println("set conn name,home,hoe")
	conn.SetProperty("name", "yibo")

}
func DoConectionEnd(conn ziface.IConnection) {
	fmt.Println("-----> call DoConectionEnd")
	if name, err := conn.GetProperty("name"); err == nil {
		fmt.Println("name=", name)
	}
}