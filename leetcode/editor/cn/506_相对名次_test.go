package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"sort"
	"strconv"
	"testing"
)

func Test_givenNormal_when506_thenSuccess(t *testing.T) {
	params := []int{10, 3, 8, 9, 4}
	ans := findRelativeRanks(params)
	assert.Equal(t, []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
时间 O(n*log₂n)
空间 O(n)
Runtime:8 ms, faster than 98.63% of Go online submissions.
相比 findRelativeRanks1 Runtime:60 ms, faster than 5.48% of Go online submissions.
*/
func findRelativeRanks(score []int) []string {
	scoreToIdx := map[int]int{}
	for i, s := range score {
		scoreToIdx[s] = i
	}
	// sort
	sort.Slice(score, func(i, j int) bool {
		return score[i] > score[j]
	})
	ans := make([]string, len(score))
	for i, s := range score {
		ans[scoreToIdx[s]] = decideTitle(i + 1)
	}
	return ans
}

// O(n²)
func findRelativeRanks1(score []int) []string {
	ans := make([]string, len(score))
	max := math.MaxInt
	for i := 0; i < len(score); i++ {
		maxIdx, maxN := 0, math.MinInt
		for j := 0; j < len(score); j++ {
			if score[j] < max && score[j] > maxN {
				maxN = score[j]
				maxIdx = j
			}
		}
		max = maxN
		ans[maxIdx] = decideTitle(i + 1)
	}
	return ans
}

func decideTitle(nth int) string {
	if nth == 1 {
		return "Gold Medal"
	} else if nth == 2 {
		return "Silver Medal"
	} else if nth == 3 {
		return "Bronze Medal"
	} else {
		return strconv.Itoa(nth)
	}

}

//leetcode submit region end(Prohibit modification and deletion)
