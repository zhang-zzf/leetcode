package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when989_thenSuccess(t *testing.T) {
	ans := addToArrayForm([]int{2, 1, 5}, 806)
	assert.Equal(t, []int{1, 0, 2, 1}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
// TODO 参考答案
func addToArrayForm(num []int, k int) []int {
	var ans []int
	for i := len(num) - 1; i >= 0 || k > 0; i-- {
		if i >= 0 {
			k += num[i]
		}
		ans = append(ans, k%10)
		k /= 10
	}
	for l, r := 0, len(ans)-1; l < r; l, r = l+1, r-1 {
		ans[l], ans[r] = ans[r], ans[l]
	}
	return ans
}

func addToArrayForm1(num []int, k int) []int {
	var ans []int
	ab, sum := 0, 0
	for i := len(num) - 1; i >= 0 || k > 0; i-- {
		sum = 0
		if i >= 0 {
			sum += num[i]
		}
		if k > 0 {
			sum += k % 10
			k /= 10
		}
		sum += ab
		ab, sum = sum/10, sum%10
		ans = append(ans, sum)
	}
	if ab > 0 {
		ans = append(ans, ab)
	}
	for l, r := 0, len(ans)-1; l < r; l, r = l+1, r-1 {
		ans[l], ans[r] = ans[r], ans[l]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
