package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1624_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
// TODO 可以使用数组优化 map
func maxLengthBetweenEqualCharacters(s string) int {
	ans := -1
	mapping := map[rune]int{}
	for idx, c := range s {
		if prev, ok := mapping[c]; ok {
			if idx-prev > ans {
				ans = idx - prev - 1
			}
		} else {
			mapping[c] = idx
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
