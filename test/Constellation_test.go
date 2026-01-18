package test

import (
	"testing"

	"github.com/6tail/tyme4go/tyme"
)

func TestConstellation0(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2020, 3, 21)
	excepted := "白羊"
	got := d.GetConstellation().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	d, _ = tyme.SolarDay{}.FromYmd(2020, 4, 19)
	got = d.GetConstellation().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestConstellation1(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2020, 4, 20)
	excepted := "金牛"
	got := d.GetConstellation().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	d, _ = tyme.SolarDay{}.FromYmd(2020, 5, 20)
	got = d.GetConstellation().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestConstellation2(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2020, 5, 21)
	excepted := "双子"
	got := d.GetConstellation().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	d, _ = tyme.SolarDay{}.FromYmd(2020, 6, 21)
	got = d.GetConstellation().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestConstellation3(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2020, 6, 22)
	excepted := "巨蟹"
	got := d.GetConstellation().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	d, _ = tyme.SolarDay{}.FromYmd(2020, 7, 22)
	got = d.GetConstellation().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
