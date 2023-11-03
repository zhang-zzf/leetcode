package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when2094_thenSuccess(t *testing.T) {
	ans := findEvenNumbers([]int{2, 1, 3, 0})
	assert.Equal(t, []int{102, 120, 130, 132, 210, 230, 302, 310, 312, 320}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func findEvenNumbers(digits []int) []int {
	// 逆行思考，3位数 100～999，偶数，且数可以由digits组成
	cnt := [10]int{}
	for _, d := range digits {
		cnt[d] += 1
	}
	// todo slice 是引用类型（copy 引用传递）， array 是基本类型（copy 数组传递）
	checkNum := func(nums [10]int, n int) bool {
		for n > 0 {
			d := n % 10
			nums[d] -= 1
			if nums[d] < 0 {
				break
			}
			n /= 10
		}
		return n == 0
	}
	ans := make([]int, 0)
	for i := 100; i < 1000; i += 2 {
		if checkNum(cnt, i) {
			ans = append(ans, i)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
