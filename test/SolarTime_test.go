package test

import (
	"testing"

	"github.com/6tail/tyme4go/tyme"
)

// TestSolarTime0 公历时刻测试
func TestSolarTime0(t *testing.T) {
	time, _ := tyme.SolarTime{}.FromYmdHms(2023, 1, 1, 13, 5, 20)
	excepted := "13:05:20"
	got := time.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "13:04:59"
	got = time.Next(-21).GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
