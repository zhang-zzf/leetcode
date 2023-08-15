package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when303_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, 1, "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
type NumArray struct {
	nums []int
}

func Constructor303(nums []int) NumArray {
	return NumArray{nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	// 暴力破解
	ans := 0
	for i := left; i <= right; i++ {
		ans += this.nums[i]
	}
	return ans
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
//leetcode submit region end(Prohibit modification and deletion)
