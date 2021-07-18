package core

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"math/rand"
	"studygo2/zinxtest/mmogame_zinx/pb"
	"studygo2/zinxtest/ziface"
	"sync"
)

type Player struct {
	Pid  int32
	Conn ziface.IConnection
	X    float32
	Y    float32
	Z    float32
	V    float32
}

var PidGen int32 = 1

var IdLock sync.Mutex

func NewPlayer(Conn ziface.IConnection) *Player {
	IdLock.Lock()
	id := PidGen
	PidGen++
	IdLock.Unlock()

	return &Player{
		Pid:  id,
		Conn: Conn,
		X:    float32(160 + rand.Intn(10)), //随机在160坐标点，基于X轴偏移
		Y:    0,
		Z:    float32(140 + rand.Intn(20)),
		V:    0,
	}
}

func (p *Player) SendMsg(msgId uint32, data proto.Message) {
	//将protomessage 序列化
	bytes, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("format failed:", err)
		return
	}
	if p.Conn == nil {
		fmt.Println("send msg failed,the conn is already closed")
		return
	}
	p.Conn.SendMsg(msgId, bytes)
}

//告诉客户端玩家pid，同步已经生成的玩家ID给客户端
func (p *Player) SyncPid() {
	//组件MSGID:0的proto数据
	data := &pb.SyncPID{
		PID: p.Pid,
	}
	//将数据发往客户端
	p.SendMsg(1, data)

}

func (p *Player) BroadCastStartPosition() {
	//组件MSGID:200的proto数据
	data := &pb.BroadCast{
		PID: p.Pid,
		Tp:  2, //代表广播的位置坐标
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	p.SendMsg(200, data)
}