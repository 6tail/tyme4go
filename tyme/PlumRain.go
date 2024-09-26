package tyme

// PlumRainNames 梅雨名称
var PlumRainNames = []string{"入梅", "出梅"}

// PlumRain 梅雨
type PlumRain struct {
	LoopTyme
}

func (PlumRain) FromIndex(index int) PlumRain {
	return PlumRain{LoopTyme{}.FromIndex(PlumRainNames, index)}
}

func (PlumRain) FromName(name string) (*PlumRain, error) {
	p, err := LoopTyme{}.FromName(PlumRainNames, name)
	if err != nil {
		return nil, err
	}
	return &PlumRain{*p}, nil
}

func (o PlumRain) Next(n int) PlumRain {
	return o.FromIndex(o.nextIndex(n))
}
