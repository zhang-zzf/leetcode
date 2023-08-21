package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when500_thenSuccess(t *testing.T) {
	params := []string{"Hello", "Alaska", "Dad", "Peace"}
	assert.Equal(t, []string{"Alaska", "Dad"}, findWords(params))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findWords(words []string) []string {
	var ans []string
	rowMap := map[rune]int{}
	for _, c := range "qwertyuiop" {
		rowMap[c] = 0
		rowMap[c+'A'-'a'] = 0
	}
	for _, c := range "asdfghjkl" {
		rowMap[c] = 1
		rowMap[c+'A'-'a'] = 1
	}
	for _, c := range "zxcvbnm" {
		rowMap[c] = 2
		rowMap[c+'A'-'a'] = 2
	}
	for _, w := range words {
		same := true
		for i := 1; i < len(w); i++ {
			if rowMap[rune(w[i])] != rowMap[rune(w[i-1])] {
				same = false
				break
			}
		}
		if same {
			ans = append(ans, w)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
