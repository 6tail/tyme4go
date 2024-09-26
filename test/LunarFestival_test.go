package test

import (
	"github.com/6tail/tyme4go/tyme"
	"testing"
)

// TestLunarFestival0 农历传统节日测试
func TestLunarFestival0(t *testing.T) {
	d, _ := tyme.LunarFestival{}.FromIndex(2023, 0)
	excepted := "农历甲辰年正月初一 春节"
	got := d.Next(13).String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "农历壬寅年十一月廿九 冬至节"
	got = d.Next(-3).String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
