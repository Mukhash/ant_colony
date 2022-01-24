package lemin

import (
	"errors"
)

func choosePaths(colony *Colony, paths [][]string) ([]int, error) {
	pcs := &pathCombState{
		paths:      paths,
		numOfRooms: len(colony.rooms) - 2,
		numOfAnts:  colony.numberOfAnts}
	for i := 0; i < len(paths); i++ {
		path, isLess := getPathComb(i, pcs)
		pcs.pathCombs = append(pcs.pathCombs, path)
		if isLess {
			pcs.minCombIndex = len(pcs.pathCombs) - 1
		}
	}
	if len(pcs.pathCombs) == 0 {
		return []int{}, errors.New("no paths can be created between start and end")
	}
	return pcs.pathCombs[pcs.minCombIndex], nil
}

func getPathComb(firstPath int, pcs *pathCombState) ([]int, bool) {
	res := []int{}
	visitedRooms := make(map[string]bool)
	pathBusy := false
	paths := pcs.paths
	for i := firstPath; i < len(paths); i++ {
		curPath := paths[i]
		if len(curPath) == 2 {
			return []int{i, 1}, true
		}
		for j := 1; j < len(curPath)-1; j++ {
			if visitedRooms[curPath[j]] {
				pathBusy = true
				break
			}
		}
		if !pathBusy {
			res = append(res, i)
			for _, path := range curPath {
				visitedRooms[path] = true
			}
		}
		pathBusy = false
	}
	ticks, lessTicks := getPathTicks(pcs, res)
	if lessTicks {
		pcs.minCombTicks = ticks
	}
	res = append(res, ticks)
	return res, lessTicks
}

func getPathTicks(pcs *pathCombState, comb []int) (int, bool) {
	numOfRooms := getCombNumRooms(pcs.paths, comb)
	numOfAnts := pcs.numOfAnts
	ticks := (numOfRooms + numOfAnts) / len(comb)
	lessTicks := false
	if pcs.minCombTicks == 0 {
		pcs.minCombTicks = ticks
	} else if pcs.minCombTicks > ticks {
		pcs.minCombTicks = ticks
		lessTicks = true
	}
	return ticks, lessTicks
}

func getCombNumRooms(paths [][]string, comb []int) int {
	res := 0
	for _, pathIndex := range comb {
		res += len(paths[pathIndex])
	}
	return res - len(comb)
}
