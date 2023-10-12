package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_givenNormal_when1668_thenSuccess(t *testing.T) {
	assert.Equal(t, 2, maxRepeating("ababc", "ab"))
	assert.Equal(t, 3, maxRepeating("ababcababab", "ab"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxRepeating(sequence string, word string) int {
	ans := 0
	n := len(sequence)
	for left := 0; left < n; left++ {
		if sequence[left] != word[0] {
			continue
		}
		for {
			repeat := strings.Repeat(word, ans+1)
			right := left + len(repeat)
			if right <= n && sequence[left:right] == repeat {
				ans += 1
			} else {
				break
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
