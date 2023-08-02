package strtucture

import (
	"container/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

const A = "A"

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
