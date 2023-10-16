package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when1710_thenSuccess(t *testing.T) {
	ans := maximumUnits([][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10)
	assert.Equal(t, 91, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func maximumUnits(boxTypes [][]int, truckSize int) int {
	var ans int
	sort.Slice(boxTypes, func(i, j int) bool { return boxTypes[i][1] > boxTypes[j][1] })
	for i := 0; truckSize > 0 && i < len(boxTypes); {
		num := boxTypes[i][0]
		if truckSize >= num {
			ans += num * boxTypes[i][1]
			truckSize -= num
			i += 1
		} else {
			ans += truckSize * boxTypes[i][1]
			truckSize = 0
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
