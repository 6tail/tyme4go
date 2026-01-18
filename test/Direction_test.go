package test

import (
	"testing"

	"github.com/6tail/tyme4go/tyme"
)

func TestDirection0(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2021, 11, 13)
	excepted := "东南"
	got := d.GetLunarDay().GetSixtyCycle().GetHeavenStem().GetMascotDirection().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestDirection1(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2024, 1, 1)
	excepted := "东南"
	got := d.GetLunarDay().GetSixtyCycle().GetHeavenStem().GetMascotDirection().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestDirection2(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2023, 11, 6)
	excepted := "东"
	got := d.GetLunarDay().GetJupiterDirection().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
