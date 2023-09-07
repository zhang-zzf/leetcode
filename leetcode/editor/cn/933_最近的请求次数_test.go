package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when933_thenSuccess(t *testing.T) {
	rc := Constructor()
	assert.Equal(t, 1, rc.Ping(1))
	assert.Equal(t, 2, rc.Ping(100))
	assert.Equal(t, 3, rc.Ping(3001))
	assert.Equal(t, 3, rc.Ping(3002))
}

//leetcode submit region begin(Prohibit modification and deletion)
type RecentCounter struct {
	count []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

// TODO 参考答案
func (this *RecentCounter) Ping(t int) int {
	this.count = append(this.count, t)
	for this.count[0] < t-3000 {
		this.count = this.count[1:]
	}
	return len(this.count)
}

func (this *RecentCounter) Ping1(t int) int {
	this.count = append(this.count, t)
	// 2 分
	l, r, ans := 0, len(this.count)-1, 0
	for l <= r {
		mid := l + (r-l)>>1
		if this.count[mid] >= t-3000 {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	this.count = this.count[ans:]
	return len(this.count)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
//leetcode submit region end(Prohibit modification and deletion)
