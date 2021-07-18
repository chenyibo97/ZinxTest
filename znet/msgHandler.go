package znet

import (
	"fmt"
	"studygo2/zinxtest/ziface"
)

type MsgHandle struct {
	Apis map[uint32]ziface.IRouter
}

func (m *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	router, ok := m.Apis[request.GetMsgId()]
	if !ok {
		fmt.Println("api msgID=", request.GetMsgId(), "not found")
	}
	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)

}
func (m *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	if _, ok := m.Apis[msgId]; ok {
		panic(fmt.Sprintf("repeat api,msg id:%d", msgId))
	}
	m.Apis[msgId] = router
	fmt.Println("add api msg:", msgId, "sucess")
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
	}
}
