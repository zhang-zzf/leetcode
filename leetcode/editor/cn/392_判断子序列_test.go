package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when392_thenSuccess(t *testing.T) {
	assert.Equal(t, true, isSubsequence("abc", "ahbgdc"))
	assert.Equal(t, false, isSubsequence("axc", "ahbgdc"))
}

/**
边界条件
"abc", ""
*/
func Test_givenFailedCase1_when392_thenSuccess(t *testing.T) {
	assert.Equal(t, false, isSubsequence("abc", ""))
	assert.Equal(t, true, isSubsequence("", ""))
}

/**
s 超出 t 部分
*/
func Test_givenFailedCase2_when392_thenSuccess(t *testing.T) {
	assert.Equal(t, false, isSubsequence("acb", "ahbadc"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isSubsequence(s string, t string) bool {
	//如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，
	//你需要依次检查它们是否为 T 的子序列
	// 分析 t 中的字母和顺序
	var dp []map[byte]int
	length := len(t)
	for idx := 0; idx < length; idx++ {
		m := map[byte]int{}
		for i := idx; i < length; i++ {
			if _, ok := m[t[i]]; !ok {
				m[t[i]] = i
			}
		}
		// TODO 声明的是一个 slice
		// dp[idx] = m
		dp = append(dp, m)
	}
	ans := true
	startIdx := 0
	for _, char := range s {
		if startIdx >= len(t) {
			ans = false
			break
		}
		if idx, ok := dp[startIdx][byte(char)]; !ok {
			ans = false
			break
		} else {
			startIdx = idx + 1
		}
	}
	return ans
}

func isSubsequence1(s string, t string) bool {
	// 双指针
	ans := true
	for ptrs, ptrt := 0, 0; ptrs < len(s); {
		if ptrt >= len(t) {
			ans = false
			break
		}
		if s[ptrs] != t[ptrt] {
			ptrt += 1
			continue
		}
		// move both ptr
		ptrs, ptrt = ptrs+1, ptrt+1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
