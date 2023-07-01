package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when191_thenSuccess(t *testing.T) {
	assert.Equal(t, 2, hammingWeight(3), "shouldEqual")
	assert.Equal(t, 32, hammingWeight(math.MaxUint32), "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func hammingWeight(num uint32) int {
	ans := 0
	for num != 0 {
		num &= num - 1
		ans += 1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
