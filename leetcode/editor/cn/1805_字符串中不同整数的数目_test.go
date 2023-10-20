package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1805_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, numDifferentIntegers("a1b01c001"))
}

func Test_givenFailedCase1_when1805_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, numDifferentIntegers("a0a0"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func numDifferentIntegers(word string) int {
	mapping := map[string]struct{}{}
	n := len(word)
	for l := 0; l < n; {
		if word[l] > '9' {
			l += 1
			continue
		}
		// todo '0' 可以是前缀/后缀/单独
		if word[l] == '0' {
			if l+1 >= n || word[l+1] > '9' { // 单独
				mapping["0"] = struct{}{}
			}
			l += 1
			continue
		}
		r := l + 1
		for r < n && word[r] <= '9' {
			r += 1
		}
		mapping[word[l:r]] = struct{}{}
		l = r
	}
	return len(mapping)
}

//leetcode submit region end(Prohibit modification and deletion)
