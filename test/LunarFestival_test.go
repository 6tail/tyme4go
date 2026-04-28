package test

import (
	"testing"

	"github.com/6tail/tyme4go/tyme"
)

// TestLunarFestival0 农历传统节日测试
func TestLunarFestival0(t *testing.T) {
	d := tyme.LunarFestival{}.FromIndex(2023, 0)
	if nil == d {
		t.Errorf("LunarFestival is nil")
		return
	}
	excepted := "农历甲辰年正月初一 春节"
	d1 := d.Next(13)
	if nil == d1 {
		t.Errorf("LunarFestival is nil")
		return
	}
	got := d1.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "农历壬寅年十一月廿九 冬至节"
	d2 := d.Next(-3)
	if nil == d2 {
		t.Errorf("LunarFestival is nil")
		return
	}
	got = d2.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestLunarFestival1(t *testing.T) {
	d, _ := tyme.LunarDay{}.FromYmd(2025, 5, 5)
	excepted := "农历乙巳年五月初五 端午节"
	f := d.GetFestival()
	if nil == f {
		t.Errorf("LunarFestival is nil")
		return
	}
	got := f.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
