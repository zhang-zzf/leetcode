package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func Test_givenNormal_when914_thenSuccess(t *testing.T) {
	gcd := new(big.Int).
		GCD(nil, nil, big.NewInt(6), big.NewInt(4)).
		Int64()
	assert.Equal(t, int64(2), gcd)
}

//leetcode submit region begin(Prohibit modification and deletion)
// TODO 参考答案
func hasGroupsSizeX(deck []int) bool {
	var gcd func(int, int) int
	gcd = func(a, b int) int {
		if b == 0 {
			return a
		}
		return gcd(b, a%b)
	}
	cardCnt := map[int]int{}
	for _, c := range deck {
		cardCnt[c] += 1
	}
	cd := cardCnt[deck[0]]
	for _, cnt := range cardCnt {
		cd = gcd(cd, cnt)
		if cd < 2 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
