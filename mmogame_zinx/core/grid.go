package core

import (
	"fmt"
	"sync"
)

/*

 */

type Grid struct {
	Gid       int
	MinX      int
	MaxX      int
	MinY      int
	MaxY      int
	PlayerIDs map[int]bool
	pIDLock   sync.RWMutex
}

func NewGrid(Gid int, MinX int, MaxX int, MinY int, MaxY int) *Grid {
	return &Grid{
		Gid:       Gid,
		MinX:      MinX,
		MinY:      MinY,
		MaxX:      MaxX,
		MaxY:      MaxY,
		PlayerIDs: make(map[int]bool),
	}
}
func (g *Grid) Add(playerId int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()

	g.PlayerIDs[playerId] = true
}

func (g *Grid) Remove(playerId int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()

	delete(g.PlayerIDs, playerId)
}

func (g *Grid) GetPlayerIDs() (playerId []int) {
	g.pIDLock.RLock()
	defer g.pIDLock.RUnlock()
	playerId = make([]int, len(g.PlayerIDs))
	for k, _ := range g.PlayerIDs {
		playerId = append(playerId, k)
	}
	return
}

func (g *Grid) String() string {

	return fmt.Sprintf("grid id is:%d,minX:%d,maxX:%d,minY:%d,maxY:%d,playerid:%v",
		g.Gid, g.MinX, g.MaxX, g.MinY, g.MaxY, g.PlayerIDs)

}
