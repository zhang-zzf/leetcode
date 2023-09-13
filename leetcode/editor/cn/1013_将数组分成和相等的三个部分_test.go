package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1013_thenSuccess(t *testing.T) {
	ans := canThreePartsEqualSum([]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1})
	assert.Equal(t, true, ans)
}

func Test_givenFailedCase1_when1013_thenSuccess(t *testing.T) {
	ans := canThreePartsEqualSum([]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4})
	assert.Equal(t, true, ans)
}

func Test_givenFailedCase2_when1013_thenSuccess(t *testing.T) {
	ans := canThreePartsEqualSum([]int{0, 0, 0, 0})
	assert.Equal(t, true, ans)
}

func Test_givenFailedCase3_when1013_thenSuccess(t *testing.T) {
	ans := canThreePartsEqualSum([]int{1, -1, 1, -1})
	assert.Equal(t, false, ans)
}

func Test_givenFailedCase4_when1013_thenSuccess(t *testing.T) {
	ans := canThreePartsEqualSum([]int{-7, 57, -76, 5, 51, -57, 27, 83, 68, 39,
		-38, 43, -79, 42, -79, 72, -58, -29, -34, 23, 20, -67, -67, 74, 10, -90,
		-91, 41, 22, -12, -75, -3, 51, 13, 9, 58, 72, -3, 48, -28, 78, 15, 13,
		22, -68, 62, 33, -18, 21, -2, 56, 41, 27, -33, -35, -36, -6, -73, -9,
		-45, -16, 4, 51, -54, -92, 79, -92, -12, -9, 39, 1, -60, 78, 16, -77,
		-45, 26, -43, 44, -21, -27, 74, -68, -58, 22, 52, 100, 37, 48, 99, 98,
		-34, 36, 25, -85, 83, 99, -64, -37, 20, 0, 13, 85, 18, 98, -91, 86, -94,
		-92, -42, -82, 89, 67, -13, 43, 11, 75, 67, -38, -79, -3, -24, 83, -1,
		74, -93, 26, -69, 92, 40, -32, -89, -100, -36, 80, 9, 92, 79, -83, 90,
		-8, -77, 1, -60, 100, 63, 87, -51, -33, 17, 24, -73, -81, -86, 52, 59,
		10, -84, -71, -12, 24, -53, -14, -44, -48, -91, -57, -82, -5, 48, -56,
		-92, 83, -64, -17, -46, 65, -15, 87, -61, -29, 79, -53, -81, -58, 70,
		70, -12, -71, 68, -57, -44, -18, -94, -84, 7, -51, 18, 44, 96, 24, 8,
		61, 36, -83, -93, 38, 82, 25, 4, 37, 93, -5, 24, -9, -31, 15, -73, -91,
		-36, 41, 56, -17, 3, -94, -26, 96, 5, -21, -20, 60, -53, 26, 2, -97, 41,
		-72, -29, -83, 32, 35, -29, 82, -38, -15, 81, -10, 50, 22, 26, 30, -31,
		74, -74, 47, 66, 53, 35, -58, -68, 47, -21, 78, 12, -88, 1, 12, 53, 94,
		1, -43, -36, 67, 6, -89, 44, -86, -22, 48, -12, 72, -90, -7, -92, 30,
		-17})
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
// 暴力
func canThreePartsEqualSum(arr []int) bool {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	if sum%3 != 0 {
		return false
	}
	i, ss, sc, s := 0, 0, 0, sum/3
	for i = 0; i < len(arr); i++ {
		ss += arr[i]
		if ss == s {
			sc += 1
			ss = 0
		}
		// 在 [0,i-1) 中找到2份的和 = sum/3 即可；剩余的部分的和一定是 sum/3
		if sc == 2 {
			break
		}
	}
	return i < len(arr)-1
}

// 暴力
// 算法是错误的
func canThreePartsEqualSumError(arr []int) bool {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	if sum%3 != 0 {
		return false
	}
	ss, s, sc := 0, sum/3, 0
	for i := 0; i < len(arr); i++ {
		ss += arr[i]
		// 错误在这里
		if ss > s {
			break
		}
		if ss == s {
			sc += 1
			ss = 0
		}
	}
	return sc == 3
}

//leetcode submit region end(Prohibit modification and deletion)
