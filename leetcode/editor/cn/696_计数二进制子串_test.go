package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when696_thenSuccess(t *testing.T) {
	assert.Equal(t, 6, countBinarySubstrings("00110011"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func countBinarySubstrings(s string) int {
	ans := 0
	bytes := []byte(s)
	n := len(bytes)
	for i := 1; i < n; i++ {
		left := bytes[i-1]
		right := bytes[i]
		if left == right {
			continue
		}
		for l, r := i-1, i; l >= 0 && r < n; l, r = l-1, r+1 {
			if bytes[l] == left && bytes[r] == right {
				ans += 1
			} else {
				break
			}
		}
	}
	return ans
}

// Time Limit Exceeded
func countBinarySubstrings1(s string) int {
	// left, right
	ans := 0
	bytes := []byte(s)
	for l := 0; l < len(bytes); l++ {
		same := true
		cnt := 1
		for r := l + 1; r < len(bytes); r++ {
			if bytes[r] != bytes[r-1] {
				if same == false {
					break
				}
				same = false
			}
			if same {
				cnt += 1
			} else {
				cnt -= 1
			}
			if cnt == 0 {
				ans += 1
				break
			}
		}

	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
