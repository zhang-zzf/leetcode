package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when551_thenSuccess(t *testing.T) {
	assert.Equal(t, true, checkRecord("PPALLP"))
	assert.Equal(t, false, checkRecord("PPALLL"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func checkRecord(s string) bool {
	ans := true
	bytes := []byte(s)
	cntA, cntP := 0, 0
	for i := 0; i < len(bytes); i++ {
		if 'A' == bytes[i] {
			cntA += 1
		}
		if 'L' == bytes[i] {
			cntP += 1
		} else {
			cntP = 0
		}
		if cntA >= 2 || cntP >= 3 {
			ans = false
			break
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
