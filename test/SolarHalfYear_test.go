package test

import (
	"testing"

	"github.com/6tail/tyme4go/tyme"
)

// TestSolarHalfYear0 公历半年测试
func TestSolarHalfYear0(t *testing.T) {
	d, _ := tyme.SolarHalfYear{}.FromIndex(2023, 0)
	excepted := "上半年"
	got := d.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "2023年上半年"
	got = d.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
