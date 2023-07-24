package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when383_thenSuccess(t *testing.T) {
	assert.Equal(t, true, canConstruct("aa", "baa"), "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func canConstruct(ransomNote string, magazine string) bool {
	ans := true
	var mapping [26]int
	for _, r := range magazine {
		mapping[r-'a'] += 1
	}
	for _, r := range ransomNote {
		cnt := mapping[r-'a']
		if cnt < 1 {
			ans = false
			break
		}
		mapping[r-'a'] -= 1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
