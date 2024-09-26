package tyme

// EarthBranchNames 地支名称
var EarthBranchNames = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

// EarthBranch 地支
type EarthBranch struct {
	LoopTyme
}

func (EarthBranch) FromIndex(index int) EarthBranch {
	return EarthBranch{LoopTyme{}.FromIndex(EarthBranchNames, index)}
}

func (EarthBranch) FromName(name string) (*EarthBranch, error) {
	p, err := LoopTyme{}.FromName(EarthBranchNames, name)
	if err != nil {
		return nil, err
	}
	return &EarthBranch{*p}, nil
}

func (o EarthBranch) Next(n int) EarthBranch {
	return o.FromIndex(o.nextIndex(n))
}

// GetElement 五行
func (o EarthBranch) GetElement() Element {
	return Element{}.FromIndex([]int{4, 2, 0, 0, 2, 1, 1, 2, 3, 3, 2, 4}[o.index])
}

// GetYinYang 阴阳
func (o EarthBranch) GetYinYang() YinYang {
	if o.index%2 == 0 {
		return YANG
	}
	return YIN
}

// GetHideHeavenStemMain 藏干之本气
func (o EarthBranch) GetHideHeavenStemMain() HeavenStem {
	return HeavenStem{}.FromIndex([]int{9, 5, 0, 1, 4, 2, 3, 5, 6, 7, 4, 8}[o.index])
}

// GetHideHeavenStemMiddle 藏干之中气，无中气返回nil
func (o EarthBranch) GetHideHeavenStemMiddle() *HeavenStem {
	n := []int{-1, 9, 2, -1, 1, 6, 5, 3, 8, -1, 7, 0}[o.index]
	if n == -1 {
		return nil
	}
	t := HeavenStem{}.FromIndex(n)
	return &t
}

// GetHideHeavenStemResidual 藏干之余气，无余气返回nil
func (o EarthBranch) GetHideHeavenStemResidual() *HeavenStem {
	n := []int{-1, 7, 4, -1, 9, 4, -1, 1, 4, -1, 3, -1}[o.index]
	if n == -1 {
		return nil
	}
	t := HeavenStem{}.FromIndex(n)
	return &t
}

// GetZodiac 生肖
func (o EarthBranch) GetZodiac() Zodiac {
	return Zodiac{}.FromIndex(o.index)
}

// GetDirection 方位
func (o EarthBranch) GetDirection() Direction {
	return Direction{}.FromIndex([]int{0, 4, 2, 2, 4, 8, 8, 4, 6, 6, 4, 0}[o.index])
}

// GetOminous 煞（逢巳日、酉日、丑日必煞东；亥日、卯日、未日必煞西；申日、子日、辰日必煞南；寅日、午日、戌日必煞北。）
func (o EarthBranch) GetOminous() Direction {
	return Direction{}.FromIndex([]int{8, 2, 0, 6}[o.index%4])
}

// GetPengZuEarthBranch 地支彭祖百忌
func (o EarthBranch) GetPengZuEarthBranch() PengZuEarthBranch {
	return PengZuEarthBranch{}.FromIndex(o.index)
}

// GetOpposite 六冲（子午冲，丑未冲，寅申冲，辰戌冲，卯酉冲，巳亥冲）
func (o EarthBranch) GetOpposite() EarthBranch {
	return o.Next(6)
}

// GetCombine 六合（子丑合，寅亥合，卯戌合，辰酉合，巳申合，午未合）
func (o EarthBranch) GetCombine() EarthBranch {
	return EarthBranch{}.FromIndex(1 - o.index)
}

// GetHarm 六害（子未害、丑午害、寅巳害、卯辰害、申亥害、酉戌害）
func (o EarthBranch) GetHarm() EarthBranch {
	return EarthBranch{}.FromIndex(19 - o.index)
}

// Combine 合化（子丑合化土，寅亥合化木，卯戌合化火，辰酉合化金，巳申合化水，午未合化土），如果无法合化，返回nil
func (o EarthBranch) Combine(target EarthBranch) *Element {
	if o.GetCombine().Equals(target) {
		t := Element{}.FromIndex([]int{2, 2, 0, 1, 3, 4, 2, 2, 4, 3, 1, 0}[o.index])
		return &t
	}
	return nil
}

func (o EarthBranch) Equals(target EarthBranch) bool {
	return o.String() == target.String()
}
