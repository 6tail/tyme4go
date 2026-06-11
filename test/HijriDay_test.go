package test

import (
	"testing"

	"github.com/6tail/tyme4go/tyme"
)

// TestHijriDay0 公历日测试
func TestHijriDay0(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(622, 7, 16)
	h := d.GetHijriDay()
	excepted := "1年穆哈兰姆月1日"
	got := h.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestHijriDay1(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2026, 5, 13)
	h := d.GetHijriDay()
	excepted := "1447年都尔喀尔德月26日"
	got := h.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	h1, _ := tyme.HijriDay{}.FromYmd(1447, 11, 26)
	d1 := h1.GetSolarDay()
	excepted = "2026年5月13日"
	got = d1.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestHijriDay2(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(100, 7, 8)
	h := d.GetHijriDay()
	excepted := "-538年都尔黑哲月12日"
	got := h.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	h1, _ := tyme.HijriDay{}.FromYmd(-538, 12, 12)
	d1 := h1.GetSolarDay()
	excepted = "100年7月8日"
	got = d1.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestHijriDay3(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(622, 7, 15)
	h := d.GetHijriDay()
	excepted := "0年都尔黑哲月29日"
	got := h.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	h1, _ := tyme.HijriDay{}.FromYmd(0, 12, 29)
	d1 := h1.GetSolarDay()
	excepted = "622年7月15日"
	got = d1.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestHijriDay4(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(1, 1, 1)
	h := d.GetHijriDay()
	excepted := "-640年主马达·敖外鲁月16日"
	got := h.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	h1, _ := tyme.HijriDay{}.FromYmd(-640, 5, 16)
	d1 := h1.GetSolarDay()
	excepted = "1年1月1日"
	got = d1.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestHijriDay5(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(9999, 12, 31)
	h := d.GetHijriDay()
	excepted := "9666年赖比尔·阿色尼月2日"
	got := h.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	h1, _ := tyme.HijriDay{}.FromYmd(9666, 4, 2)
	d1 := h1.GetSolarDay()
	excepted = "9999年12月31日"
	got = d1.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
