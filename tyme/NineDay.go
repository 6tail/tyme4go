package tyme

// NineDay 数九天
type NineDay struct {
	AbstractCultureDay
}

func (NineDay) New(o Nine, index int) NineDay {
	return NineDay{AbstractCultureDay{}.New(o, index)}
}

// GetNine 数九
func (o NineDay) GetNine() Nine {
	return o.culture.(Nine)
}
