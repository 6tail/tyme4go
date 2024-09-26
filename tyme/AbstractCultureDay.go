package tyme

import "fmt"

// AbstractCultureDay 带天索引的传统文化抽象
type AbstractCultureDay struct {
	AbstractCulture
	culture  Culture
	dayIndex int
}

func (AbstractCultureDay) New(culture Culture, dayIndex int) AbstractCultureDay {
	return AbstractCultureDay{culture: culture, dayIndex: dayIndex}
}

// GetDayIndex 天索引
func (o AbstractCultureDay) GetDayIndex() int {
	return o.dayIndex
}

func (o AbstractCultureDay) GetCulture() Culture {
	return o.culture
}

func (o AbstractCultureDay) GetName() string {
	return o.culture.GetName()
}

func (o AbstractCultureDay) String() string {
	return fmt.Sprintf("%v第%d天", o.culture, o.dayIndex+1)
}
