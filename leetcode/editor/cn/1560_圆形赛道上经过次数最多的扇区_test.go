package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when1560_thenSuccess(t *testing.T) {
	assert.Equal(t, []int{1, 2}, mostVisited(4, []int{1, 3, 1, 2}))
	assert.Equal(t, []int{2}, mostVisited(2, []int{2, 1, 2, 1, 2, 1, 2, 1, 2}))
}

//leetcode submit region begin(Prohibit modification and deletion)

//  TODO 参考答案
func mostVisited(n int, rounds []int) (ans []int) {
	s, e := rounds[0], rounds[len(rounds)-1]
	if s <= e {
		for i := s; i <= e; i++ {
			ans = append(ans, i)
		}
	} else {
		for i := 1; i <= e; i++ {
			ans = append(ans, i)
		}
		for i := s; i <= n; i++ {
			ans = append(ans, i)
		}
	}
	return ans
}

//
func mostVisited1(n int, rounds []int) []int {
	ptr, visited, most := rounds[0]-1, make([]int, n+1), math.MinInt
	for i := 0; i < len(rounds); i++ {
		for ptr != rounds[i] {
			ptr = (ptr + 1) % n
			if ptr == 0 {
				ptr = n
			}
			visited[ptr] += 1
			if visited[ptr] > most {
				most = visited[ptr]
			}
		}
	}
	var ans []int
	for i := 1; i <= n; i++ {
		if visited[i] == most {
			ans = append(ans, i)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
