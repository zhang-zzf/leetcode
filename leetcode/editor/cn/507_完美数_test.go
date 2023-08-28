package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when507_thenSuccess(t *testing.T) {
	ans := checkPerfectNumber(28)
	assert.Equal(t, true, ans)
}

func Test_givenFailedCase1_when507_thenSuccess(t *testing.T) {
	assert.Equal(t, true, checkPerfectNumber(6))
	assert.Equal(t, false, checkPerfectNumber(16))
}

//leetcode submit region begin(Prohibit modification and deletion)
func checkPerfectNumber(num int) bool {
	sum := 1
	// TODO 注意，这里是向下取整
	sqrt := int(math.Sqrt(float64(num)))
	for i := 2; i <= sqrt; i++ {
		if num%i == 0 {
			sum += i
			sum += num / i
		}
	}
	if sqrt*sqrt == num {
		sum -= sqrt
	}
	return sum == num
}

//Time Limit Exceeded Your input: 99999995
func checkPerfectNumber1(num int) bool {
	sum := 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	return sum == num
}

//leetcode submit region end(Prohibit modification and deletion)
