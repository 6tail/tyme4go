package tyme

// AbstractChildLimitProvider 童限计算抽象
type AbstractChildLimitProvider struct {
	IChildLimitProvider
}

func (o AbstractChildLimitProvider) next(birthTime SolarTime, addYear int, addMonth int, addDay int, addHour int, addMinute int, addSecond int) ChildLimitInfo {
	d := birthTime.GetDay() + addDay
	h := birthTime.GetHour() + addHour
	mi := birthTime.GetMinute() + addMinute
	s := birthTime.GetSecond() + addSecond
	mi += s / 60
	s %= 60
	h += mi / 60
	mi %= 60
	d += h / 24
	h %= 24

	mt, _ := SolarMonth{}.FromYm(birthTime.GetYear()+addYear, birthTime.GetMonth())
	sm := mt.Next(addMonth)

	dc := sm.GetDayCount()
	for d > dc {
		d -= dc
		sm = sm.Next(1)
		dc = sm.GetDayCount()
	}

	t, _ := SolarTime{}.FromYmdHms(sm.GetYear(), sm.GetMonth(), d, h, mi, s)
	return ChildLimitInfo{}.New(birthTime, *t, addYear, addMonth, addDay, addHour, addMinute)
}
