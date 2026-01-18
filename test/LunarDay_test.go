package test

import (
	"testing"

	"github.com/6tail/tyme4go/tyme"
)

// TestLunarDay0 农历日测试
func TestLunarDay0(t *testing.T) {
	d, _ := tyme.LunarDay{}.FromYmd(0, 11, 18)
	excepted := "1年1月1日"
	got := d.GetSolarDay().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestLunarDay1(t *testing.T) {
	d, _ := tyme.LunarDay{}.FromYmd(9999, 12, 2)
	excepted := "9999年12月31日"
	got := d.GetSolarDay().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestLunarDay2(t *testing.T) {
	d, _ := tyme.LunarDay{}.FromYmd(2023, 8, 24)
	excepted := "己亥"
	got := d.GetSixtyCycle().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestLunarDay3(t *testing.T) {
	d, _ := tyme.LunarDay{}.FromYmd(2012, 3, 1)
	excepted := "农历壬辰年闰四月初一"
	got := d.Next(60).String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

// TestLunarDay4 二十八宿测试
func TestLunarDay4(t *testing.T) {
	d, _ := tyme.LunarDay{}.FromYmd(2020, 4, 13)
	star := d.GetTwentyEightStar()
	excepted := "南"
	got := star.GetZone().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "朱雀"
	got = star.GetZone().GetBeast().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "翼"
	got = star.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "火"
	got = star.GetSevenStar().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "蛇"
	got = star.GetAnimal().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "凶"
	got = star.GetLuck().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "阳天"
	got = star.GetLand().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "东南"
	got = star.GetLand().GetDirection().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestLunarDay5(t *testing.T) {
	d, _ := tyme.LunarDay{}.FromYmd(2024, 1, 1)
	excepted := "农历甲辰年二月初三"
	got := d.Next(31).String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
