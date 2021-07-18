package main

import (
	"fmt"
	"studygo2/zinxtest/mmogame_zinx/core"
	"studygo2/zinxtest/utils"
	"studygo2/zinxtest/ziface"
	"studygo2/zinxtest/znet"
)

func main() {
	//创建zinxserer句柄

	s := znet.NewServer("MMO GAME ZINX")

	//创建hook函数
	fmt.Println(utils.GlobalObject.Name)

	//注册路由服务
	s.SetOnConnStart(OnConnectionAdd)

	//启动服务
	s.Server()

}

//hook函数
func OnConnectionAdd(conn ziface.IConnection) {
	//创建一个player对象
	player := core.NewPlayer(conn)
	//给客户端发送msgid=1的消息
	player.SyncPid()
	//给客户端发送msgid==200的消息
	player.BroadCastStartPosition()
	//将玩家添加到world
	core.WorldMgrObj.AddPlayer(player)
	fmt.Println("------>player pi d:", player.Pid, "is arrived")

}
