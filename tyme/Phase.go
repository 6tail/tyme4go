package tyme

// PhaseNames 月相名称
var PhaseNames = []string{"朔月", "既朔月", "蛾眉新月", "蛾眉新月", "蛾眉月", "夕月", "上弦月", "上弦月", "九夜月", "宵月", "宵月", "宵月", "渐盈凸月", "小望月", "望月", "既望月", "立待月", "居待月", "寝待月", "更待月", "渐亏凸月", "下弦月", "下弦月", "有明月", "有明月", "蛾眉残月", "蛾眉残月", "残月", "晓月", "晦月"}

// Phase 月相
type Phase struct {
	LoopTyme
}

func (Phase) FromIndex(index int) Phase {
	return Phase{LoopTyme{}.FromIndex(PhaseNames, index)}
}

func (Phase) FromName(name string) (*Phase, error) {
	p, err := LoopTyme{}.FromName(PhaseNames, name)
	if err != nil {
		return nil, err
	}
	return &Phase{*p}, nil
}

func (o Phase) Next(n int) Phase {
	return o.FromIndex(o.nextIndex(n))
}
