package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when705_thenSuccess(t *testing.T) {
	ans := true
	assert.Equal(t, true, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
type MyHashSet struct {
	data map[int]struct{}
}

func Constructor705() MyHashSet {
	return MyHashSet{map[int]struct{}{}}
}

func (this *MyHashSet) Add(key int) {
	this.data[key] = struct{}{}
}

func (this *MyHashSet) Remove(key int) {
	delete(this.data, key)
}

func (this *MyHashSet) Contains(key int) bool {
	_, ok := this.data[key]
	return ok
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
//leetcode submit region end(Prohibit modification and deletion)
