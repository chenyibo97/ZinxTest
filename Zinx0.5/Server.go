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

func main() {
	server := znet.NewServer("test")
	server.AddRouter(&PingRouter{})
	server.Server()
}
