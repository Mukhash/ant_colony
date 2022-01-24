package lemin

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Lemin(colony *Colony) {
	paths := bfs(colony)
	disjointPaths, err := choosePaths(colony, paths)
	if err != nil {
		log.Fatal(err)
	}
	resPaths := make([][]string, len(disjointPaths)-1)
	for i := 0; i < len(disjointPaths)-1; i++ {
		resPaths[i] = paths[disjointPaths[i]]
	}
	pathStates := distributeAnts(resPaths, colony.numberOfAnts)
	resMap := setOFF(pathStates)
	printTicks(resMap, colony.numberOfAnts)
}

func PrintGraph(col *Colony) {
	m := col.graph
	for k, v := range m {
		fmt.Print(k, ": ")
		fmt.Println(v)
	}
}

func ValidateFile(fileContent []byte) bool {
	countLines := 0
	for i := 0; i < len(fileContent); i++ {
		if fileContent[i] == '\n' {
			countLines++
		}
	}
	if countLines < 5 {
		return false
	}
	return true
}

func PopulateColony(colonyFile []byte, buffer *Colony) error {
	strArr := strings.Split(string(colonyFile), "\n")
	numOfAnts, err := strconv.Atoi(strArr[0])
	if err != nil {
		return err
	}
	buffer.numberOfAnts = numOfAnts
	offset, err := populateRooms(strArr, buffer)
	if err != nil {
		return err
	}
	if err = getGraph(strArr, offset, buffer); err != nil {
		return err
	}
	return nil
}

func CheckColony(colony *Colony) error {
	// Check if number of ants is OK
	if colony.numberOfAnts <= 0 {
		return errors.New("Invalid number of ants")
	}
	// Check if Start & End room exist
	if colony.endRoom.name == "" || colony.StartRoom.name == "" {
		return errors.New("No start or end room")
	}
	return nil
}

func isNotVisited(room string, path []string) bool {
	for _, mroom := range path {
		if room == mroom {
			return false
		}
	}
	return true
}

func contains(s string, arr []string) bool {
	if isNotVisited(s, arr) {
		return false
	}
	return true
}

func printTicks(resMap map[int]map[int]string, ant int) {
	first := true
	for i := 1; i <= len(resMap); i++ {
		for j := 1; j <= ant; j++ {
			if _, b := resMap[i][j]; b {
				if !first {
					fmt.Print(" ")
				}
				first = false
				fmt.Print(resMap[i][j])
			}
		}
		first = true
		fmt.Println()
	}
}

func iterateRooms(ticks map[int]map[int]string, current *roomState, ants []int) map[int]map[int]string {
	res := ticks
	if current.next == nil {

		res[1] = make(map[int]string)
		for _, ant := range ants {
			res[1][ant] = "L" + strconv.Itoa(ant) + "-" + current.name
		}
		return res
	}
	initStep := 1
	for current != nil {
		temp := initStep
		for _, ant := range ants {
			if res[temp] == nil {
				res[temp] = make(map[int]string)
			}
			res[temp][ant] = "L" + strconv.Itoa(ant) + "-" + current.name
			temp++
		}
		initStep++
		current = current.next
	}
	return res
}
