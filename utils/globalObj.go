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

	WorkerPoolSize   uint32 `json:"worker_pool_size"`    //WORKER工作池的队列的个数
	MaxWorkerTaskLen uint32 `json:"max_worker_task_len"` //每个worker对应的消息队列的任务数量最大值
}

func init() {
	GlobalObject = &GlobalObj{
		Name:             "zinxServer",
		Version:          "v0.6",
		TcpPort:          8999,
		Host:             "0.0.0.0",
		MaxConn:          1000,
		MaxPackageSize:   4096,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
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
