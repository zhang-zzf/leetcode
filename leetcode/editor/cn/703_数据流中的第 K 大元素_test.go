package leetcode

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when703_thenSuccess(t *testing.T) {
	kth := Constructor703(3, []int{4, 5, 8, 2})
	assert.Equal(t, 4, kth.Add(3))
	assert.Equal(t, 5, kth.Add(5))
	assert.Equal(t, 5, kth.Add(10))
	assert.Equal(t, 8, kth.Add(9))
	assert.Equal(t, 8, kth.Add(4))
}

//leetcode submit region begin(Prohibit modification and deletion)
type KthLargest struct {
	// TODO 继承 IntSlice
	sort.IntSlice
	k int
}

// TODO KthLargest 实现了 heap.Interface。
// TODO KthLargest 继承 sort.IntSlice, 也实现了 Sort.Interface
func (r *KthLargest) Push(x any) {
	r.IntSlice = append(r.IntSlice, x.(int))
}

func (r *KthLargest) Pop() any {
	n := r.Len()
	ans := r.IntSlice[n-1]
	r.IntSlice = r.IntSlice[:n-1]
	return ans
}

func Constructor703(k int, nums []int) KthLargest {
	kth := KthLargest{k: k}
	for _, n := range nums {
		kth.Add(n)
	}
	return kth
}

func (r *KthLargest) Add(val int) int {
	if r.Len() == r.k && val < r.IntSlice[0] {
		return r.IntSlice[0]
	}
	heap.Push(r, val)
	if r.Len() > r.k {
		heap.Pop(r)
	}
	return r.IntSlice[0]
}

func (r *KthLargest) Add1(val int) int {
	heap.Push(r, val)
	if r.Len() > r.k {
		heap.Pop(r)
	}
	return r.IntSlice[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
//leetcode submit region end(Prohibit modification and deletion)
