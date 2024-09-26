package test

import (
	"github.com/6tail/tyme4go/tyme"
	"testing"
)

// TestSolarTerm0 公历半年测试
func TestSolarTerm0(t *testing.T) {
	d, _ := tyme.SolarTerm{}.FromName(2023, "冬至")
	excepted := "冬至"
	got := d.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "2022年12月22日"
	got = d.GetJulianDay().GetSolarDay().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
