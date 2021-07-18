package main

import (
	"fmt"
	"io"
	"net"
	"studygo2/zinxtest/utils"
	_ "studygo2/zinxtest/utils"
	"studygo2/zinxtest/znet"
	"time"
)

func main() {
	fmt.Println("client start")
	conn, err := net.Dial("tcp", "0.0.0.0:8999")
	if err != nil {
		fmt.Println("dial fail,err：", err)
	}
	dp := znet.NewDataPack()
	msg1 := znet.NewMessage(0, []byte("hello"))
	msg2 := znet.NewMessage(1, []byte("world"))
	pack1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("pack msg1 fail,err:", err)
	}
	pack2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("pack msg2 fail,err:", err)
	}
	pack1 = append(pack1, pack2...)
	//buf:=make([]byte,512)
	fmt.Println(utils.GlobalObject.MaxPackageSize)
	for {
		conn.Write(pack1)
		/*	conn.Read(buf)
			fmt.Println(len(buf))
			fmt.Println(string(buf))
			message, err := dp.Unpack(buf)
			if err != nil {
				fmt.Println("unpack msg fail,err:", err)
			}
			buf2:=make([]byte,message.GetMsgLen())
			io.ReadFull(conn, buf2)
			fmt.Println(string(buf2))*/
		time.Sleep(5 * time.Second)

		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(conn, headData)
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

			_, err := io.ReadFull(conn, data)
			if err != nil {
				fmt.Println("read data failed,err:", err)
				continue
			}

			fmt.Println(string(data))
		}

	}
}
