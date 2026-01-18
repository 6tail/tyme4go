package test

import (
	"testing"

	"github.com/6tail/tyme4go/tyme"
)

// TestPlumRainDay0 梅雨天测试
func TestPlumRainDay0(t *testing.T) {
	solarDay, _ := tyme.SolarDay{}.FromYmd(2024, 6, 11)
	d := solarDay.GetPlumRainDay()

	excepted := "入梅"
	got := d.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "入梅"
	got = d.GetPlumRain().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "入梅第1天"
	got = d.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
