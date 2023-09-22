package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_givenNormal_when1360_thenSuccess(t *testing.T) {
	assert.Equal(t, 1, daysBetweenDates("2019-06-29", "2019-06-30"))
	assert.Equal(t, 15, daysBetweenDates("2020-01-15", "2019-12-31"))
}

func Test_givenFailedCase1_when1360_thenSuccess(t *testing.T) {
	d1, _ := time.Parse("2006-01-02", "2080-08-08")
	d2, _ := time.Parse("2006-01-02", "2009-08-18")
	duration := d1.Sub(d2)
	days := duration.Hours() / 24
	assert.Equal(t, 25923, int(days))
	assert.Equal(t, 25923, daysBetweenDates("2080-08-08", "2009-08-18"))
}

func Test_givenFailedCase11_when1360_thenSuccess(t *testing.T) {
	days1 := daysSince1971("2009-08-18")
	days2 := daysSince1971("2080-08-08")
	days11 := daysSince197111("2009-08-18")
	days21 := daysSince197111("2080-08-08")
	assert.Equal(t, days1, days11)
	assert.Equal(t, days2, days21)
}

func daysSince197111(date string) int {
	d, _ := time.Parse("2006-01-02", date)
	_1971, _ := time.Parse("2006-01-02", "1971-01-01")
	duration := d.Sub(_1971)
	return int(duration.Hours() / 24)
}

//leetcode submit region begin(Prohibit modification and deletion)
var daysBefore = [...]int32{
	0,
	31,
	31 + 28,
	31 + 28 + 31,
	31 + 28 + 31 + 30,
	31 + 28 + 31 + 30 + 31,
	31 + 28 + 31 + 30 + 31 + 30,
	31 + 28 + 31 + 30 + 31 + 30 + 31,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30 + 31,
}

func leapYear(y int) bool {
	return y%400 == 0 || (y%4 == 0 && y%100 != 0)
}

func daysSince1971(date string) int {
	ans := 0
	d, _ := time.Parse("2006-01-02", date)
	for y := 1971; y < d.Year(); y++ {
		ans += 365
		if leapYear(y) {
			ans += 1
		}
	}
	ans += int(daysBefore[d.Month()-1])
	// TODO
	//if leapYear(d.Year()) {
	//	ans += 1
	//}
	// TODO
	//if d.Month() > time.January && leapYear(d.Year()) {
	//	ans += 1
	//}
	if d.Month() > time.February && leapYear(d.Year()) {
		ans += 1
	}
	// TODO
	// ans += d.Day()
	ans += d.Day() - 1
	return ans
}

func daysBetweenDates(date1 string, date2 string) int {
	days := daysSince1971(date1) - daysSince1971(date2)
	if days < 0 {
		days = -days
	}
	return days
}

//leetcode submit region end(Prohibit modification and deletion)
