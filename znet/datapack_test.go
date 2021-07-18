package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
	"time"
)

func TestNewDataPack(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("listen fail")
	}
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("listen fail")
			}
			go func(conn net.Conn) {
				dp := NewDataPack()
				for {
					headdata := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headdata)
					if err != nil {
						fmt.Println("read fail")
					}

					unpack, err := dp.Unpack(headdata)
					if err != nil {
						fmt.Println("pack fail", err)
					}

					msg := unpack.(*Message)
					msg.Data = make([]byte, msg.GetMsgLen())
					_, err = io.ReadFull(conn, msg.Data)
					if err != nil {
						fmt.Println("read data fail", err)
					}
					fmt.Println("recv msgid:", msg.Id, "len:", msg.DataLen, "data:", string(msg.Data))
				}
			}(conn)
		}
	}()

	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial fail,err:", err)
	}
	dp := NewDataPack()
	msg1 := &Message{
		Id:      1,
		DataLen: 5,
		Data:    []byte{'z', 'i', 'n', 'x', '.'},
	}
	msg2 := &Message{
		Id:      1,
		DataLen: 5,
		Data:    []byte{'t', 'e', 's', 't', '.'},
	}
	pack1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("pack msg1 fail,err:", err)
	}
	pack2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("pack msg2 fail,err:", err)
	}
	pack1 = append(pack1, pack2...)
	buf := make([]byte, 512)

	for {
		conn.Write(pack1)
		conn.Read(buf)
		message, err := dp.Unpack(buf)
		if err != nil {
			fmt.Println("unpack msg fail,err:", err)
		}
		buf2 := make([]byte, message.GetMsgLen())
		io.ReadFull(conn, buf2)
		fmt.Println(string(buf2))
		time.Sleep(1 * time.Second)
	}

}
