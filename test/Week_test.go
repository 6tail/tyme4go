package test

import (
	"testing"

	"github.com/6tail/tyme4go/tyme"
)

// TestWeek0 星期测试
func TestWeek0(t *testing.T) {
	w, _ := tyme.SolarDay{}.FromYmd(1582, 10, 1)
	excepted := "一"
	got := w.GetWeek().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestWeek1(t *testing.T) {
	w, _ := tyme.SolarWeek{}.FromYm(2023, 10, 0, 0)
	excepted := "第一周"
	got := w.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "2023年10月第一周"
	got = w.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestWeek2(t *testing.T) {
	solarDay, _ := tyme.SolarDay{}.FromYmd(2025, 3, 6)
	solarWeek := solarDay.GetSolarWeek(0)

	excepted := "2025年3月第二周"
	got := solarWeek.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "2025年3月2日"
	got = solarWeek.GetFirstDay().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
