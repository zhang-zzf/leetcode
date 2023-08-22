package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when605_thenSuccess(t *testing.T) {
	assert.Equal(t, true, canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
}

func Test_givenFailedCase1_when605_thenSuccess(t *testing.T) {
	assert.Equal(t, true, canPlaceFlowers([]int{0, 0, 1, 0, 1}, 1))
}

//leetcode submit region begin(Prohibit modification and deletion)
func canPlaceFlowers(flowerbed []int, n int) bool {
	ans := 0
	lng := len(flowerbed)
	for i := 0; i < lng; i++ {
		// i 前后都是0
		if flowerbed[i] == 0 && (i-1 < 0 || flowerbed[i-1] == 0) && (i+1 >= lng || flowerbed[i+1] == 0) {
			ans += 1
			flowerbed[i] = 1
			// 下一个一定不可以
			i += 1
		}
	}
	return ans >= n
}

// failed
// no consider border [0,0,1,0,1]
func canPlaceFlowers1(flowerbed []int, n int) bool {
	ans := 0
	zeroCnt := 0
	for _, v := range flowerbed {
		if v == 1 {
			zeroCnt = 0
			continue
		}
		zeroCnt += 1
		if zeroCnt == 3 {
			ans += 1
			zeroCnt = 1
		}
	}
	return ans >= n
}

//leetcode submit region end(Prohibit modification and deletion)
