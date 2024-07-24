package main

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)

	for _, val := range asteroids {
		if val > 0 {
			stack = append(stack, val)
			continue
		}

		survived := true

		for i := len(stack) - 1; i >= 0; i-- {
			if stack[i] < 0 {
				continue
			}

			if stack[i] > val*-1 {
				survived = false
				break
			}

			if stack[i] == val*-1 {
				survived = false
				stack = stack[0 : len(stack)-1]
				break
			}

			stack = stack[0 : len(stack)-1]
		}

		if survived {
			stack = append(stack, val)
		}
	}

	return stack
}

func main() {
	_ = asteroidCollision([]int{8, -8})
}
