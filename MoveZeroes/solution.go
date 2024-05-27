package movezeroes

func moveZeroes(nums []int) {
	for currPtr := 0; currPtr < len(nums); currPtr++ {
		if nums[currPtr] != 0 {
			continue
		}
		for searchPtr := currPtr; searchPtr < len(nums); searchPtr++ {
			if nums[searchPtr] != 0 {
				nums[currPtr] = nums[searchPtr]
				nums[searchPtr] = 0
				break
			}
		}
	}
}
