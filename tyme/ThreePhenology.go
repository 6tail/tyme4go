package tyme

// ThreePhenologyNames 三候名称
var ThreePhenologyNames = []string{"初候", "二候", "三候"}

// ThreePhenology 三候
type ThreePhenology struct {
	LoopTyme
}

func (ThreePhenology) FromIndex(index int) ThreePhenology {
	return ThreePhenology{LoopTyme{}.FromIndex(ThreePhenologyNames, index)}
}

func (ThreePhenology) FromName(name string) (*ThreePhenology, error) {
	p, err := LoopTyme{}.FromName(ThreePhenologyNames, name)
	if err != nil {
		return nil, err
	}
	return &ThreePhenology{*p}, nil
}

func (o ThreePhenology) Next(n int) ThreePhenology {
	return o.FromIndex(o.nextIndex(n))
}
