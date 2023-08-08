package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when225_thenSuccess(t *testing.T) {
	stack := Constructor()
	assert.True(t, stack.Empty())
	stack.Push(1)
	stack.Push(2)
	assert.Equal(t, 2, stack.Top())
	assert.Equal(t, 2, stack.Pop())
	assert.False(t, stack.Empty())
	assert.Equal(t, 1, stack.Pop())
	assert.True(t, stack.Empty())
}

//leetcode submit region begin(Prohibit modification and deletion)
type MyStack struct {
	high []int
	low  []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.high = append(this.high, x)
	if len(this.high) > 1 {
		this.low = append(this.low, this.high[0])
		this.high = this.high[1:]
	}
}

func (this *MyStack) Pop() int {
	// poll from header
	ans := this.high[0]
	this.high = this.high[1:]
	// switch queues
	this.high, this.low = this.low, this.high
	for len(this.high) > 1 {
		// poll from high and offer to the low
		this.low = append(this.low, this.high[0])
		this.high = this.high[1:]
	}
	return ans
}

func (this *MyStack) Top() int {
	return this.high[0]
}

func (this *MyStack) Empty() bool {
	return len(this.high) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)
