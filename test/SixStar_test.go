package test

import (
	"github.com/6tail/tyme4go/tyme"
	"testing"
)

// TestSixStar0 六曜测试
func TestSixStar0(t *testing.T) {
	solarDay, _ := tyme.SolarDay{}.FromYmd(2020, 4, 23)
	d := solarDay.GetLunarDay().GetSixStar()

	excepted := "佛灭"
	got := d.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
