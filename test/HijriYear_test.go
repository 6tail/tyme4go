package test

import (
	"testing"

	"github.com/6tail/tyme4go/tyme"
)

// TestHijriYear0 回历年测试
func TestHijriYear0(t *testing.T) {
	y, _ := tyme.HijriYear{}.FromYear(1)
	got := y.IsLeap()
	if got {
		t.Errorf("excepted: %v, got: %v", false, got)
	}
}

func TestHijriYear1(t *testing.T) {
	y, _ := tyme.HijriYear{}.FromYear(2)
	got := y.IsLeap()
	if !got {
		t.Errorf("excepted: %v, got: %v", true, got)
	}
}

func TestHijriYear2(t *testing.T) {
	y, _ := tyme.HijriYear{}.FromYear(0)
	got := y.IsLeap()
	if got {
		t.Errorf("excepted: %v, got: %v", false, got)
	}
}

func TestHijriYear3(t *testing.T) {
	y, _ := tyme.HijriYear{}.FromYear(-1)
	got := y.IsLeap()
	if !got {
		t.Errorf("excepted: %v, got: %v", true, got)
	}
}
