package test

import (
	"github.com/6tail/tyme4go/tyme"
	"testing"
)

// TestSolarMonth0 公历半年测试
func TestSolarMonth0(t *testing.T) {
	d, _ := tyme.SolarMonth{}.FromYm(2019, 5)
	excepted := "5月"
	got := d.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "2019年5月"
	got = d.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
