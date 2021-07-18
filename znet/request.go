package znet

import "studygo2/zinxtest/ziface"

type Request struct {
	conn ziface.IConnection
	msg  ziface.IMessage
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}
func (r *Request) GetMsgId() uint32 {
	return r.msg.GetMsgId()
}

//func(r *Request)G
