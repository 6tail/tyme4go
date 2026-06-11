package tyme

import (
	"fmt"
	"strconv"
)

// SecondUnit 秒
type SecondUnit struct {
	DayUnit

	// 时
	hour int
	// 分
	minute int
	// 秒
	second int
}

// GetHour 时
func (o SecondUnit) GetHour() int {
	return o.hour
}

// GetMinute 分
func (o SecondUnit) GetMinute() int {
	return o.minute
}

// GetSecond 秒
func (o SecondUnit) GetSecond() int {
	return o.second
}

func (SecondUnit) Validate(_ int, _ int, _ int, hour int, minute int, second int) error {
	if hour < 0 || hour > 23 {
		return fmt.Errorf("illegal hour: " + strconv.Itoa(hour))
	}
	if minute < 0 || minute > 59 {
		return fmt.Errorf("illegal minute: " + strconv.Itoa(minute))
	}
	if second < 0 || second > 59 {
		return fmt.Errorf("illegal second: " + strconv.Itoa(second))
	}
	return nil
}

// GetSecondsInDay 当天秒数
func (o SecondUnit) GetSecondsInDay() int {
	return o.hour*3600 + o.minute*60 + o.second
}

// GetCompareIndex 用于比较大小的索引
func (o SecondUnit) GetCompareIndex() int {
	return o.DayUnit.GetCompareIndex()*SecondPerDay + o.GetSecondsInDay()
}
