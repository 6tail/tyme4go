package tyme

import "fmt"

// HideHeavenStemDay 人元司令分野（地支藏干+天索引）
type HideHeavenStemDay struct {
	AbstractCultureDay
}

func (HideHeavenStemDay) New(o HideHeavenStem, index int) HideHeavenStemDay {
	return HideHeavenStemDay{AbstractCultureDay{}.New(o, index)}
}

// GetHideHeavenStem 藏干
func (o HideHeavenStemDay) GetHideHeavenStem() HideHeavenStem {
	return o.culture.(HideHeavenStem)
}

func (o HideHeavenStemDay) GetName() string {
	heavenStem := o.GetHideHeavenStem().GetHeavenStem()
	return fmt.Sprintf("%v%v", heavenStem.GetName(), heavenStem.GetElement().GetName())
}

func (o HideHeavenStemDay) String() string {
	return fmt.Sprintf("%v第%d天", o.GetName(), o.GetDayIndex()+1)
}
