package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1534_thenSuccess(t *testing.T) {
	ans := countGoodTriplets([]int{3, 0, 1, 1, 9, 7}, 7, 2, 3)
	assert.Equal(t, 4, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func countGoodTriplets(arr []int, a int, b int, c int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		} else {
			return x
		}
	}
	ans := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if abs(arr[i]-arr[j]) > a {
				continue
			}
			for k := j + 1; k < n; k++ {
				if abs(arr[j]-arr[k]) > b {
					continue
				}
				if abs(arr[i]-arr[k]) <= c {
					ans += 1
				}
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
