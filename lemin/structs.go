package lemin

type room struct {
	name string
	X    int
	Y    int
}

type Colony struct {
	numberOfAnts    int
	graph           map[string][]string
	rooms           []room
	StartRoom       room
	endRoom         room
	uniqueRoomNames map[string]int
	uniqueRoomDesc  map[int]int
}

func (colony *Colony) Init() {
	colony.graph = make(map[string][]string)
	colony.uniqueRoomNames = make(map[string]int)
	colony.uniqueRoomDesc = make(map[int]int)
}

type pathState struct {
	roomStates *roomState
	ants       []int
	rooms      int
}

type roomState struct {
	name string
	next *roomState
}

type queue [][]string

func (q queue) poll() ([]string, queue) {
	res := q[0]
	return res, q[1:]
}

type pathCombState struct {
	paths        [][]string
	pathCombs    [][]int
	numOfRooms   int
	numOfAnts    int
	minCombTicks int
	minCombIndex int
}
