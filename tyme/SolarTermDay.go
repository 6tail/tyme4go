package tyme

// SolarTermDay 节气第几天
type SolarTermDay struct {
	AbstractCultureDay
}

func (SolarTermDay) New(o SolarTerm, index int) SolarTermDay {
	return SolarTermDay{AbstractCultureDay{}.New(o, index)}
}

// GetSolarTerm 节气
func (o SolarTermDay) GetSolarTerm() SolarTerm {
	return o.culture.(SolarTerm)
}
