package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1201_thenSuccess(t *testing.T) {
	assert.Equal(t, 4, nthUglyNumber(3, 2, 3, 5))
	assert.Equal(t, 6, nthUglyNumber(4, 2, 3, 4))
}

func Test_givenFailedCase1_when1201_thenSuccess(t *testing.T) {
	assert.Equal(t, 10, nthUglyNumber(5, 2, 11, 13))
}

func Test_givenFailedCase2_when1201_thenSuccess(t *testing.T) {
	assert.Equal(t, 1999999984, nthUglyNumber(1000000000, 2, 217983653, 336916467))
}

func Test_givenGcd_when1201_thenSuccess(t *testing.T) {
	assert.Equal(t, 7, gcd1201(98, 63))
	assert.Equal(t, 7, gcd120101(98, 63))
	assert.Equal(t, 52, gcd120101(260, 104))
	assert.Equal(t, 52, gcd1201(260, 104))
}

//leetcode submit region begin(Prohibit modification and deletion)
func nthUglyNumber(n int, a int, b int, c int) int {
	ans := 0
	// TODO 值域2分
	left, right := 0, nthUglyNumberMin1201(a, b, c)*n
	ab, ac, bc := lcm1201(a, b), lcm1201(a, c), lcm1201(b, c)
	abc := lcm1201(ab, c)
	for left <= right {
		mid := left + (right-left)>>1
		num := mid/a + mid/b + mid/c - mid/ab - mid/ac - mid/bc + mid/abc
		if num == n {
			if mid%a == 0 || mid%b == 0 || mid%c == 0 {
				ans = mid
				break
			}
			right = mid - 1
		} else if num > n {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

func nthUglyNumberMin1201(args ...int) int {
	ans := args[0]
	for i := 0; i < len(args); i++ {
		if args[i] < ans {
			ans = args[i]
		}
	}
	return ans
}

// TODO  lcm gcd 都是针对正整数
/**
最小公倍数

least common multiple
*/
func lcm1201(n1, n2 int) int {
	return n1 * n2 / gcd1201(n1, n2)
}

/**
最大公约数

greatest common divisor

highest common factor
*/
// TODO 欧几里得 辗转相除法
func gcd1201(n1, n2 int) int {
	if n2 == 0 {
		return n1
	}
	return gcd1201(n2, n1%n2)
}

// TODO 算术九章 更相减损术
func gcd120101(n1, n2 int) int {
	if n1 < n2 {
		n1, n2 = n2, n1
	}
	subtract := n1 - n2
	if n2 == subtract {
		return subtract
	}
	return gcd120101(n2, subtract)
}

//leetcode submit region end(Prohibit modification and deletion)
