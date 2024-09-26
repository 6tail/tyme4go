package tyme

// AnimalNames 动物名称
var AnimalNames = []string{"蛟", "龙", "貉", "兔", "狐", "虎", "豹", "獬", "牛", "蝠", "鼠", "燕", "猪", "獝", "狼", "狗", "彘", "鸡", "乌", "猴", "猿", "犴", "羊", "獐", "马", "鹿", "蛇", "蚓"}

// Animal 动物
type Animal struct {
	LoopTyme
}

func (Animal) FromIndex(index int) Animal {
	return Animal{LoopTyme{}.FromIndex(AnimalNames, index)}
}

func (Animal) FromName(name string) (*Animal, error) {
	p, err := LoopTyme{}.FromName(AnimalNames, name)
	if err != nil {
		return nil, err
	}
	return &Animal{*p}, nil
}

func (o Animal) Next(n int) Animal {
	return o.FromIndex(o.nextIndex(n))
}
