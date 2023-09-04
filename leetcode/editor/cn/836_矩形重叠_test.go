package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when836_thenSuccess(t *testing.T) {
	assert.Equal(t, true, isRectangleOverlap([]int{0, 0, 2, 2}, []int{1, 1, 3, 3}))
	assert.Equal(t, false, isRectangleOverlap([]int{0, 0, 1, 1}, []int{1, 0, 2, 1}))
}

func Test_givenFailedCase1_when836_thenSuccess(t *testing.T) {
	assert.Equal(t, true, isRectangleOverlap([]int{7, 8, 13, 15}, []int{10, 8, 12, 20}))
}

//leetcode submit region begin(Prohibit modification and deletion)

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// TODO 参考了答案
	// 降维
	xOverlap := !(rec1[0] >= rec2[2] || rec1[2] <= rec2[0])
	yOverlap := !(rec1[1] >= rec2[3] || rec1[3] <= rec2[1])
	return xOverlap && yOverlap
}

// 算法是错误的
func isRectangleOverlap11(rec1 []int, rec2 []int) bool {
	// 面积法: 2个矩形的左下角 + 右上角的面积和 < 2 个矩形的面积和 一定存在重合
	area := func(rec []int) int {
		return (rec[2] - rec[0]) * (rec[3] - rec[1])
	}
	// TODO 认知异常
	// target 和 rec1 底层指向同一个数组
	// target := rec1[:]
	target := make([]int, 4)
	copy(target, rec1)
	if target[0] > rec2[0] || target[1] > rec2[1] {
		target[0], target[1] = rec2[0], rec2[1]
	}
	if target[2] < rec2[2] || target[3] < rec2[3] {
		target[2], target[3] = rec2[2], rec2[3]
	}
	return area(target) < area(rec1)+area(rec2)
}

// error 2个矩形完全重合
func isRectangleOverlap1(rec1 []int, rec2 []int) bool {
	inner := func(point []int, rec []int) bool {
		x, y := point[0], point[1]
		if x > rec[0] && x < rec[2] && y > rec[1] && y < rec[3] {
			return true
		}
		return false
	}
	return inner(rec1[:2], rec2) ||
		inner(rec1[2:], rec2) ||
		inner(rec2[:2], rec1) ||
		inner(rec2[2:], rec1)
}

//leetcode submit region end(Prohibit modification and deletion)
