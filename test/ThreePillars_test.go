package test

import (
	"github.com/6tail/tyme4go/tyme"
	"testing"
)

// TestThreePillars1 三柱测试
func TestThreePillars1(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(1034, 10, 2)
	excepted := "甲戌 甲戌 甲戌"
	got := d.GetSixtyCycleDay().GetThreePillars().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
