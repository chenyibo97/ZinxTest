package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("client start")
	dial, err := net.Dial("tcp", "0.0.0.0:8999")
	if err != nil {
		fmt.Println("dial fail,err：", err)
	}
	for {
		time.Sleep(1 * time.Second)
		buf := make([]byte, 512)
		buf = []byte("i am ok")
		dial.Write(buf)
		_, err := dial.Read(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(string(buf))
	}

}
