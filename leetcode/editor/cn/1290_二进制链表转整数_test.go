package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when1290_thenSuccess(t *testing.T) {
	ans := getDecimalValue(NewList([]int{1, 0, 1}))
	assert.Equal(t, 5, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	n := 0
	for ptr := head; ptr != nil; ptr = ptr.Next {
		n += 1
	}
	ans := 0
	for ptr := head; ptr != nil; ptr, n = ptr.Next, n-1 {
		if ptr.Val == 1 {
			ans += int(math.Pow(2, float64(n-1)))
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
