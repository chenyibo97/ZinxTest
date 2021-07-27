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
func (p *Player) Talk(content string) {
	proto_msg := &pb.BroadCast{
		PID: p.Pid,
		Tp:  1, //代表聊天广播
		Data: &pb.BroadCast_Content{
			Content: content,
		},
	}
	//得到当前世界所有的在线玩家
	Players := WorldMgrObj.GetAllPlayer()
	for _, player := range Players {
		fmt.Println(player)
		player.SendMsg(200, proto_msg)
	}
}
func (p *Player) SyncSurround() {
	//获取周边玩家
	pids := WorldMgrObj.AOIMgr.GetPidsByPos(p.X, p.Z)
	fmt.Println("pids:", pids)
	//将当前玩家的位置信息通过200发给周围玩家
	player := make([]*Player, 0, len(pids))
	for _, pid := range pids {
		fmt.Println("get pids", WorldMgrObj.GetPlayerByPid(int32(pid)))
		player = append(player, WorldMgrObj.GetPlayerByPid(int32(pid)))
	}

	//组建msgid=200的proto数据

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
	fmt.Println(player)
	//让别人看到自己
	for _, plyer := range player {
		plyer.SendMsg(200, data)
	}
	//3让自己看到别人
	//3.1组件MSGID:202的Proto数据
	//3.1.1制作pb.player.slice切片
	players_proto_msg := make([]*pb.Player, 0, len(player))
	for _, plyer := range player {
		p := &pb.Player{
			PID: plyer.Pid,
			P: &pb.Position{
				X: plyer.X,
				Y: plyer.Y,
				Z: plyer.Z,
				V: plyer.V,
			},
		}
		players_proto_msg = append(players_proto_msg, p)
	}
	//3.1.2 syncplaer
	SyncPlayer_proto_msg := &pb.SyncPlayers{
		Ps: players_proto_msg,
	}

	//将组建好的数据发送给当前玩家的客户端
	p.SendMsg(202, SyncPlayer_proto_msg)

}

func (p *Player) UpdatePos(x, y, z, v float32) {
	//更新当前玩家的Player对象的坐标

	p.X = x
	p.Y = y
	p.Z = z
	p.V = v
	//组件广播proto协议,msgID:=200 tp-4

	proto_msg := &pb.BroadCast{
		PID: p.Pid,
		Tp:  4,
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	//获取当前玩家的周边玩家AOI九宫格之内的玩家
	players := p.GetSurroudingPlayers()
	for _, player := range players {
		player.SendMsg(200, proto_msg)
	}
}

//获取当前玩家的周边玩家AOU九宫格之内的玩家
func (p *Player) GetSurroudingPlayers() []*Player {
	pids := WorldMgrObj.AOIMgr.GetPidsByPos(p.X, p.Z)
	PlayerS := make([]*Player, 0, len(pids))

	for _, pid := range pids {
		PlayerS = append(PlayerS, WorldMgrObj.GetPlayerByPid(int32(pid)))
	}
	return PlayerS

}

//玩家下线任务
func (p *Player) Offline() {
	fmt.Println("调用了吗？3")
	players := p.GetSurroudingPlayers()
	proto_msg := &pb.SyncPID{
		PID: p.Pid,
	}
	fmt.Println("调用了吗？4")
	for _, player := range players {
		player.SendMsg(201, proto_msg)
	}
	fmt.Println("调用了吗？5")
	//将当前玩家从世界树删除
	WorldMgrObj.AOIMgr.RemoveFromGridByPos(int(p.Pid), p.X, p.Z)
	fmt.Println("调用了吗？6")
	WorldMgrObj.RemovePlayerByPid(p.Pid)
	fmt.Println("调用了吗？7")
}
