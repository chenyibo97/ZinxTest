package znet

import (
	"fmt"
	"studygo2/zinxtest/utils"
	"studygo2/zinxtest/ziface"
)

type MsgHandle struct {
	Apis           map[uint32]ziface.IRouter
	TaskQueue      []chan ziface.IRequest
	WorkerPoolSize uint32
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
		Apis:           make(map[uint32]ziface.IRouter),
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
	}
}

//启动一个work工作池
func (m *MsgHandle) StartWorkPool() {
	for i := 0; i < int(m.WorkerPoolSize); i++ {
		m.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)

		go m.startOneWork(i, m.TaskQueue[i])
	}
}

//启动一个worker工作流程
func (m *MsgHandle) startOneWork(workerId int, taskQueue chan ziface.IRequest) {
	fmt.Println("work id=", workerId, "is starting")

	for {
		select {
		case request := <-taskQueue:
			m.DoMsgHandler(request)
		}
	}
}
func (m *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	//1平均分配
	workId := request.GetConnection().GetConnID() % m.WorkerPoolSize
	fmt.Println("add connid=", request.GetConnection().GetConnID(),
		"request MSGID", request.GetMsgId(),
		"to", workId)

	m.TaskQueue[workId] <- request

	//将消息发给对应worker的taskqueue
}
