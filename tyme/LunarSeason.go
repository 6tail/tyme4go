package tyme

// LunarSeasonNames 农历季节名称
var LunarSeasonNames = []string{"孟春", "仲春", "季春", "孟夏", "仲夏", "季夏", "孟秋", "仲秋", "季秋", "孟冬", "仲冬", "季冬"}

// LunarSeason 农历季节
type LunarSeason struct {
	LoopTyme
}

func (LunarSeason) FromIndex(index int) LunarSeason {
	return LunarSeason{LoopTyme{}.FromIndex(LunarSeasonNames, index)}
}

func (LunarSeason) FromName(name string) (*LunarSeason, error) {
	p, err := LoopTyme{}.FromName(LunarSeasonNames, name)
	if err != nil {
		return nil, err
	}
	return &LunarSeason{*p}, nil
}

func (o LunarSeason) Next(n int) LunarSeason {
	return o.FromIndex(o.nextIndex(n))
}
