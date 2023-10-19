package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1773_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	mapping := map[string]int{}
	switch ruleKey {
	case "type":
		for _, item := range items {
			mapping[item[0]] += 1
		}
	case "color":
		for _, item := range items {
			mapping[item[1]] += 1
		}
	case "name":
		for _, item := range items {
			mapping[item[2]] += 1
		}
	}
	return mapping[ruleValue]
}

//leetcode submit region end(Prohibit modification and deletion)
