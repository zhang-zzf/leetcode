package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_givenNormal_when682_thenSuccess(t *testing.T) {
	params := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	assert.Equal(t, 27, calPoints(params))
}

//leetcode submit region begin(Prohibit modification and deletion)
func calPoints(operations []string) int {
	var points []int
	for _, c := range operations {
		switch c {
		case "C":
			points = points[:len(points)-1]
		case "D":
			points = append(points, points[len(points)-1]*2)
		case "+":
			points = append(points, points[len(points)-1]+points[len(points)-2])
		default:
			n, _ := strconv.Atoi(c)
			points = append(points, n)
		}
	}
	ans := 0
	for _, x := range points {
		ans += x
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
