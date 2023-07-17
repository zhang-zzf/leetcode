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
