package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when242_thenSuccess(t *testing.T) {
	assert.Equal(t, true, isAnagram("anagram", "nagaram"))
	assert.Equal(t, false, isAnagram("anagrm", "nagaram"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	var cnt [26]int
	for _, c := range s {
		cnt[c-'a'] += 1
	}
	for _, c := range t {
		cnt[c-'a'] -= 1
	}
	for _, c := range cnt {
		if c != 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
