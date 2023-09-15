package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func Test_givenNormal_when1154_thenSuccess(t *testing.T) {
	assert.Equal(t, 9, dayOfYear("2019-01-09"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func dayOfYear(date string) int {
	daysOfMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	ymd := strings.Split(date, "-")
	y, _ := strconv.Atoi(ymd[0])
	m, _ := strconv.Atoi(ymd[1])
	d, _ := strconv.Atoi(ymd[2])
	ans := 0
	ans += d
	for i := 1; i < m; i++ {
		ans += daysOfMonth[i-1]
	}
	if m > 2 && (y%400 == 0 || (y%4 == 0 && y%100 != 0)) {
		ans += 1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
