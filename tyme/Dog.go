package tyme

// DogNames 三伏名称
var DogNames = []string{"初伏", "中伏", "末伏"}

// Dog 三伏
type Dog struct {
	LoopTyme
}

func (Dog) FromIndex(index int) Dog {
	return Dog{LoopTyme{}.FromIndex(DogNames, index)}
}

func (Dog) FromName(name string) (*Dog, error) {
	p, err := LoopTyme{}.FromName(DogNames, name)
	if err != nil {
		return nil, err
	}
	return &Dog{*p}, nil
}

func (o Dog) Next(n int) Dog {
	return o.FromIndex(o.nextIndex(n))
}
