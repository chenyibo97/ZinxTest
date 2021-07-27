package apis

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"studygo2/zinxtest/mmogame_zinx/core"
	"studygo2/zinxtest/mmogame_zinx/pb"
	"studygo2/zinxtest/ziface"
)

type MoveApi struct{}

func (m *MoveApi) PreHandle(request ziface.IRequest) {
	//panic("implement me")
}

func (m *MoveApi) Handle(request ziface.IRequest) {
	proto_msg := &pb.Position{}
	err := proto.Unmarshal(request.GetData(), proto_msg)
	if err != nil {
		fmt.Println("move：position Unmarshal err：", err)
		return
	}

	pid, err := request.GetConnection().GetProperty("pid")
	if err != nil {
		fmt.Println("Get property error：", err)
		return
	}

	fmt.Printf("Player pid:=%d,move(%f,%f,%f,%f)", pid, proto_msg.X,
		proto_msg.Y, proto_msg.Z, proto_msg.V)

	player := core.WorldMgrObj.GetPlayerByPid(pid.(int32))
	player.UpdatePos(proto_msg.X, proto_msg.Y, proto_msg.Z, proto_msg.V)
}

func (m *MoveApi) PostHandle(request ziface.IRequest) {
	//panic("implement me")
}
