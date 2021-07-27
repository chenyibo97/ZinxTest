package apis

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"studygo2/zinxtest/mmogame_zinx/core"
	"studygo2/zinxtest/mmogame_zinx/pb"
	"studygo2/zinxtest/ziface"
	"studygo2/zinxtest/znet"
)

type WorldChatApi struct {
	znet.BaseRouter
}

/*func (wm *WorldChatApi) PreHandle(request ziface.IRequest) {
//	panic("implement me")
}*/

func (wm *WorldChatApi) Handle(request ziface.IRequest) {
	//解析协议

	proto_msg := &pb.Talk{}
	err := proto.Unmarshal(request.GetData(), proto_msg)
	if err != nil {
		fmt.Println("talk unmarshal err", err)
	}
	//根据PID找到玩家
	pid, err := request.GetConnection().GetProperty("pid")
	if err != nil {
		fmt.Println("can not found pid property:", err)
	}
	//根据PID找到玩家对象
	player := core.WorldMgrObj.GetPlayerByPid(pid.(int32))
	//将这个消息广播给其他玩家
	player.Talk(proto_msg.Content)
}

/*func (wm *WorldChatApi) PostHandle(request ziface.IRequest) {
	//panic("implement me")
}
*/
