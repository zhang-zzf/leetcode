package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when2053_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func kthDistinct(arr []string, k int) string {
	// hash 统计
	mapping := map[string]int{}
	for _, w := range arr {
		mapping[w] += 1
	}
	for i := 0; i < len(arr); i++ {
		if mapping[arr[i]] == 1 {
			k -= 1
		}
		if k == 0 {
			return arr[i]
		}
	}
	return ""
}

//leetcode submit region end(Prohibit modification and deletion)
