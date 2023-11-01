package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_givenNormal_when2011_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func finalValueAfterOperations(operations []string) int {
	ans := 0
	for _, op := range operations {
		if strings.Index(op, "+") != -1 {
			ans += 1
		} else {
			ans -= 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
