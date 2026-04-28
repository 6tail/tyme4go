package test

import (
	"testing"

	"github.com/6tail/tyme4go/tyme"
)

// TestSolarFestival0 公历现代节日测试
func TestSolarFestival0(t *testing.T) {
	d := tyme.SolarFestival{}.FromIndex(2023, 0)
	if d == nil {
		t.Errorf("SolarFestival is nil")
		return
	}
	excepted := "2024年5月1日 劳动节"
	d1 := d.Next(13)
	if d1 == nil {
		t.Errorf("SolarFestival is nil")
		return
	}
	got := d1.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "2022年8月1日 建军节"
	d2 := d.Next(-3)
	if d2 == nil {
		t.Errorf("SolarFestival is nil")
		return
	}
	got = d2.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
