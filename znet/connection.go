package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"studygo2/zinxtest/ziface"
)

type Connection struct {
	Conn     *net.TCPConn
	ConnId   uint32
	IsClosed bool
	//  HandleApi ziface.HandleFunc
	Exit chan bool

	//Router ziface.IRouter
	MsgHandle ziface.IMsgHandle
}

func NewConnection(conn *net.TCPConn, ConnId uint32, MsgHandle ziface.IMsgHandle) *Connection {
	return &Connection{
		conn,
		ConnId,
		false,
		//callback_Api,
		make(chan bool, 1),
		MsgHandle,
	}
}

func (c *Connection) StartReader() {
	fmt.Println("read start,connID=", c.ConnId)

	defer c.Stop()
	/*buf := make([]byte, 512)*/

	dp := NewDataPack()

	for {

		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTcpConnection(), headData)
		if err != nil {
			fmt.Println("read head failed,err:", err)
			continue
		}
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack  failed,err:", err)
			continue
		}
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())

			_, err := io.ReadFull(c.GetTcpConnection(), data)
			if err != nil {
				fmt.Println("read data failed,err:", err)
				continue
			}
		}

		msg.SetData(data)

		/*	_, err = c.Conn.Read(buf)
			if err != nil {
				fmt.Println("read buf fail,err:", err)
				continue
			}*/
		/*req:=&Request{
			conn: c,
			data: buf,
		}*/
		/*err = c.HandleApi(c.Conn, buf, cnt)
		if err != nil {
			fmt.Println("ConnID",c.ConnId,"handle is err",err)
			break
		}*/
		req := &Request{
			conn: c,
			msg:  msg,
		}
		go c.MsgHandle.DoMsgHandler(req)

	}

}

/*func (c *Connection)SendMsg(msgId uint32,data []byte){

}*/

func (c *Connection) Start() {

	fmt.Println("connection start,connID=", c.ConnId)
	//TODO 启动从当前连接写数据的业务
	go c.StartReader()
}
func (c *Connection) Stop() {
	fmt.Println("connection stop,connID=", c.ConnId)
	if c.IsClosed == true {
		return
	}
	c.IsClosed = true
	c.Conn.Close()
	close(c.Exit)
	return
}
func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}
func (c *Connection) GetConnID() uint32 {
	return c.ConnId
}
func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) SendMsg(msgid uint32, data []byte) (err error) {
	if c.IsClosed == true {
		err = errors.New("connecting has been closed")
	}
	dp := NewDataPack()
	// msg:=NewMessage()

	binaryMsg, err := dp.Pack(NewMessage(msgid, data))
	if err != nil {
		fmt.Println("pack fail,err:", err)
	}

	_, err = c.Conn.Write(binaryMsg)
	if err != nil {
		fmt.Println("send pack fail,err:", err)
	}
	return
}
