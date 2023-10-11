package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1656_thenSuccess(t *testing.T) {
	os := Constructor(5)
	assert.Equal(t, []string{}, os.Insert(3, "ccccc"))
	assert.Equal(t, []string{"aaaaa"}, os.Insert(1, "aaaaa"))
	assert.Equal(t, []string{"bbbbb", "ccccc"}, os.Insert(2, "bbbbb"))
	assert.Equal(t, []string{}, os.Insert(5, "eeeee"))
	assert.Equal(t, []string{"ddddd", "eeeee"}, os.Insert(4, "ddddd"))
}

//leetcode submit region begin(Prohibit modification and deletion)
type OrderedStream struct {
	stream []string
	ptr    int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{stream: make([]string, n+1), ptr: 1}

}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.stream[idKey] = value
	if idKey == this.ptr {
		var ans []string
		for this.ptr < len(this.stream) {
			str := this.stream[this.ptr]
			if str != "" {
				ans = append(ans, str)
				this.ptr += 1
			} else {
				break
			}
		}
		return ans
	}
	return []string{}
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
