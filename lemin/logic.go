package lemin

// BFS Breadth First Search
func bfs(colony *Colony) [][]string {
	res := [][]string{}
	var queue queue = [][]string{}
	path := []string{}
	path = append(path, colony.StartRoom.name)
	queue = append(queue, path)
	for len(queue) != 0 {
		var queuePath []string
		queuePath, queue = queue.poll()
		last := queuePath[len(queuePath)-1]
		if last == colony.endRoom.name {
			res = append(res, queuePath)
			continue
		}
		adjRooms := colony.graph[last]
		for i := 0; i < len(adjRooms); i++ {
			if isNotVisited(adjRooms[i], queuePath) {
				newPath := make([]string, len(queuePath))
				copy(newPath, queuePath)
				newPath = append(newPath, adjRooms[i])
				queue = append(queue, newPath)
			}
		}
	}
	return res
}

func distributeAnts(paths [][]string, ants int) []pathState {
	pathStates := make([]pathState, len(paths))
	for i, path := range paths {
		pathStates[i] = pathState{
			rooms:      len(path) - 1,
			roomStates: getRoomStates(paths[i]),
		}
	}
	setAnts(pathStates, ants)
	return pathStates
}

func getRoomStates(rooms []string) *roomState {
	arr := make([]roomState, len(rooms))
	arr[len(arr)-1] = roomState{
		name: rooms[len(arr)-1],
		next: nil,
	}
	for i := 0; i < len(rooms)-1; i++ {
		arr[i] = roomState{
			name: rooms[i],
			next: &arr[i+1],
		}
	}
	return &arr[0]
}

func setAnts(ps []pathState, ants int) {
	size := len(ps)
	if size == 1 {
		for i := 1; i <= ants; i++ {
			ps[0].ants = append(ps[0].ants, i)
		}
		return
	}
	i := 0
	cur := 1
	next := getNextIndex(i, size)
	for cur <= ants {
		if ps[i].rooms+len(ps[i].ants) > ps[next].rooms+len(ps[next].ants) {
			ps[next].ants = append(ps[next].ants, cur)
			cur++
			i = getNextIndex(i, size)
			next = getNextIndex(i, size)
			continue
		} else if i == size-1 {
			i = 0
			next = getNextIndex(i, size)
			continue
		}
		ps[i].ants = append(ps[i].ants, cur)
		cur++
	}
}

func getNextIndex(cur int, len int) int {
	if cur+1 == len {
		return 0
	}
	cur++
	return cur
}

func setOFF(paths []pathState) map[int]map[int]string {
	ticks := make(map[int]map[int]string)
	for _, path := range paths {
		current := path.roomStates.next
		ticks = iterateRooms(ticks, current, path.ants)
	}
	return ticks
}
