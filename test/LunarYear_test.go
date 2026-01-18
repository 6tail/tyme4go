package test

import (
	"testing"

	"github.com/6tail/tyme4go/tyme"
)

// TestLunarYear0 农历年测试
func TestLunarYear0(t *testing.T) {
	m, _ := tyme.LunarYear{}.FromYear(2023)
	excepted := "农历癸卯年"
	got := m.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestLunarYear1(t *testing.T) {
	m, _ := tyme.LunarYear{}.FromYear(2020)
	excepted := "庚子"
	got := m.GetSixtyCycle().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestLunarYear2(t *testing.T) {
	m, _ := tyme.LunarYear{}.FromYear(1864)
	excepted := "上元"
	got := m.GetTwenty().GetSixty().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestLunarYear3(t *testing.T) {
	m, _ := tyme.LunarYear{}.FromYear(1884)
	excepted := "二运"
	got := m.GetTwenty().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
