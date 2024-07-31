package graph

var output []int

func findOrder(numCourses int, prerequisites [][]int) []int {
	output = make([]int, 0)

	if len(prerequisites) == 0 {
		for i := 0; i < numCourses; i++ {
			output = append(output, i)
		}
		return output
	}

	prereqs := make(map[int][]int)

	for _, pair := range prerequisites {
		if _, exists := prereqs[pair[0]]; !exists {
			prereqs[pair[0]] = []int{pair[1]}
		} else {
			prereqs[pair[0]] = append(prereqs[pair[0]], pair[1])
		}
	}

	visited := make(map[int]bool)
	visitedForCurrCourse := make(map[int]bool)

	for course := 0; course < numCourses; course++ {
		if !dfs(course, prereqs, visited, visitedForCurrCourse) {
			return []int{}
		}
	}

	return output
}

func dfs(currCourse int, prereqs map[int][]int, visited map[int]bool, visitedForCurrCourse map[int]bool) bool {
	if _, isVisited := visitedForCurrCourse[currCourse]; isVisited {
		return false
	}

	if _, isVisited := visited[currCourse]; isVisited {
		return true
	}

	visitedForCurrCourse[currCourse] = true

	for _, prereqCourse := range prereqs[currCourse] {
		if !dfs(prereqCourse, prereqs, visited, visitedForCurrCourse) {
			return false
		}
	}

	delete(visitedForCurrCourse, currCourse)
	visited[currCourse] = true
	output = append(output, currCourse)

	return true
}
