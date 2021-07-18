package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"studygo2/zinxtest/utils"
	"studygo2/zinxtest/ziface"
)

type Connection struct {
	Conn     *net.TCPConn
	ConnId   uint32
	IsClosed bool
	//  HandleApi ziface.HandleFunc
	Exit chan bool

	MsgChan chan []byte
	//Router ziface.IRouter
	MsgHandle ziface.IMsgHandle
	//当前CONN隶属于哪个指针
	TcpServer ziface.IServer
}

func NewConnection(conn *net.TCPConn, ConnId uint32, MsgHandle ziface.IMsgHandle, TcpServer ziface.IServer) *Connection {
	return &Connection{
		conn,
		ConnId,
		false,
		//callback_Api,
		make(chan bool, 1),
		make(chan []byte),
		MsgHandle,
		TcpServer,
	}

}
func (c *Connection) StartWriter() {
	fmt.Println("[Write goutine is runing]")
	defer fmt.Println(c.GetRemoteAddr().String(), "[conn writer eixt]")
	for {
		select {
		case data := <-c.MsgChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("send data error,", err)
				return
			}
		case <-c.Exit:
			return
		}
	}
}

func (c *Connection) StartReader() {
	fmt.Println("read start,connID=", c.ConnId)
	defer fmt.Println(c.GetRemoteAddr().String(), "[conn reader exit]")
	defer c.Stop()
	/*buf := make([]byte, 512)*/

	dp := NewDataPack()

	for {

		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTcpConnection(), headData)
		if err != nil {
			fmt.Println("read head failed,err:", err)
			break
		}
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack  failed,err:", err)
			break
		}
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())

			_, err := io.ReadFull(c.GetTcpConnection(), data)
			if err != nil {
				fmt.Println("read data failed,err:", err)
				break
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
		if utils.GlobalObject.WorkerPoolSize > 0 {
			c.MsgHandle.SendMsgToTaskQueue(req)
		} else {
			go c.MsgHandle.DoMsgHandler(req)
		}

	}

}

/*func (c *Connection)SendMsg(msgId uint32,data []byte){

}*/

func (c *Connection) Start() {

	fmt.Println("connection start,connID=", c.ConnId)
	//TODO 启动从当前连接写数据的业务
	go c.StartReader()
	go c.StartWriter()
	c.TcpServer.CallOnConnStart(c)
}
func (c *Connection) Stop() {

	fmt.Println("connection stop,connID=", c.ConnId)
	if c.IsClosed == true {
		return
	}
	c.IsClosed = true
	c.Conn.Close()
	c.Exit <- true
	close(c.Exit)
	close(c.MsgChan)
	c.TcpServer.CallOnConnStop(c)
	c.TcpServer.GetConnmgr().Remove(c)
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

	c.MsgChan <- binaryMsg
	return
}
