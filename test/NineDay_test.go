package test

import (
	"testing"

	"github.com/6tail/tyme4go/tyme"
)

// TestNineDay0 数九测试
func TestNineDay0(t *testing.T) {
	m, _ := tyme.SolarDay{}.FromYmd(2020, 12, 21)
	excepted := "一九"
	got := m.GetNineDay().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "一九"
	got = m.GetNineDay().GetNine().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "一九第1天"
	got = m.GetNineDay().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestNineDay1(t *testing.T) {
	m, _ := tyme.SolarDay{}.FromYmd(2020, 1, 7)
	excepted := "二九"
	got := m.GetNineDay().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "二九"
	got = m.GetNineDay().GetNine().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "二九第8天"
	got = m.GetNineDay().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
