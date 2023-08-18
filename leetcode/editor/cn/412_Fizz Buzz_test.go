package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_givenNormal_when412_thenSuccess(t *testing.T) {
	ans := []string{"1", "2", "Fizz", "4", "Buzz",
		"Fizz", "7", "8", "Fizz", "Buzz",
		"11", "Fizz", "13", "14", "FizzBuzz"}
	assert.Equal(t, ans, fizzBuzz(15))
}

//leetcode submit region begin(Prohibit modification and deletion)
func fizzBuzz(n int) []string {
	ans := make([]string, n)
	for i := 0; i < n; i++ {
		if (i+1)%3 == 0 {
			ans[i] += "Fizz"
		}
		if (i+1)%5 == 0 {
			ans[i] += "Buzz"
		}
		if ans[i] == "" {
			ans[i] = strconv.Itoa(i + 1)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
