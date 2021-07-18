package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"studygo2/zinxtest/ziface"
)

var GlobalObject *GlobalObj

type GlobalObj struct {
	TcpServer ziface.IServer
	Host      string `json:"host"`
	TcpPort   int    `json:"tcp_port"`
	Name      string `json:"name"`

	Version        string
	MaxConn        int
	MaxPackageSize uint32
}

func init() {
	GlobalObject = &GlobalObj{
		Name:           "zinxServer",
		Version:        "v0.6",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}
	GlobalObject.Reload()
}

func (g *GlobalObj) Reload() {
	file, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		fmt.Println("read config fail", err)
	}

	err = json.Unmarshal(file, &GlobalObject)
	if err != nil {
		fmt.Println("unmarshal config fail", err)
	}
}
