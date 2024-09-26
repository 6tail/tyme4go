package test

import (
	"github.com/6tail/tyme4go/tyme"
	"testing"
)

// TestSolarFestival0 公历现代节日测试
func TestSolarFestival0(t *testing.T) {
	d, _ := tyme.SolarFestival{}.FromIndex(2023, 0)
	excepted := "2024年5月1日 五一劳动节"
	got := d.Next(13).String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "2022年8月1日 八一建军节"
	got = d.Next(-3).String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
