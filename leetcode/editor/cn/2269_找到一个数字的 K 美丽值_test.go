package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_givenNormal_when2269_thenSuccess(t *testing.T) {
	ans := divisorSubstrings(430043, 2)
	assert.Equal(t, 2, ans, "shouldEqual")
}

func Test_givenZero_when2269_thenSuccess(t *testing.T) {
	val := 0
	// error division by zero
	// val = 100 % 0
	assert.Equal(t, 0, val, "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func divisorSubstrings(num int, k int) int {
	// num to string
	// string to num
	ans := 0
	numStr := strconv.FormatInt(int64(num), 10)
	for i := 0; i <= len(numStr)-k; i++ {
		val, _ := strconv.Atoi(numStr[i : i+k])
		if val == 0 {
			continue
		}
		if num%val == 0 {
			ans += 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
