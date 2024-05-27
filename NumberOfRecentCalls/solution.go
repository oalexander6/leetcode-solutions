package numberofrecentcalls

type RecentCounter struct {
	calls []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		calls: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.calls = append(this.calls, t)

	cutoff := t - 3000
	count := 0
	for i := len(this.calls) - 1; i >= 0; i-- {
		if this.calls[i] >= cutoff {
			count++
		} else {
			break
		}
	}

	return count
}

/**
* Your RecentCounter object will be instantiated and called as such:
* obj := Constructor();
* param_1 := obj.Ping(t);
 */
