package tyme

// ZodiacNames 生肖名称
var ZodiacNames = []string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}

// Zodiac 生肖
type Zodiac struct {
	LoopTyme
}

func (Zodiac) FromIndex(index int) Zodiac {
	return Zodiac{LoopTyme{}.FromIndex(ZodiacNames, index)}
}

func (Zodiac) FromName(name string) (*Zodiac, error) {
	p, err := LoopTyme{}.FromName(ZodiacNames, name)
	if err != nil {
		return nil, err
	}
	return &Zodiac{*p}, nil
}

func (o Zodiac) Next(n int) Zodiac {
	return o.FromIndex(o.nextIndex(n))
}
