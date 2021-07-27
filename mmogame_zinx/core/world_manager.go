package core

import (
	"sync"
)

type WorldManager struct {
	AOIMgr  *AOIManager
	Players map[int32]*Player
	PLock   sync.RWMutex
}

var WorldMgrObj *WorldManager

func init() {
	WorldMgrObj = &WorldManager{
		//创建世界
		AOIMgr: NewAOIManager(AOI_MIN_X, AOI_MAX_X, AOI_CNTS_X, AOI_MIN_Y, AOI_MAX_Y, AOI_CNTS_Y),
		//初始化Player集合
		Players: make(map[int32]*Player, 32),
	}
}

func (wm *WorldManager) AddPlayer(player *Player) {
	wm.PLock.Lock()
	wm.Players[player.Pid] = player
	wm.PLock.Unlock()

	//将玩家添加到大地图
	wm.AOIMgr.AddToGridByPos(int(player.Pid), player.X, player.Z)

}
func (wm *WorldManager) RemovePlayerByPid(pid int32) {
	player := wm.Players[pid]
	wm.AOIMgr.RemoveFromGridByPos(int(pid), player.X, player.Z)
	wm.PLock.Lock()
	delete(wm.Players, pid)
	wm.PLock.Unlock()

}

func (wm *WorldManager) GetPlayerByPid(pid int32) *Player {
	wm.PLock.RLock()
	defer wm.PLock.RUnlock()
	return wm.Players[pid]
}
func (wm *WorldManager) GetAllPlayer() []*Player {
	wm.PLock.RLock()
	defer wm.PLock.RUnlock()

	players := make([]*Player, 0)
	for _, v := range wm.Players {
		players = append(players, v)
	}
	//fmt.Println("players:",players)
	return players
}
