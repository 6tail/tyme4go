package tyme

// SixtyCycleNames 六十甲子(六十干支周)名称
var SixtyCycleNames = []string{"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉", "甲戌", "乙亥", "丙子", "丁丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未", "甲申", "乙酉", "丙戌", "丁亥", "戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳", "甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥", "庚子", "辛丑", "壬寅", "癸卯", "甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥", "壬子", "癸丑", "甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥"}

// SixtyCycle 六十甲子(六十干支周)
type SixtyCycle struct {
	LoopTyme
}

func (SixtyCycle) FromIndex(index int) SixtyCycle {
	return SixtyCycle{LoopTyme{}.FromIndex(SixtyCycleNames, index)}
}

func (SixtyCycle) FromName(name string) (*SixtyCycle, error) {
	p, err := LoopTyme{}.FromName(SixtyCycleNames, name)
	if err != nil {
		return nil, err
	}
	return &SixtyCycle{*p}, nil
}

func (o SixtyCycle) Next(n int) SixtyCycle {
	return o.FromIndex(o.nextIndex(n))
}

// GetHeavenStem 天干
func (o SixtyCycle) GetHeavenStem() HeavenStem {
	return HeavenStem{}.FromIndex(o.index % len(HeavenStemNames))
}

// GetEarthBranch 地支
func (o SixtyCycle) GetEarthBranch() EarthBranch {
	return EarthBranch{}.FromIndex(o.index % len(EarthBranchNames))
}

// GetSound 纳音
func (o SixtyCycle) GetSound() Sound {
	return Sound{}.FromIndex(o.index / 2)
}

// GetPengZu 彭祖百忌
func (o SixtyCycle) GetPengZu() PengZu {
	return PengZu{}.FromSixtyCycle(o)
}

// GetTen 旬
func (o SixtyCycle) GetTen() Ten {
	return Ten{}.FromIndex((o.GetHeavenStem().GetIndex() - o.GetEarthBranch().GetIndex()) / 2)
}

// GetExtraEarthBranches 旬空(空亡)，因地支比天干多2个，旬空则为每一轮干支一一配对后多出来的2个地支
func (o SixtyCycle) GetExtraEarthBranches() []EarthBranch {
	var l []EarthBranch
	l[0] = EarthBranch{}.FromIndex(10 + o.GetEarthBranch().GetIndex() - o.GetHeavenStem().GetIndex())
	l[1] = l[0].Next(1)
	return l
}
