package leetcode

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when1046_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
}

//leetcode submit region begin(Prohibit modification and deletion)
type MaxHeap struct {
	sort.IntSlice
}

func (h *MaxHeap) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }

//add x as element Len()
func (h *MaxHeap) Push(x interface{}) { h.IntSlice = append(h.IntSlice, x.(int)) }

//remove and return element Len() - 1.
func (h *MaxHeap) Pop() any {
	n := h.Len()
	val := h.IntSlice[n-1]
	h.IntSlice = h.IntSlice[:n-1]
	return val
}

func (h *MaxHeap) push(x int) { heap.Push(h, x) }
func (h *MaxHeap) pop() int   { return heap.Pop(h).(int) }

func lastStoneWeight(stones []int) int {
	h := &MaxHeap{}
	for _, stone := range stones {
		heap.Push(h, stone)
	}
	for h.Len() > 1 {
		h.push(h.pop() - h.pop())
	}
	return h.pop()
}

//leetcode submit region end(Prohibit modification and deletion)
