package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when1370_thenSuccess(t *testing.T) {
	assert.Equal(t, "abccbaabccba", sortString("aaaabbbbcccc"))
	assert.Equal(t, "ggggggg", sortString("ggggggg"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func sortString(s string) string {
	// 基数统计
	const total = 26
	mapping := [total]int{}
	for _, c := range s {
		mapping[c-'a'] += 1
	}
	ans := make([]byte, len(s))
	for selected := 0; selected < len(s); {
		for c := 0; c < total; c++ {
			if mapping[c] > 0 {
				ans[selected] = byte(c + 'a')
				selected += 1
				mapping[c] -= 1
			}
		}
		for c := total - 1; c >= 0; c-- {
			if mapping[c] > 0 {
				ans[selected] = byte(c + 'a')
				selected += 1
				mapping[c] -= 1
			}
		}
	}
	return string(ans)
}

func sortString1(s string) string {
	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
	n := len(bytes)
	ans := make([]byte, n)
	pivot := byte('a' - 1)
	for selected := 0; selected < n; {
		min, max := 0, 0
		for i := 0; i < n; i++ {
			if bytes[i] == pivot {
				continue
			}
			if min == 0 || bytes[i] != ans[selected-1] {
				ans[selected] = bytes[i]
				bytes[i] = pivot
				selected += 1
				min += 1
			}
		}
		// TODO
		// for i := n - 1; i >= 0 && bytes[i] != pivot; {
		for i := n - 1; i >= 0; i-- {
			if bytes[i] == pivot {
				continue
			}
			if max == 0 || bytes[i] != ans[selected-1] {
				ans[selected] = bytes[i]
				bytes[i] = pivot
				selected += 1
				max += 1
			}
		}
	}
	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
