package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"sort"
	"testing"
)

func Test_givenNormal_when645_thenSuccess(t *testing.T) {
	ans := findErrorNums([]int{1, 2, 2, 4})
	assert.Equal(t, []int{2, 3}, ans)
}

func Test_givenFailedCase1_when645_thenSuccess(t *testing.T) {
	params := []int{37, 62, 43, 27, 12, 66, 36, 18, 39, 54, 61, 65, 47, 32, 23, 2, 46, 8, 4, 24, 29, 38, 63, 39, 25, 11, 45, 28, 44, 52, 15, 30, 21, 7, 57, 49, 1, 59, 58, 14, 9, 40, 3, 42, 56, 31, 20, 41, 22, 50, 13, 33, 6, 10, 16, 64, 53, 51, 19, 17, 48, 26, 34, 60, 35, 5}
	sort.Ints(params)
	ans := findErrorNums(params)
	assert.Equal(t, []int{39, 55}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func findErrorNums(nums []int) []int {
	n := len(nums)
	flag := make([]int, n+1)
	for _, n := range nums {
		flag[n] += 1
	}
	ans := make([]int, 2)
	for i := 1; i < n+1; i++ {
		if flag[i] == 2 {
			ans[0] = i
		} else if flag[i] == 0 {
			ans[1] = i
		}
	}
	return ans
}

// #ERROR
// 一组数字积超出 float64 阀值
func findErrorNums1(nums []int) []int {
	// 先找重复的数
	// sum / multiply 数学方法
	s, p, sn, pn := 0, float64(1), 0, float64(1)
	for i, v := range nums {
		s += v
		p *= float64(v)
		sn += i + 1
		pn *= float64(i + 1)
	}

	d := int(math.Round(float64(sn-s) * p / (pn - p)))
	m := d + (sn - s)
	return []int{d, m}
}

//leetcode submit region end(Prohibit modification and deletion)
