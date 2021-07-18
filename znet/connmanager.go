package znet

import (
	"errors"
	"fmt"
	"studygo2/zinxtest/ziface"
	"sync"
)

type ConnManager struct {
	connection map[uint32]ziface.IConnection
	connLock   sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connection: make(map[uint32]ziface.IConnection),
	}
}

func (c *ConnManager) Add(conn ziface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	c.connection[conn.GetConnID()] = conn
	fmt.Println("connection add to connmgr suncessfully,conn num:", c.Len())
}

func (c *ConnManager) Remove(conn ziface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	delete(c.connection, conn.GetConnID())
	fmt.Println("connid=", conn.GetConnID(), "has been remove sucessfully")

}

func (c *ConnManager) Get(connId uint32) (ziface.IConnection, error) {
	c.connLock.RLock()
	defer c.connLock.RUnlock()

	if conn, ok := c.connection[connId]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found!")
	}

}

func (c *ConnManager) Len() int {
	return len(c.connection)
}

func (c *ConnManager) ClearConn() {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	for connid, conn := range c.connection {
		conn.Stop()
		delete(c.connection, connid)
	}
	fmt.Println("clear all connection sucess!conn num=", c.Len())
}
