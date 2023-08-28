package strtucture

import (
	"container/heap"
	"container/list"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

const A = "A"

/**
Stack
using Slice
*/
func Test_givenStack_whenUsingSlice_thenSuccess(t *testing.T) {
	var stack []any
	// Push
	stack = append(stack, 1, 2, 3)
	// Pop
	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	assert.Equal(t, 3, top)
	// Pop again
	top = stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	assert.Equal(t, 2, top)
	// pop all
	for len(stack) > 0 {
		top = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
}

/**
Queue
*/
func Test_givenQueue_when_thenSuccess(t *testing.T) {
	var queue []any
	// add to tail
	queue = append(queue, 0, 1, 2)
	// peek the header
	assert.Equal(t, 0, queue[0])
	// remove from the header
	_ = queue[0]
	queue = queue[1:]
	assert.Equal(t, 1, queue[0])
}

/**
Stack
*/
func Test_givenStack_when_thenSuccess(t *testing.T) {
	stack := list.New()
	// push
	stack.PushBack(A)
	stack.PushBack(nil)
	// Peek
	top := stack.Back()
	assert.Equal(t, nil, top.Value)
	// Pop
	for stack.Len() > 0 {
		topVal := stack.Remove(stack.Back())
		assert.Contains(t, []any{nil, A}, topVal)
	}
}

type IntMinPQ struct {
	sort.IntSlice
}

//add x as element Len()
func (r *IntMinPQ) Push(x any) { r.IntSlice = append(r.IntSlice, x.(int)) }

//remove and return element Len() - 1.
func (r *IntMinPQ) Pop() any {
	n := r.Len()
	val := r.IntSlice[n-1]
	r.IntSlice = r.IntSlice[:n-1]
	return val
}

/**
heap
min-heap
*/
func Test_givenMinHeap_when_thenSuccess(t *testing.T) {
	aSlice := []int{5, 4, 3, 2, 1}
	pq := &IntMinPQ{}
	for _, n := range aSlice {
		heap.Push(pq, n)
		assert.Equal(t, n, pq.IntSlice[0])
	}
	for i := len(aSlice) - 1; i >= 0; i-- {
		assert.Equal(t, aSlice[i], heap.Pop(pq))
	}
}

type IntMaxPQ struct {
	IntMinPQ
}

func (x *IntMaxPQ) Less(i, j int) bool { return x.IntSlice[i] > x.IntSlice[j] }

/**
heap
max-heap
*/
func Test_givenMaxHeap_when_thenSuccess(t *testing.T) {
	aSlice := []int{1, 2, 3, 4, 5}
	pq := &IntMaxPQ{}
	for _, n := range aSlice {
		heap.Push(pq, n)
		assert.Equal(t, n, pq.IntSlice[0])
	}
	for i := len(aSlice) - 1; i >= 0; i-- {
		assert.Equal(t, aSlice[i], heap.Pop(pq))
	}
}
