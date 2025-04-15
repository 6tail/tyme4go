package tyme

// ChildLimitProvider 童限计算接口
var ChildLimitProvider IChildLimitProvider = DefaultChildLimitProvider{}

// ChildLimit 童限
type ChildLimit struct {
	// 八字
	eightChar EightChar
	// 性别
	gender Gender
	// 顺逆
	forward bool
	// 童限信息
	info ChildLimitInfo
}

func (ChildLimit) FromSolarTime(birthTime SolarTime, gender Gender) ChildLimit {
	eightChar := birthTime.GetLunarHour().GetEightChar()
	// 阳男阴女顺推，阴男阳女逆推
	yang := YANG == eightChar.GetYear().GetHeavenStem().GetYinYang()
	man := MAN == gender
	forward := (yang && man) || (!yang && !man)
	term := birthTime.GetTerm()
	if !term.IsJie() {
		term = term.Next(-1)
	}
	if forward {
		term = term.Next(2)
	}
	info := ChildLimitProvider.GetInfo(birthTime, term)

	return ChildLimit{
		eightChar: eightChar,
		gender:    gender,
		forward:   forward,
		info:      info,
	}
}

// GetEightChar 八字
func (o ChildLimit) GetEightChar() EightChar {
	return o.eightChar
}

// GetGender 性别
func (o ChildLimit) GetGender() Gender {
	return o.gender
}

// IsForward 是否顺推
func (o ChildLimit) IsForward() bool {
	return o.forward
}

// GetYearCount 年数
func (o ChildLimit) GetYearCount() int {
	return o.info.GetYearCount()
}

// GetMonthCount 月数
func (o ChildLimit) GetMonthCount() int {
	return o.info.GetMonthCount()
}

// GetDayCount 日数
func (o ChildLimit) GetDayCount() int {
	return o.info.GetDayCount()
}

// GetHourCount 小时数
func (o ChildLimit) GetHourCount() int {
	return o.info.GetHourCount()
}

// GetMinuteCount 分钟数
func (o ChildLimit) GetMinuteCount() int {
	return o.info.GetMinuteCount()
}

// GetStartTime 开始(即出生)的公历时刻
func (o ChildLimit) GetStartTime() SolarTime {
	return o.info.GetStartTime()
}

// GetEndTime 结束(即开始起运)的公历时刻
func (o ChildLimit) GetEndTime() SolarTime {
	return o.info.GetEndTime()
}

// GetStartDecadeFortune 起运大运
func (o ChildLimit) GetStartDecadeFortune() DecadeFortune {
	return DecadeFortune{}.FromChildLimit(o, 0)
}

// GetDecadeFortune 所属大运
func (o ChildLimit) GetDecadeFortune() DecadeFortune {
	return DecadeFortune{}.FromChildLimit(o, -1)
}

// GetStartFortune 小运
func (o ChildLimit) GetStartFortune() Fortune {
	return Fortune{}.FromChildLimit(o, 0)
}

// Deprecated: Use GetEndSixtyCycleYear instead.
func (o ChildLimit) GetEndLunarYear() LunarYear {
	y, _ := LunarYear{}.FromYear(o.GetStartTime().GetLunarHour().GetYear() + o.GetEndTime().GetYear() - o.GetStartTime().GetYear())
	return *y
}

// GetStartSixtyCycleYear 开始(即出生)干支年
func (o ChildLimit) GetStartSixtyCycleYear() SixtyCycleYear {
	y, _ := SixtyCycleYear{}.FromYear(o.GetStartTime().GetYear())
	return *y
}

// GetEndSixtyCycleYear 结束(即起运)干支年
func (o ChildLimit) GetEndSixtyCycleYear() SixtyCycleYear {
	y, _ := SixtyCycleYear{}.FromYear(o.GetEndTime().GetYear())
	return *y
}

// GetStartAge 开始年龄
func (o ChildLimit) GetStartAge() int {
	return 1
}

// GetEndAge 结束年龄
func (o ChildLimit) GetEndAge() int {
	n := o.GetEndSixtyCycleYear().GetYear() - o.GetStartSixtyCycleYear().GetYear()
	if n > 1 {
		return n
	}
	return 1
}
