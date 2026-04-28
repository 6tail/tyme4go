package tyme

import (
	"fmt"
)

// AbstractFestival 节日抽象
type AbstractFestival struct {
	AbstractTyme
	// 类型
	festivalType FestivalType
	// 序号
	index int
	// 日
	day DayUnit
	// 事件
	event Event
}

func (AbstractFestival) New(festivalType FestivalType, index int, event Event, day DayUnit) AbstractFestival {
	return AbstractFestival{
		festivalType: festivalType,
		index:        index,
		event:        event,
		day:          day,
	}
}

// Deprecated:
func (o AbstractFestival) GetType() FestivalType {
	return o.festivalType
}

// GetDay 日
func (o AbstractFestival) GetDay() DayUnit {
	return o.day
}

// GetIndex 索引
func (o AbstractFestival) GetIndex() int {
	return o.index
}

func (o AbstractFestival) GetName() string {
	return o.event.name
}

func (o AbstractFestival) String() string {
	return fmt.Sprintf("%v %v", o.GetDay(), o.GetName())
}
