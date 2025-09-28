package test

import (
	"github.com/6tail/tyme4go/tyme"
	"testing"
)

// TestSolarTerm0 节气测试
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

func TestSolarTerm1(t *testing.T) {
	d, _ := tyme.SolarTerm{}.FromName(2023, "冬至")
	excepted := "2022年12月22日 05:48:12"
	got := d.GetJulianDay().GetSolarTime().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarTerm2(t *testing.T) {
	d, _ := tyme.SolarTerm{}.FromName(2025, "惊蛰")
	excepted := "2025年3月5日 16:07:18"
	got := d.GetJulianDay().GetSolarTime().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
