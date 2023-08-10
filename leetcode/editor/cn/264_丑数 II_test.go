package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when264_thenSuccess(t *testing.T) {
	assert.Equal(t, 12, nthUglyNumber(10))
}

//leetcode submit region begin(Prohibit modification and deletion)
func nthUglyNumber(n int) int {
	// TODO PriorityQueue
	ugly := make([]int, n+1)
	ugly[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		n2, n3, n5 := ugly[p2]*2, ugly[p3]*3, ugly[p5]*5
		nextUgly := nthUglyNumberMin(n2, n3, n5)
		ugly[i] = nextUgly
		if nextUgly == n2 {
			p2 += 1
		}
		if nextUgly == n3 {
			p3 += 1
		}
		if nextUgly == n5 {
			p5 += 1
		}
	}
	return ugly[n]
}

func nthUglyNumberMin(args ...int) int {
	ans := args[0]
	for i := 0; i < len(args); i++ {
		if args[i] < ans {
			ans = args[i]
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
