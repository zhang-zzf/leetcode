package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when733_thenSuccess(t *testing.T) {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	ans := [][]int{
		{2, 2, 2},
		{2, 2, 0},
		{2, 0, 1},
	}
	filled := floodFill(image, 1, 1, 2)
	assert.Equal(t, ans, filled)
}

//leetcode submit region begin(Prohibit modification and deletion)
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	src := image[sr][sc]
	if src == color {
		return image
	}
	m, n := len(image), len(image[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}
		if image[i][j] != src {
			return
		}
		image[i][j] = color
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i+1, j)
		dfs(i, j-1)
	}
	dfs(sr, sc)
	return image
}

//leetcode submit region end(Prohibit modification and deletion)
