package tyme

import "fmt"

// PlumRainDay 梅雨天
type PlumRainDay struct {
	AbstractCultureDay
}

func (PlumRainDay) New(o PlumRain, index int) PlumRainDay {
	return PlumRainDay{AbstractCultureDay{}.New(o, index)}
}

// GetPlumRain 梅雨
func (o PlumRainDay) GetPlumRain() PlumRain {
	return o.culture.(PlumRain)
}

func (o PlumRainDay) String() string {
	if o.GetPlumRain().GetIndex() == 0 {
		return fmt.Sprintf("%v第%d天", o.culture, o.dayIndex+1)
	}
	return o.culture.GetName()
}
