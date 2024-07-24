package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	numPlanted := 0

	paddedFlowerbed := make([]int, len(flowerbed)+2)
	paddedFlowerbed[0] = 0
	paddedFlowerbed[len(paddedFlowerbed)-1] = 0
	copy(paddedFlowerbed[1:], flowerbed)

	for i := 1; i < len(paddedFlowerbed)-1; i++ {
		if paddedFlowerbed[i] == 0 && paddedFlowerbed[i-1] == 0 && paddedFlowerbed[i+1] == 0 {
			paddedFlowerbed[i] = 1
			numPlanted++
		}
	}

	return numPlanted >= n
}

func main() {
	_ = canPlaceFlowers([]int{1, 0, 0, 0, 1, 0, 0}, 2)
}
