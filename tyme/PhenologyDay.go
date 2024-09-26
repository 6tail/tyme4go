package tyme

// PhenologyDay 七十二候
type PhenologyDay struct {
	AbstractCultureDay
}

func (PhenologyDay) New(o Phenology, index int) PhenologyDay {
	return PhenologyDay{AbstractCultureDay{}.New(o, index)}
}

// GetPhenology 候
func (o PhenologyDay) GetPhenology() Phenology {
	return o.culture.(Phenology)
}
