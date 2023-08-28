package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when482_thenSuccess(t *testing.T) {
	assert.Equal(t, "5F3Z-2E9W", licenseKeyFormatting("5F3Z-2e-9-w", 4))
}

/**
"--a-a-a-a--"
*/
func Test_givenFailedCase1_when482_thenSuccess(t *testing.T) {
	assert.Equal(t, "AA-AA", licenseKeyFormatting("--a-a-a-a--", 2))
}

//leetcode submit region begin(Prohibit modification and deletion)

/**
licenseKeyFormatting1 是在数组0位置插入元素
licenseKeyFormatting 是在数组末尾插入元素，然后反转数组，性能提升 n 倍
Runtime:0 ms, faster than 100.00% of Go online submissions.
*/
func licenseKeyFormatting(s string, k int) string {
	var chars []byte
	ck := 0
	for i := len(s) - 1; i >= 0; i-- {
		// TODO 注意 判断顺序
		c := s[i]
		if c == '-' {
			continue
		}
		if ck == k {
			chars = append(chars, '-')
			ck = 0
		}
		if c >= 'a' && c <= 'z' {
			c = c + 'A' - 'a'
		}
		chars = append(chars, c)
		ck += 1
	}
	// reverse
	for l, r := 0, len(chars)-1; l <= r; l, r = l+1, r-1 {
		chars[l], chars[r] = chars[r], chars[l]
	}
	return string(chars)
}

/**
Runtime:268 ms, faster than 9.52% of Go online submissions.
*/
func licenseKeyFormatting1(s string, k int) string {
	// s = strings.ToUpper(s)
	var chars []byte
	ck := 0
	for i := len(s) - 1; i >= 0; i-- {
		// 注意 判断顺序
		c := s[i]
		if c == '-' {
			continue
		}
		if ck == k {
			chars = append([]byte{'-'}, chars...)
			ck = 0
		}
		if c >= 'a' && c <= 'z' {
			c = c + 'A' - 'a'
		}
		// TODO go 数组 0 位置插入
		chars = append([]byte{c}, chars...)
		ck += 1
	}
	return string(chars)
}

//leetcode submit region end(Prohibit modification and deletion)
