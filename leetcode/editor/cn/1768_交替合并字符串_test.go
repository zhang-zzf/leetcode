package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1768_thenSuccess(t *testing.T) {
	assert.Equal(t, "apbqcr", mergeAlternately("abc", "pqr"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func mergeAlternately(word1 string, word2 string) string {
	var ans []byte
	strArr := []string{word1, word2}
	n := len(strArr)
	ptrs := make([]int, n)
	endFlag := make([]int, n)
	endCnt := 0
	for i := 0; ; i++ {
		if endCnt >= n {
			break
		}
		idx := i % n
		ptr := ptrs[idx]
		if ptr < len(strArr[idx]) {
			ans = append(ans, strArr[idx][ptr])
			ptrs[idx] += 1
		} else if endFlag[idx] == 0 {
			endCnt += 1
			endFlag[idx] = 1
		}
	}
	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
