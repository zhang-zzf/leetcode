package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_givenNormal_when884_thenSuccess(t *testing.T) {
	ans := uncommonFromSentences("this apple is sweet", "this apple is sour")
	assert.ElementsMatch(t, []string{"sweet", "sour"}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func uncommonFromSentences(s1 string, s2 string) []string {
	var ans []string
	wordCnt := map[string]int{}
	countWord := func(s string) {
		for _, w := range strings.Split(s, " ") {
			wordCnt[w] += 1
		}
	}
	countWord(s1)
	countWord(s2)
	for w, c := range wordCnt {
		if c == 1 {
			ans = append(ans, w)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
