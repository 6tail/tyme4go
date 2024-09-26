package tyme

// PengZuEarthBranchNames 地支彭祖百忌名称
var PengZuEarthBranchNames = []string{"子不问卜自惹祸殃", "丑不冠带主不还乡", "寅不祭祀神鬼不尝", "卯不穿井水泉不香", "辰不哭泣必主重丧", "巳不远行财物伏藏", "午不苫盖屋主更张", "未不服药毒气入肠", "申不安床鬼祟入房", "酉不会客醉坐颠狂", "戌不吃犬作怪上床", "亥不嫁娶不利新郎"}

// PengZuEarthBranch 地支彭祖百忌
type PengZuEarthBranch struct {
	LoopTyme
}

func (PengZuEarthBranch) FromIndex(index int) PengZuEarthBranch {
	return PengZuEarthBranch{LoopTyme{}.FromIndex(PengZuEarthBranchNames, index)}
}

func (PengZuEarthBranch) FromName(name string) (*PengZuEarthBranch, error) {
	p, err := LoopTyme{}.FromName(PengZuEarthBranchNames, name)
	if err != nil {
		return nil, err
	}
	return &PengZuEarthBranch{*p}, nil
}

func (o PengZuEarthBranch) Next(n int) PengZuEarthBranch {
	return o.FromIndex(o.nextIndex(n))
}
