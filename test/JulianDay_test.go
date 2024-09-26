package test

import (
	"github.com/6tail/tyme4go/tyme"
	"testing"
)

// TestJulianDay0 儒略日测试
func TestJulianDay0(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2023, 1, 1)
	excepted := "2023年1月1日"
	got := d.GetJulianDay().GetSolarDay().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
