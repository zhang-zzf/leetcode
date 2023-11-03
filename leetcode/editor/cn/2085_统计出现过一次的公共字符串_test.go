package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when2085_thenSuccess(t *testing.T) {
	words1 := []string{"leetcode", "is", "amazing", "as", "is"}
	words2 := []string{"amazing", "leetcode", "is"}
	assert.Equal(t, 2, countWords(words1, words2))
}

//leetcode submit region begin(Prohibit modification and deletion)
func countWords(words1 []string, words2 []string) int {
	mapping := map[string]int{}
	for _, w := range words1 {
		mapping[w] += 1
	}
	for w, c := range mapping {
		if c != 1 {
			mapping[w] = math.MaxInt
		}
	}
	for _, w := range words2 {
		mapping[w] -= 1
	}
	ans := 0
	for _, c := range mapping {
		if c == 0 {
			ans += 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
