package test

import (
	"github.com/6tail/tyme4go/tyme"
	"testing"
)

// TestLunarHour0 农历时辰测试
func TestLunarHour0(t *testing.T) {
	h, _ := tyme.LunarHour{}.FromYmdHms(2020, -4, 5, 23, 0, 0)
	excepted := "子时"
	got := h.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "农历庚子年闰四月初五戊子时"
	got = h.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
