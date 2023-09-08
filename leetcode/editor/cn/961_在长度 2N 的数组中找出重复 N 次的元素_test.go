package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func Test_givenNormal_when961_thenSuccess(t *testing.T) {
	assert.Equal(t, 5, repeatedNTimes([]int{5, 1, 5, 2, 5, 3, 5, 4}))
}

//leetcode submit region begin(Prohibit modification and deletion)
// TODO 参考答案
func repeatedNTimes(nums []int) int {
	n := len(nums)
	for {
		i, j := rand.Intn(n), rand.Intn(n)
		if i != j && nums[i] == nums[j] {
			return nums[i]
		}
	}
}

func repeatedNTimes1(nums []int) int {
	// 基数排序
	numCnt := map[int]int{}
	for _, n := range nums {
		numCnt[n] += 1
		if numCnt[n] > 1 {
			return n
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
