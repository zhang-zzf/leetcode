package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1854_thenSuccess(t *testing.T) {
	ans := maximumPopulation([][]int{{1982, 1998}, {2013, 2042}, {2010, 2035}, {2022, 2050}, {2047, 2048}})
	assert.Equal(t, 2022, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func maximumPopulation(logs [][]int) int {
	year := make([]int, 101)
	for _, log := range logs {
		for i := log[0]; i < log[1]; i++ {
			year[i-1950] += 1
		}
	}
	ans := 0
	for i := 0; i < 101; i++ {
		if year[i] > year[ans] {
			ans = i
		}
	}
	return ans + 1950
}

//leetcode submit region end(Prohibit modification and deletion)
