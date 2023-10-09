package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when1588_thenSuccess(t *testing.T) {
	assert.Equal(t, 58, sumOddLengthSubarrays([]int{1, 4, 2, 5, 3}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	ans := 0
	for i := 1; i <= n; i += 2 {
		maxCnt := int(math.Min(float64(n-i+1), float64(i)))
		for j := 0; j < n; j++ {
			distance := int(math.Min(float64(j+1), float64(n-j)))
			ans += int(math.Min(float64(maxCnt), float64(distance))) * arr[j]
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
