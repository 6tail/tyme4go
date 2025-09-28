package tyme

// PhaseDay 月相第几天
type PhaseDay struct {
	AbstractCultureDay
}

func (PhaseDay) New(o Phase, index int) PhaseDay {
	return PhaseDay{AbstractCultureDay{}.New(o, index)}
}

// GetPhase 月相
func (o PhaseDay) GetPhase() Phase {
	return o.culture.(Phase)
}
