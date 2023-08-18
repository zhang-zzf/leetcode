package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when434_thenSuccess(t *testing.T) {
	assert.Equal(t, 5, countSegments("Hello, my name is John"))
	assert.Equal(t, 5, countSegments("Hello, my name is John  "))
}

//leetcode submit region begin(Prohibit modification and deletion)
func countSegments(s string) int {
	ans := 0
	word := false
	for _, char := range s {
		if ' ' == char {
			if word {
				ans += 1
				word = false
			}
		} else {
			word = true
		}
	}
	if word {
		ans += 1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
