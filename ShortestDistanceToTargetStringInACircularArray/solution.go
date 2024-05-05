package shortestdistancetotargetstringinacirculararray

// https://leetcode.com/problems/shortest-distance-to-target-string-in-a-circular-array/submissions/1241280707

func closetTarget(words []string, target string, startIndex int) int {
	i1 := startIndex
	i2 := startIndex
	counter := 0

	for {
		if words[i1] == target || words[i2] == target {
			return counter
		}

		i1--
		i2++

		if i1 == -1 {
			i1 = len(words) - 1
		}

		if i2 == len(words) {
			i2 = 0
		}

		counter++

		if counter == len(words) {
			break
		}
	}

	return -1
}
