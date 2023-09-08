package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when953_thenSuccess(t *testing.T) {
	ans := true
	ans = isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz")
	assert.Equal(t, false, ans)
	ans = isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz")
}
func Test_givenFailedCase1_when953_thenSuccess(t *testing.T) {
	ans := true
	params := []string{"fxasxpc", "dfbdrifhp", "nwzgs", "cmwqriv", "ebulyfyve", "miracx", "sxckdwzv", "dtijzluhts", "wwbmnge", "qmjwymmyox"}
	ans = isAlienSorted(params, "zkgwaverfimqxbnctdplsjyohu")
	assert.Equal(t, false, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func isAlienSorted(wordArr []string, order string) bool {
	charOrder := make([]int, 26)
	for i, b := range order {
		charOrder[b-'a'] = i
	}
	check := func(w1, w2 string) bool {
		i, m, n := 0, len(w1), len(w2)
		for ; i < m && i < n; i++ {
			b := charOrder[w2[i]-'a'] - charOrder[w1[i]-'a']
			if b == 0 {
				continue
			}
			return b > 0
		}
		return i == m
	}
	n := len(wordArr)
	ans := true
	for i := 1; i < n; i++ {
		if !check(wordArr[i-1], wordArr[i]) {
			ans = false
			break
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
