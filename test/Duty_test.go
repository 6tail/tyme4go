package test

import (
	"github.com/6tail/tyme4go/tyme"
	"testing"
)

func TestDuty0(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2023, 10, 30)
	excepted := "闭"
	got := d.GetLunarDay().GetDuty().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestDuty1(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2023, 10, 19)
	excepted := "建"
	got := d.GetLunarDay().GetDuty().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestDuty2(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2023, 10, 7)
	excepted := "除"
	got := d.GetLunarDay().GetDuty().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestDuty3(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2023, 10, 8)
	excepted := "除"
	got := d.GetLunarDay().GetDuty().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
