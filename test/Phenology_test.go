package test

import (
	"github.com/6tail/tyme4go/tyme"
	"testing"
)

// TestPhenology0 物候测试
func TestPhenology0(t *testing.T) {
	solarDay, _ := tyme.SolarDay{}.FromYmd(2020, 4, 23)
	// 七十二候
	phenology := solarDay.GetPhenologyDay()
	// 三候
	threePhenology := phenology.GetPhenology().GetThreePhenology()

	excepted := "谷雨"
	got := solarDay.GetTerm().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "初候"
	got = threePhenology.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "萍始生"
	got = phenology.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted2 := 4
	got2 := phenology.GetDayIndex()
	if excepted2 != got2 {
		t.Errorf("excepted: %v, got: %v", excepted2, got2)
	}
}

func TestPhenology1(t *testing.T) {
	p := tyme.Phenology{}.FromIndex(2026, 1)
	jd := p.GetJulianDay()

	excepted := "麋角解"
	got := p.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "2025年12月26日"
	got = jd.GetSolarDay().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "2025年12月26日 20:49:56"
	got = jd.GetSolarTime().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestPhenology2(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2025, 12, 26)
	p := d.GetPhenology()
	jd := p.GetJulianDay()

	excepted := "麋角解"
	got := p.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "2025年12月26日"
	got = jd.GetSolarDay().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "2025年12月26日 20:49:56"
	got = jd.GetSolarTime().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestPhenology3(t *testing.T) {
	d, _ := tyme.SolarTime{}.FromYmdHms(2025, 12, 26, 20, 49, 55)
	p := d.GetPhenology()

	excepted := "蚯蚓结"
	got := p.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	d, _ = tyme.SolarTime{}.FromYmdHms(2025, 12, 26, 20, 49, 56)
	p = d.GetPhenology()

	excepted = "麋角解"
	got = p.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
