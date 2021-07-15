package znet

import (
	"fmt"
	"net"
	"studygo2/zinxtest/ziface"
)

type Connection struct {
	Conn     *net.TCPConn
	ConnId   uint32
	IsClosed bool
	//  HandleApi ziface.HandleFunc
	Exit chan bool

	Router ziface.IRouter
}

func NewConnection(conn *net.TCPConn, ConnId uint32, router ziface.IRouter) *Connection {
	return &Connection{
		conn,
		ConnId,
		false,
		//callback_Api,
		make(chan bool, 1),
		router,
	}
}

func (c *Connection) StartReader() {
	fmt.Println("read start,connID=", c.ConnId)

	defer c.Stop()
	buf := make([]byte, 512)
	for {
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("read buf fail,err:", err)
			continue
		}
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
			data: buf,
		}
		func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.Handle(request)
		}(req)

	}

}

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
func (c *Connection) Send(data []byte) error {
	return nil
}
