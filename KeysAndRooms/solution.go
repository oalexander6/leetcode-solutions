package keysandrooms

type RoomInfo struct {
	HaveKey   bool
	IsVisited bool
}

func canVisitAllRooms(rooms [][]int) bool {
	visitQ := []int{0}
	visited := make(map[int]bool)
	currRoom := 0

	for {
		visited[currRoom] = true

		for i := 0; i < len(rooms[currRoom]); i++ {
			if _, isVisited := visited[rooms[currRoom][i]]; !isVisited {
				visitQ = append(visitQ, rooms[currRoom][i])
			}
		}

		if len(visitQ) == 0 {
			break
		}

		currRoom = visitQ[len(visitQ)-1]
		visitQ = visitQ[:len(visitQ)-1]
	}

	return len(visited) == len(rooms)
}
