package test

import (
	"github.com/6tail/tyme4go/tyme"
	"testing"
)

// TestNineStar0 九星测试
func TestNineStar0(t *testing.T) {
	m, _ := tyme.LunarYear{}.FromYear(1985)
	excepted := "六"
	got := m.GetNineStar().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "六白金"
	got = m.GetNineStar().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestNineStar1(t *testing.T) {
	m, _ := tyme.SolarDay{}.FromYmd(1985, 2, 19)
	excepted := "五黄土"
	got := m.GetLunarDay().GetNineStar().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "玉衡"
	got = m.GetLunarDay().GetNineStar().GetDipper().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
