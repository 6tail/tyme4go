package tyme

// DogDay 三伏天
type DogDay struct {
	AbstractCultureDay
}

func (DogDay) New(o Dog, index int) DogDay {
	return DogDay{AbstractCultureDay{}.New(o, index)}
}

// GetDog 三伏
func (o DogDay) GetDog() Dog {
	return o.culture.(Dog)
}
