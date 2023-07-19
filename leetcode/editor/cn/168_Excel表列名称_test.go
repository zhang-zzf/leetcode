package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when168_thenSuccess(t *testing.T) {
	assert.Equal(t, "ZY", convertToTitle(701), "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func convertToTitle(columnNumber int) string {
	var runes []rune
	for columnNumber != 0 {
		idx := columnNumber % 26
		r := rune('A' - 1 + idx)
		columnNumber /= 26
		if idx == 0 {
			r = 'Z'
			columnNumber -= 1
		}
		runes = append([]rune{r}, runes...)
	}
	return string(runes)
}

//leetcode submit region end(Prohibit modification and deletion)
