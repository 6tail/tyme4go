package test

import (
	"github.com/6tail/tyme4go/tyme"
	"testing"
)

func TestHideHeavenStemDay0(t *testing.T) {
	solar, _ := tyme.SolarDay{}.FromYmd(2024, 12, 4)
	d := solar.GetHideHeavenStemDay()

	excepted := "本气"
	got := d.GetHideHeavenStem().GetType().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "壬"
	got = d.GetHideHeavenStem().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "壬"
	got = d.GetHideHeavenStem().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "壬水"
	got = d.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "壬水第16天"
	got = d.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
