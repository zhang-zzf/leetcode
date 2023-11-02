package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when2062_thenSuccess(t *testing.T) {
	assert.Equal(t, 7, countVowelSubstrings("cuaieuouac"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func countVowelSubstrings(word string) (ans int) {
	// todo 参考更优解
	var success int
	success |= 1 << int('a'-'a')
	success |= 1 << int('e'-'a')
	success |= 1 << int('i'-'a')
	success |= 1 << int('o'-'a')
	success |= 1 << int('u'-'a')
	n := len(word)
	for i := 0; i < n; i++ {
		state := 0
		for j := i; j < n; j++ {
			bits := 1 << int(word[j]-'a')
			if success&bits == 0 {
				break
			}
			state |= bits
			if state == success {
				ans += 1
			}
		}
	}
	return ans
}

func countVowelSubstrings1(word string) int {
	mapping := map[byte]int{
		'a': 0x1,
		'e': 0x10,
		'i': 0x100,
		'o': 0x1000,
		'u': 0x10000,
	}
	ans := 0
	for left := 0; left < len(word); left++ {
		state := 0
		for right := left; right < len(word); right++ {
			m, ok := mapping[word[right]]
			if !ok {
				break
			}
			state |= m
			if state == 0x11111 {
				ans += 1
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
