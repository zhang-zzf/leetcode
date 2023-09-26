package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1436_thenSuccess(t *testing.T) {
	ans := destCity([][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}})
	assert.Equal(t, "Sao Paulo", ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func destCity(paths [][]string) string {
	cities := map[string]int{}
	for _, path := range paths {
		cities[path[0]] += 1
		cities[path[1]] += 0
	}
	for k, v := range cities {
		if v == 0 {
			return k
		}
	}
	return ""
}

//leetcode submit region end(Prohibit modification and deletion)
