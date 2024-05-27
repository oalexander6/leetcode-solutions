package findhighestaltitude

func largestAltitude(gain []int) int {
	highestAltitude := 0
	currAlt := 0

	for i := 0; i < len(gain); i++ {
		currAlt += gain[i]

		if currAlt > highestAltitude {
			highestAltitude = currAlt
		}
	}

	return highestAltitude
}
