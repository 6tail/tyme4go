package test

import (
	"github.com/6tail/tyme4go/tyme"
	"testing"
)

// TestPhenology0 物候测试
func TestPhenology0(t *testing.T) {
	solarDay, _ := tyme.SolarDay{}.FromYmd(2020, 4, 23)
	// 七十二候
	phenology := solarDay.GetPhenologyDay()
	// 三候
	threePhenology := phenology.GetPhenology().GetThreePhenology()

	excepted := "谷雨"
	got := solarDay.GetTerm().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "初候"
	got = threePhenology.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "萍始生"
	got = phenology.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted2 := 4
	got2 := phenology.GetDayIndex()
	if excepted2 != got2 {
		t.Errorf("excepted: %v, got: %v", excepted2, got2)
	}
}
