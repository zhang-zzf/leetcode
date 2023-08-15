package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when345_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, 1, "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
func reverseVowels(s string) string {
	void := struct{}{}
	vowels := map[byte]struct{}{
		'a': void, 'e': void, 'i': void, 'o': void, 'u': void,
		'A': void, 'E': void, 'I': void, 'O': void, 'U': void,
	}
	// TODO string to bytes
	bytes := []byte(s)
	for left, right := 0, len(bytes)-1; left < right; {
		if _, ok := vowels[bytes[left]]; !ok {
			left += 1
			continue
		}
		if _, ok := vowels[bytes[right]]; !ok {
			right -= 1
			continue
		}
		bytes[left], bytes[right] = bytes[right], bytes[left]
		left, right = left+1, right-1
	}
	return string(bytes)
}

//leetcode submit region end(Prohibit modification and deletion)
