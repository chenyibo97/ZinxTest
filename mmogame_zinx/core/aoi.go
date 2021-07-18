package core

import "fmt"

const (
	AOI_MIN_X  int = 85
	AOI_MAX_X  int = 410
	AOI_CNTS_X int = 10
	AOI_MIN_Y  int = 75
	AOI_MAX_Y  int = 400
	AOI_CNTS_Y int = 20
)

type AOIManager struct {
	MinX  int
	MaxX  int
	CntsX int
	MinY  int
	MaxY  int
	CntsY int
	Grid  map[int]*Grid
}

func NewAOIManager(MinX int, MaxX int, CntsX int, MinY int, MaxY int, CntsY int) *AOIManager {
	aoimgr := &AOIManager{
		MaxY:  MaxY,
		MinY:  MinY,
		MaxX:  MaxX,
		MinX:  MinX,
		CntsX: CntsX,
		CntsY: CntsY,
		Grid:  make(map[int]*Grid),
	}
	for y := 0; y < CntsY; y++ {
		for x := 0; x < CntsX; x++ {
			gid := y*CntsX + x
			aoimgr.Grid[gid] = NewGrid(gid,
				aoimgr.MinX+x*aoimgr.gridWidth(),
				aoimgr.MinX+(x+1)*aoimgr.gridWidth(),
				aoimgr.MinY+y*aoimgr.gridLenth(),
				aoimgr.MinY+(y+1)*aoimgr.gridLenth())
		}
	}
	return aoimgr

}

func (m *AOIManager) gridWidth() int {
	return (m.MaxX - m.MinX) / m.CntsX
}

func (m *AOIManager) gridLenth() int {
	return (m.MaxY - m.MinY) / m.CntsY
}
func (m *AOIManager) String() string {
	s := fmt.Sprintf("AOIManager:\n ,minX:%d,maxX:%d,minY:%d,maxY:%d\",cntsX:%d,cntsY:%d\n",
		m.MinX, m.MaxX, m.MinY, m.MaxY, m.CntsX, m.CntsY)
	for _, grid := range m.Grid {
		s += fmt.Sprintln(grid)
	}
	return s
}

//根据格子GID获得周边格子ID的集合

func (m *AOIManager) GetSurroundGridByGid(gid int) (grids []*Grid) {
	if _, ok := m.Grid[gid]; !ok {
		return
	}
	grids = append(grids, m.Grid[gid])

	idx := gid % m.CntsX
	if idx > 0 {
		grids = append(grids, m.Grid[gid-1])
	}

	if idx < m.CntsX-1 {
		grids = append(grids, m.Grid[gid+1])
	}

	gridsX := make([]int, 0, len(grids))
	for _, v := range grids {
		gridsX = append(gridsX, v.Gid)
	}
	for _, v := range grids {
		if v == nil {
			fmt.Println("4添加了nil")
		}
	}
	for _, v := range gridsX {
		idy := v / m.CntsY

		if idy > 0 {

			grids = append(grids, m.Grid[v-m.CntsX])

		}

		if idy < m.CntsY-1 {
			grids = append(grids, m.Grid[v+m.CntsX])
		}

	}
	return

}
func (m *AOIManager) GetGidsByPos(x, y float32) int {
	idx := (int(x) - m.MinX) / m.gridWidth()
	idy := (int(y) - m.MinY) / m.gridLenth()
	return idy*m.CntsX + idx
}

func (m *AOIManager) GetPidsByPos(x, y float32) (playerIds []int) {
	gid := m.GetGidsByPos(x, y)

	grids := m.GetSurroundGridByGid(gid)
	for _, v := range grids {
		playerIds = append(playerIds, v.GetPlayerIDs()...)
		//fmt.Println("_____>gridi")
	}
	return
}

//添加一个player到格子中
func (m *AOIManager) AddPidToGrid(pid, gid int) {
	m.Grid[gid].Add(pid)
}

//移除一个格子中的PlayerId
func (m *AOIManager) RemovePidFromGrid(pid, gid int) {
	m.Grid[gid].Remove(pid)
}

//通过GID获取全部的playerID
func (m *AOIManager) GetPidByGid(gid int) (playerIds []int) {
	playerIds = m.Grid[gid].GetPlayerIDs()
	return
}

//通过坐标把player添加到一个格子中
func (m *AOIManager) AddToGridByPos(pid int, x, y float32) {
	Gid := m.GetGidsByPos(x, y)
	grid := m.Grid[Gid]
	grid.Add(pid)
}

//通过坐标把一个Player从一个格子中删除
func (m *AOIManager) RemoveFromGridByPos(pid int, x, y float32) {
	Gid := m.GetGidsByPos(x, y)
	grid := m.Grid[Gid]
	grid.Remove(pid)
}
