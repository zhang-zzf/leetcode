package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when232_thenSuccess(t *testing.T) {
	queue := Constructor()
	assert.Equal(t, true, queue.Empty())
	queue.Push(1)
	queue.Push(2)
	assert.Equal(t, 1, queue.Peek())
	queue.Push(3)
	assert.Equal(t, 1, queue.Pop())
	assert.Equal(t, 2, queue.Pop())
	assert.Equal(t, false, queue.Empty())
	assert.Equal(t, 3, queue.Pop())
	assert.Equal(t, true, queue.Empty())
}

//leetcode submit region begin(Prohibit modification and deletion)
type MyQueue struct {
	head []int
	tail []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.tail = append(this.tail, x)
}

func (this *MyQueue) Pop() int {
	this.tailToHead()
	top := this.head[len(this.head)-1]
	this.head = this.head[:len(this.head)-1]
	return top
}

func (this *MyQueue) tailToHead() {
	if len(this.head) == 0 {
		if len(this.tail) != 0 {
			for len(this.tail) > 0 {
				this.head = append(this.head, this.tail[len(this.tail)-1])
				this.tail = this.tail[:len(this.tail)-1]
			}
		}
	}
}

func (this *MyQueue) Peek() int {
	this.tailToHead()
	return this.head[len(this.head)-1]
}

func (this *MyQueue) Empty() bool {
	this.tailToHead()
	return len(this.head) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)
