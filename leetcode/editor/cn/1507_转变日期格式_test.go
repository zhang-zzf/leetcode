package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_givenNormal_when1507_thenSuccess(t *testing.T) {
	assert.Equal(t, "2052-10-20", reformatDate("20th Oct 2052"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func reformatDate(date string) string {
	dayMapping := map[string]string{}
	for i := 1; i < 32; i++ {
		dd := fmt.Sprintf("%02d", i)
		dayMapping[fmt.Sprintf("%dth", i)] = dd
		switch i {
		case 1, 21, 31:
			dayMapping[fmt.Sprintf("%dst", i)] = dd
		case 2, 22:
			dayMapping[fmt.Sprintf("%dnd", i)] = dd
		case 3, 23:
			dayMapping[fmt.Sprintf("%drd", i)] = dd
		}
	}
	monthMapping := map[string]string{
		"Jan": "01", "Feb": "02", "Mar": "03",
		"Apr": "04", "May": "05", "Jun": "06",
		"Jul": "07", "Aug": "08", "Sep": "09",
		"Oct": "10", "Nov": "11", "Dec": "12"}
	d := strings.Split(date, " ")
	return fmt.Sprintf("%s-%s-%s", d[2], monthMapping[d[1]], dayMapping[d[0]])
}

//leetcode submit region end(Prohibit modification and deletion)
