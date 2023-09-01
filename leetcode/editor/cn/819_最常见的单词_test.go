package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"strings"
	"testing"
)

func Test_givenNormal_when819_thenSuccess(t *testing.T) {
	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	ans := mostCommonWord(paragraph, []string{"hit"})
	assert.Equal(t, "ball", ans)
}

// 边界条件
func Test_givenFailedCase1_when819_thenSuccess(t *testing.T) {
	paragraph := "Bob"
	ans := mostCommonWord(paragraph, []string{})
	assert.Equal(t, "bob", ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func mostCommonWord(paragraph string, banned []string) string {
	ans, maxCnt := "", 0
	cntMap := map[string]int{}
	for _, w := range banned {
		cntMap[w] = math.MinInt
	}
	word := strings.Builder{}
	// for _, c := range paragraph {
	// 解决边界条件
	for _, c := range paragraph + " " {
		if c >= 'a' && c <= 'z' {
			word.WriteRune(c)
		} else if c >= 'A' && c <= 'Z' {
			word.WriteRune(c + 'a' - 'A') // 'A' -> 'a'
		} else {
			if word.Len() != 0 {
				w := word.String()
				cntMap[w] += 1
				if cntMap[w] > maxCnt {
					ans, maxCnt = w, cntMap[w]
				}
			}
			// create a new  word
			word = strings.Builder{}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
