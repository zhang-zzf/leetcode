package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when387_thenSuccess(t *testing.T) {
	assert.Equal(t, 2, firstUniqChar("loveleetcode"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func firstUniqChar(s string) int {
	ans := -1
	var cntMap [26]int
	for _, char := range s {
		cntMap[char-'a'] += 1
	}
	for idx, char := range s {
		if cntMap[char-'a'] == 1 {
			ans = idx
			break
		}
	}
	return ans
}

func firstUniqChar1(s string) int {
	ans := -1
	cntMap := map[rune]int{}
	for _, char := range s {
		cntMap[char] += 1
	}
	for idx, char := range s {
		if cntMap[char] == 1 {
			ans = idx
			break
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
