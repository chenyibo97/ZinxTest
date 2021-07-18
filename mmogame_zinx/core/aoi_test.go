package core

import (
	"fmt"
	"testing"
)

/*func TestNewAOIManager(t *testing.T) {
	aoimgr:=NewAOIManager(100,300,4,200,450,5)
	fmt.Println(aoimgr)
}*/
func TestAOIManager_String(t *testing.T) {
	aoimgr := NewAOIManager(100, 300, 4, 200, 450, 5)
	defer func() {

		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	for gid, _ := range aoimgr.Grid {

		grids := aoimgr.GetSurroundGridByGid(gid)
		fmt.Println("gid:", gid, "grids len=", len(grids))
		fmt.Println(grids)
		gIDs := make([]int, 0)
		for _, grid := range grids {
			gIDs = append(gIDs, grid.Gid)
		}

		fmt.Println("surroutnd grid ids are:", gIDs)
	}
}
