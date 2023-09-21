package leetcode

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_givenNormal_when1337_thenSuccess(t *testing.T) {
	param := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1}}
	ans := kWeakestRows(param, 5)
	assert.Equal(t, []int{2, 0, 3, 1, 4}, ans)
}

func Test_givenFailedCase1_when1337_thenSuccess(t *testing.T) {
	param := [][]int{
		{1, 1, 0},
		{1, 0, 0},
		{1, 0, 0},
		{1, 1, 1},
		{1, 1, 0},
		{0, 0, 0},
	}
	ans := kWeakestRows(param, 6)
	assert.Equal(t, []int{5, 1, 2, 0, 4, 3}, ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
// todo 参考答案
func kWeakestRows(mat [][]int, k int) (ans []int) {
	m, n := len(mat), len(mat[0])
	h := &hp1337{}
	for r := 0; r < m; r++ {
		// 2分查找
		cnt := sort.Search(n, func(c int) bool { return mat[r][c] == 0 })
		heap.Push(h, pair{cnt, r})
	}
	for i := 0; i < k; i++ {
		ans = append(ans, heap.Pop(h).(pair).idx)
	}
	return ans
}

type pair struct {
	cnt, idx int
}

// implement sort.Interface
type pairSlice []pair

func (x pairSlice) Len() int { return len(x) }
func (x pairSlice) Less(i, j int) bool {
	return x[i].cnt < x[j].cnt || (x[i].cnt == x[j].cnt && x[i].idx < x[j].idx)
}
func (x pairSlice) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

type hp1337 struct {
	pairSlice
}

func (h *hp1337) Push(a any) { h.pairSlice = append(h.pairSlice, a.(pair)) }
func (h *hp1337) Pop() (ans any) {
	n := len(h.pairSlice)
	ans = h.pairSlice[n-1]
	h.pairSlice = h.pairSlice[:n-1]
	return ans
}

func kWeakestRows2Heap(mat [][]int, k int) (ans []int) {
	m, n := len(mat), len(mat[0])
	// 小堆
	h := &MinHeap1337{}
	for i := 0; i < m; i++ {
		cnt := 0
		for j := 0; j < n && mat[i][j] == 1; j++ {
			cnt += 1
		}
		h.push([]int{i, cnt})
	}
	for i := 0; i < k; i++ {
		ans = append(ans, h.pop()[0])
	}
	return ans
}

type PairSlice [][]int

func (x PairSlice) Len() int { return len(x) }
func (x PairSlice) Less(i, j int) bool {
	return x[i][1] < x[j][1] || (x[i][1] == x[j][1] && x[i][0] < x[j][0])
}
func (x PairSlice) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

type MinHeap1337 struct {
	PairSlice
}

// add x as element Len()
func (h *MinHeap1337) Push(x any) { h.PairSlice = append(h.PairSlice, x.([]int)) }

// remove and return element Len() - 1.
func (h *MinHeap1337) Pop() (ans any) {
	n := len(h.PairSlice)
	ans = h.PairSlice[n-1]
	h.PairSlice = h.PairSlice[:n-1]
	return ans
}

func (h *MinHeap1337) push(x []int) { heap.Push(h, x) }
func (h *MinHeap1337) pop() []int   { return heap.Pop(h).([]int) }
func (h *MinHeap1337) peek() []int  { return h.PairSlice[0] }

func kWeakestRows1(mat [][]int, k int) []int {
	m, n := len(mat), len(mat[0])
	var ans []int
	ansMap := map[int]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[j][i] == 1 {
				continue
			}
			if _, ok := ansMap[j]; ok {
				continue
			}
			ans = append(ans, j)
			ansMap[j] += 1
			// 每列可能找到多个， failed case1
			// break
		}
	}
	if len(ans) < m {
		for i := 0; i < m; i++ {
			if _, ok := ansMap[i]; !ok {
				ans = append(ans, i)
			}
		}
	}
	return ans[:k]
}

//leetcode submit region end(Prohibit modification and deletion)
