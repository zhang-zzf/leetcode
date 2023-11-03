package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when2103_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, countPoints("B0B6G0R6R0R6G9"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func countPoints(rings string) int {
	cnt := [10]int{}
	success := 0
	success |= 1 << ('R' - 'A')
	success |= 1 << ('G' - 'A')
	success |= 1 << ('B' - 'A')
	for i := 1; i < len(rings); i += 2 {
		cnt[rings[i]-'0'] |= 1 << (rings[i-1] - 'A')
	}

	ans := 0
	for _, c := range cnt {
		if c == success {
			ans += 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
