package tyme

// LunarSect2EightCharProvider Lunar流派2的八字计算（晚子时日柱算当天）
type LunarSect2EightCharProvider struct {
	IEightCharProvider
}

func (o LunarSect2EightCharProvider) GetEightChar(hour LunarHour) EightChar {
	h := hour.GetSixtyCycleHour()
	return EightChar{
		year:  h.GetYear(),
		month: h.GetMonth(),
		day:   hour.GetLunarDay().GetSixtyCycle(),
		hour:  h.GetSixtyCycle(),
	}
}
