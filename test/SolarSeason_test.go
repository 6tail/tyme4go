package test

import (
	"github.com/6tail/tyme4go/tyme"
	"testing"
)

// TestSolarSeason0 公历半年测试
func TestSolarSeason0(t *testing.T) {
	d, _ := tyme.SolarSeason{}.FromIndex(2023, 0)
	excepted := "2023年一季度"
	got := d.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "2021年四季度"
	got = d.Next(-5).String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
