package tyme

// IEightCharProvider 八字计算接口
type IEightCharProvider interface {

	// GetEightChar 根据农历时辰计算八字
	GetEightChar(hour LunarHour) EightChar
}
