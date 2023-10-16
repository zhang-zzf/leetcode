package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1700_thenSuccess(t *testing.T) {
	ans := countStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1})
	assert.Equal(t, 3, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func countStudents(students []int, sandwiches []int) int {
	ans := len(students)
	studentFavorites := make([]int, 2)
	for _, s := range students {
		studentFavorites[s] += 1
	}
	for _, s := range sandwiches {
		if studentFavorites[s] > 0 {
			studentFavorites[s] -= 1
			ans -= 1
		} else {
			break
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
