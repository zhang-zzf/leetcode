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
	stack.PushBack(nil)
	stack.PushBack(A)
	// Peek
	top := stack.Back()
	assert.Equal(t, A, top.Value)
	// Pop
	for stack.Len() > 0 {
		topVal := stack.Remove(stack.Back())
		assert.Contains(t, []any{nil, A}, topVal)
	}
}
