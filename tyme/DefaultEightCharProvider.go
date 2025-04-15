package tyme

// DefaultEightCharProvider 默认的八字计算（晚子时算第二天）
type DefaultEightCharProvider struct {
	IEightCharProvider
}

func (o DefaultEightCharProvider) GetEightChar(hour LunarHour) EightChar {
	return hour.GetSixtyCycleHour().GetEightChar()
}
