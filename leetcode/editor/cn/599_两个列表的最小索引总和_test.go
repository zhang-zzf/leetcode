package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_givenNormal_when599_thenSuccess(t *testing.T) {
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 := []string{"KFC", "Shogun", "Burger King"}
	ans := findRestaurant(list1, list2)
	assert.Equal(t, []string{"Shogun"}, ans)
}

func Test_givenNormal1_when599_thenSuccess(t *testing.T) {
	list1 := []string{"Shogun", "Tapioca Express", "KFC", "Burger King"}
	list2 := []string{"KFC", "Shogun", "Burger King"}
	ans := findRestaurant(list1, list2)
	assert.Equal(t, []string{"Shogun"}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func findRestaurant(list1 []string, list2 []string) []string {
	if len(list1) > len(list2) {
		list1, list2 = list2, list1
	}
	idxMap := map[string]int{}
	for i, v := range list1 {
		idxMap[v] = i
	}
	var ans []string
	maxSum := math.MaxInt
	for i2, v := range list2 {
		if i1, ok := idxMap[v]; ok {
			sum := i2 + i1
			if sum == maxSum {
				ans = append(ans, v)
			} else if sum < maxSum {
				ans = []string{v}
				maxSum = sum
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
