package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1128_thenSuccess(t *testing.T) {
	ans := numEquivDominoPairs([][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}})
	assert.Equal(t, 1, ans)
}

func Test_givenFailedCase1_when1128_thenSuccess(t *testing.T) {
	ans := numEquivDominoPairs([][]int{{1, 4}, {4, 1}, {2, 2}, {1, 4}})
	assert.Equal(t, 3, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func numEquivDominoPairs(dominoes [][]int) int {
	ans := 0
	sumCnt := map[int]int{}
	for _, d := range dominoes {
		n1, n2 := d[0], d[1]
		if n1 > n2 {
			n1, n2 = n2, n1
		}
		n := n1*10 + n2
		ans += sumCnt[n]
		sumCnt[n] += 1
		// TODO error
		//if sumCnt[sum] > 1 {
		//	ans += 1
		//}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
