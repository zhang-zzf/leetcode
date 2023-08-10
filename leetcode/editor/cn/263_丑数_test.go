package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when263_thenSuccess(t *testing.T) {
	// var aInt32 int32
	// The value of '-math.MinInt32' (type untyped int) cannot be represented by the type int32
	// aInt32 = -math.MinInt32
	var aVar int64
	aVar = -math.MinInt32
	assert.Equal(t, int64(1<<31), aVar)
	assert.Equal(t, true, isUgly(10))
	assert.Equal(t, false, isUgly(14))
}

func Test_givenFailedCase1_when263_thenSuccess(t *testing.T) {
	assert.Equal(t, true, isUgly(1))
}

func Test_givenFailedCase2_when263_thenSuccess(t *testing.T) {
	assert.Equal(t, false, isUgly(0))
}

func Test_givenFailedCase3_when263_thenSuccess(t *testing.T) {
	assert.Equal(t, false, isUgly(-14))
}

func Test_givenFailedCase4_when263_thenSuccess(t *testing.T) {
	assert.Equal(t, false, isUgly(math.MinInt32))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isUgly(n int) bool {
	factors := []int{2, 3, 5}
	for n > 1 {
		ugly := false
		for _, f := range factors {
			if n%f == 0 {
				n /= f
				ugly = true
			}
		}
		if !ugly {
			break
		}
	}
	return n == 1
}

//leetcode submit region end(Prohibit modification and deletion)
