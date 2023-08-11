package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1201_thenSuccess(t *testing.T) {
	assert.Equal(t, 4, nthUglyNumber(3, 2, 3, 5))
	assert.Equal(t, 6, nthUglyNumber(4, 2, 3, 4))
}

func Test_givenFailedCase1_when1201_thenSuccess(t *testing.T) {
	assert.Equal(t, 10, nthUglyNumber(5, 2, 11, 13))
}

func Test_givenFailedCase2_when1201_thenSuccess(t *testing.T) {
	assert.Equal(t, 1999999984, nthUglyNumber(1000000000, 2, 217983653, 336916467))
}

//leetcode submit region begin(Prohibit modification and deletion)
func nthUglyNumber(n int, a int, b int, c int) int {
	ugly := 0
	pa, pb, pc := 1, 1, 1
	for i := 1; i < n+1; i++ {
		na, nb, nc := pa*a, pb*b, pc*c
		ugly = nthUglyNumberMin1201(na, nb, nc)
		if ugly == na {
			pa += 1
		}
		if ugly == nb {
			pb += 1
		}
		if ugly == nc {
			pc += 1
		}
	}
	return ugly
}

func nthUglyNumberMin1201(args ...int) int {
	ans := args[0]
	for i := 0; i < len(args); i++ {
		if args[i] < ans {
			ans = args[i]
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
