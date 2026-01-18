package test

import (
	"testing"

	"github.com/6tail/tyme4go/tyme"
)

// TestLunarMonth0 农历月测试
func TestLunarMonth0(t *testing.T) {
	m, _ := tyme.LunarMonth{}.FromYm(2023, 3)
	excepted := "农历癸卯年三月"
	got := m.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "农历癸卯年闰二月"
	got = m.Next(-1).String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
