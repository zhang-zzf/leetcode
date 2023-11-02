package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when2037_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	ans := 0
	for i := 0; i < len(students); i++ {
		distance := students[i] - seats[i]
		if distance < 0 {
			distance = -distance
		}
		ans += distance
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
