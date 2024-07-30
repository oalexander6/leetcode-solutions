package graph

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	prereqs := make(map[int][]int)

	for _, pair := range prerequisites {
		if _, exists := prereqs[pair[0]]; !exists {
			prereqs[pair[0]] = []int{pair[1]}
		} else {
			prereqs[pair[0]] = append(prereqs[pair[0]], pair[1])
		}
	}

	canTakeCount := 0
	visited := make(map[int]bool)

	for course := 0; course < numCourses; course++ {
		if dfs(course, prereqs, visited) {
			canTakeCount++
		}

		if canTakeCount >= numCourses {
			return true
		}
	}

	return canTakeCount >= numCourses
}

func dfs(currCourse int, prereqs map[int][]int, visited map[int]bool) bool {
	if _, isVisited := visited[currCourse]; isVisited {
		return false
	}

	if _, hasPrereqs := prereqs[currCourse]; !hasPrereqs {
		return true
	}

	visited[currCourse] = true

	for _, prereqCourse := range prereqs[currCourse] {
		if !dfs(prereqCourse, prereqs, visited) {
			return false
		}
	}

	delete(visited, currCourse)
	delete(prereqs, currCourse)

	return true
}
