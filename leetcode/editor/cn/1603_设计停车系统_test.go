package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_givenNormal_when1603_thenSuccess(t *testing.T) {
	assert.Equal(t, true, true)
}

//leetcode submit region begin(Prohibit modification and deletion)
type ParkingSystem struct {
	leftCnt [4]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{leftCnt: [4]int{0, big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	this.leftCnt[carType] -= 1
	return this.leftCnt[carType] >= 0
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
//leetcode submit region end(Prohibit modification and deletion)
