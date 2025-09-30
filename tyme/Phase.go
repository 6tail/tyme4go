package tyme

import "math"

// PhaseNames 月相名称
var PhaseNames = []string{"新月", "蛾眉月", "上弦月", "盈凸月", "满月", "亏凸月", "下弦月", "残月"}

// Phase 月相
type Phase struct {
	LoopTyme
	lunarYear  int
	lunarMonth int
}

func (Phase) FromIndex(lunarYear int, lunarMonth int, index int) Phase {
	m, _ := LunarMonth{}.FromYm(lunarYear, lunarMonth)
	month := m.Next(index / len(PhaseNames))
	return Phase{
		LoopTyme{}.FromIndex(PhaseNames, index),
		month.GetYear(),
		month.GetMonthWithLeap(),
	}
}

func (Phase) FromName(lunarYear int, lunarMonth int, name string) (*Phase, error) {
	p, err := LoopTyme{}.FromName(PhaseNames, name)
	if err != nil {
		return nil, err
	}
	return &Phase{
		*p,
		lunarYear,
		lunarMonth,
	}, nil
}

func (o Phase) Next(n int) Phase {
	size := o.GetSize()
	i := o.GetIndex() + n
	if i < 0 {
		i -= size
	}
	i /= size
	m, _ := LunarMonth{}.FromYm(o.lunarYear, o.lunarMonth)
	month := *m
	if i != 0 {
		month = month.Next(i)
	}
	return o.FromIndex(month.GetYear(), month.GetMonthWithLeap(), o.nextIndex(n))
}

func (o Phase) getStartSolarTime() SolarTime {
	n := int(math.Floor(float64(o.lunarYear-2000) * 365.2422 / 29.53058886))
	i := 0
	jd := J2000 + OneThird
	lunarDay, _ := LunarDay{}.FromYmd(o.lunarYear, o.lunarMonth, 1)
	d := lunarDay.GetSolarDay()
	for {
		t := msaLonT(float64(n+i)*Pi2) * 36525
		if (!JulianDay{}.FromJulianDay(jd + t - DtT(t)).GetSolarDay().IsBefore(d)) {
			break
		}
		i++
	}
	t := msaLonT((float64(n+i)+float64([]int{0, 90, 180, 270}[o.GetIndex()/2])/360.0)*Pi2) * 36525
	return JulianDay{}.FromJulianDay(jd + t - DtT(t)).GetSolarTime()
}

func (o Phase) GetSolarTime() SolarTime {
	t := o.getStartSolarTime()
	if o.GetIndex()%2 == 1 {
		return t.Next(1)
	}
	return t
}

func (o Phase) GetSolarDay() SolarDay {
	d := o.getStartSolarTime().GetSolarDay()
	if o.GetIndex()%2 == 1 {
		return d.Next(1)
	}
	return d
}
