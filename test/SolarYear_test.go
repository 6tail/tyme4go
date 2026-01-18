package test

import (
	"testing"

	"github.com/6tail/tyme4go/tyme"
)

// TestSolarYear0 公历时刻测试
func TestSolarYear0(t *testing.T) {
	y, _ := tyme.SolarYear{}.FromYear(2023)
	excepted := "2023年"
	got := y.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarYear1(t *testing.T) {
	y, _ := tyme.SolarYear{}.FromYear(1500)
	got := y.IsLeap()
	if true != got {
		t.Errorf("excepted: %v, got: %v", true, got)
	}
}

func TestSolarYear2(t *testing.T) {
	y, _ := tyme.SolarYear{}.FromYear(1500)
	got := y.GetYear()
	if 1500 != got {
		t.Errorf("excepted: %v, got: %v", true, got)
	}
}
