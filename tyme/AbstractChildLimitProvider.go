package tyme

// AbstractChildLimitProvider 童限计算抽象
type AbstractChildLimitProvider struct {
	IChildLimitProvider
}

func (o AbstractChildLimitProvider) next(birthTime SolarTime, addYear int, addMonth int, addDay int, addHour int, addMinute int, addSecond int) ChildLimitInfo {
	mt, err := SolarMonth{}.FromYm(birthTime.GetYear()+addYear, birthTime.GetMonth())
	if err != nil {
		panic(err)
	}
	sm := mt.Next(addMonth)

	// 年月平移后保留合法日号，避免把1582年10月15日当作当月第15天。
	t, err := SolarTime{}.FromYmdHms(sm.GetYear(), sm.GetMonth(), birthTime.GetDay(), birthTime.GetHour(), birthTime.GetMinute(), birthTime.GetSecond())
	dayOffset := 0
	if err != nil {
		// 日号不存在时从月初顺延，保留月末溢出，并跳过公历缺日。
		t, err = SolarTime{}.FromYmdHms(sm.GetYear(), sm.GetMonth(), 1, birthTime.GetHour(), birthTime.GetMinute(), birthTime.GetSecond())
		if err != nil {
			panic(err)
		}
		dayOffset = birthTime.GetDay() - 1
	}

	// 日时分秒按实际经过时间推进，由历法层处理1582年的公历跳日。
	endTime := t.Next((dayOffset+addDay)*SecondPerDay + addHour*3600 + addMinute*60 + addSecond)
	return ChildLimitInfo{}.New(birthTime, endTime, addYear, addMonth, addDay, addHour, addMinute)
}
