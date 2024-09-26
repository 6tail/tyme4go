package test

import (
	"github.com/6tail/tyme4go/tyme"
	"testing"
)

// TestSolarDay0 公历日测试
func TestSolarDay0(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2023, 1, 1)
	excepted := "2023年1月1日"
	got := d.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "1日"
	got = d.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarDay1(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2020, 5, 24)
	excepted := "农历庚子年闰四月初二"
	got := d.GetLunarDay().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
