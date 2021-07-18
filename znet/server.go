package znet

//import "zinx/ziface"
import (
	"errors"
	"fmt"
	"net"
	"studygo2/zinxtest/utils"
	"studygo2/zinxtest/ziface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	//router    ziface.IRouter
	MsgHandle ziface.IMsgHandle
	ConnMgr   ziface.IConnManager
	//server创建连接后自动调用的HOOK函数
	OnConnStart func(conn ziface.IConnection)
	//server销毁连接后自动调用的HOOK函数
	OnConnStop func(conn ziface.IConnection)
}

func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[conn handle]callbacktoclient...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err", err)
		return errors.New("CallBacktToClient error")
	}

	return nil
}

func (s *Server) Start() {
	fmt.Println("[zinx]Server config: ip:", s.IP, "port:", s.Port, "name", s.Name)
	//开启工作池
	s.MsgHandle.StartWorkPool()
	//创建连接
	fmt.Printf("[Start] Server Listening at IP:%s,port:%d,is starting", s.IP, s.Port)
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr fail,", err)
			return
		}
		Listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println(" listening addr fail,", err)
			return
		}
		defer Listener.Close()
		var cid uint32
		fmt.Println("start zinx serverr succ", s.Name, "sucess,listening")
		for {
			conn, err := Listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept err ,", err)
				continue
			}

			if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
				//TODO 给客户端响应一个超出最大连接的错误包
				fmt.Println("too many connection ,maxconn=", utils.GlobalObject.MaxConn)
				conn.Close()
				continue
			}
			connection := NewConnection(conn, cid, s.MsgHandle, s)
			connection.TcpServer.GetConnmgr().Add(connection)
			fmt.Println("添加成功，当前连接数为:", s.ConnMgr.Len(), "最大连接数为:", utils.GlobalObject.MaxConn)
			cid++
			connection.Start()
		}
	}()

}
func (s *Server) Stop() {
	//todo 将一些服务器的资源回收和停止
	fmt.Println("server stop")
	s.ConnMgr.ClearConn()
}

func (s *Server) Server() {
	s.Start()
	select {}
}
func (s *Server) AddRouter(msgid uint32, router ziface.IRouter) {
	s.MsgHandle.AddRouter(msgid, router)
	fmt.Println("add router sucess")
}

func NewServer(name string) ziface.IServer {
	return &Server{
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Name:      utils.GlobalObject.Name,
		Port:      utils.GlobalObject.TcpPort,
		MsgHandle: NewMsgHandle(),
		ConnMgr:   NewConnManager(),
	}
}
func (s *Server) GetConnmgr() ziface.IConnManager {
	return s.ConnMgr
}

func (s *Server) SetOnConnStart(f func(connection ziface.IConnection)) {
	s.OnConnStart = f
}

func (s *Server) SetOnConnStop(f func(connection ziface.IConnection)) {
	s.OnConnStop = f
}

func (s *Server) CallOnConnStart(connection ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("CallOnConnStart")
		s.OnConnStart(connection)
	}
}

func (s *Server) CallOnConnStop(connection ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("CallOnConnStop")
		s.OnConnStop(connection)
	}
}
