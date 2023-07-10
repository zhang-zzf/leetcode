package leetcode

import (
	"container/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when20_thenSuccess(t *testing.T) {
	assert.Equal(t, true, isValid("()"))
	assert.Equal(t, false, isValid("(]"))
	assert.Equal(t, true, isValid("([])"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	ans := true
	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := list.New()
	for _, r := range s {
		// 判断 key 是否存在
		_, exists := pairs[r]
		if exists {
			stack.PushBack(r)
		} else {
			back := stack.Back()
			// go nil
			// back.Value 为 any (interface{})类型，强制转换为 rune 类型
			if back == nil || r != pairs[back.Value.(rune)] {
				ans = false
				break
			}
			stack.Remove(back)
		}
	}
	if stack.Len() != 0 {
		ans = false
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
