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
	_, err := request.GetConnection().GetTcpConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("call back ping error")
	}

}
func (b *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call preHandle")
	_, err := request.GetConnection().GetTcpConnection().Write([]byte("ping png  ping...\n"))
	if err != nil {
		fmt.Println("ping ping ping error")
	}
}
func (b *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("call preHandle")
	_, err := request.GetConnection().GetTcpConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("call back  after ping error")
	}
}

func main() {
	server := znet.NewServer("test")
	server.AddRouter(&PingRouter{})
	server.Server()
}
