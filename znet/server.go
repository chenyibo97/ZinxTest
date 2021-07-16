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
	router    ziface.IRouter
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
	fmt.Printf("[zinx]Server config: ip:%s", s.IP, "port:", s.Port, "name", s.Name)
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
		fmt.Printf("start zinx serverr succ", s.Name, "sucess,listening")
		for {
			conn, err := Listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept err ,", err)
				continue
			}

			connection := NewConnection(conn, cid, s.router)
			cid++
			connection.Start()
		}
	}()

}
func (s *Server) Stop() {
	//todo 将一些服务器的资源回收和停止
}

func (s *Server) Server() {
	s.Start()
	select {}
}
func (s *Server) AddRouter(router ziface.IRouter) {
	s.router = router
	fmt.Println("add router sucess")
}

func NewServer(name string) ziface.IServer {
	return &Server{
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Name:      utils.GlobalObject.Name,
		Port:      utils.GlobalObject.TcpPort,
		router:    nil,
	}
}
