package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when278_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, 1, "shouldEqual")
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	left, right, ans := 1, n, 1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if isBadVersion(mid) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func isBadVersion(version int) bool {
	return true
}
