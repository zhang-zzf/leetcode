package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when925_thenSuccess(t *testing.T) {
	assert.Equal(t, true, isLongPressedName("alex", "aaleex"))
	assert.Equal(t, true, isLongPressedName("saeed", "ssaaeeedd"))
	assert.Equal(t, false, isLongPressedName("saeed", "ssaaedd"))
	//assert.Equal(t, false, isLongPressedName("saeed", "ssaaeeeddm"))
	assert.Equal(t, false, isLongPressedName("saeedd", "ssaaeed"))
	assert.Equal(t, false, isLongPressedName("saeed", "ssaaedd"))
	assert.Equal(t, false, isLongPressedName("alex", "aaaa"))
}

//leetcode submit region begin(Prohibit modification and deletion)
// TODO 参考答案
func isLongPressedName(name string, typed string) bool {
	if name[0] != typed[0] {
		return false
	}
	i, j, m, n := 0, 0, len(name), len(typed)
	for j < n {
		if i < m && typed[j] == name[i] {
			i, j = i+1, j+1
		} else if typed[j] == typed[j-1] {
			j += 1
		} else {
			break
		}
	}
	return i == m
}

// 算法比较复杂，可以应付比原题目更多的场景
func isLongPressedName1(name string, typed string) bool {
	if name[0] != typed[0] {
		return false
	}
	i, j, m, n := 0, 0, len(name), len(typed)
	for i < m && j < n {
		for i < m && j < n && typed[j] == name[i] {
			i, j = i+1, j+1
		}
		for j < n && typed[j] == name[i-1] {
			j += 1
		}
		if i < m && j < n && typed[j] != name[i] {
			break
		}
	}
	return i == m && j == n
}

//leetcode submit region end(Prohibit modification and deletion)
