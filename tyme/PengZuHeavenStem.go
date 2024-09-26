package tyme

// PengZuHeavenStemNames 天干彭祖百忌名称
var PengZuHeavenStemNames = []string{"甲不开仓财物耗散", "乙不栽植千株不长", "丙不修灶必见灾殃", "丁不剃头头必生疮", "戊不受田田主不祥", "己不破券二比并亡", "庚不经络织机虚张", "辛不合酱主人不尝", "壬不泱水更难提防", "癸不词讼理弱敌强"}

// PengZuHeavenStem 天干彭祖百忌
type PengZuHeavenStem struct {
	LoopTyme
}

func (PengZuHeavenStem) FromIndex(index int) PengZuHeavenStem {
	return PengZuHeavenStem{LoopTyme{}.FromIndex(PengZuHeavenStemNames, index)}
}

func (PengZuHeavenStem) FromName(name string) (*PengZuHeavenStem, error) {
	p, err := LoopTyme{}.FromName(PengZuHeavenStemNames, name)
	if err != nil {
		return nil, err
	}
	return &PengZuHeavenStem{*p}, nil
}

func (o PengZuHeavenStem) Next(n int) PengZuHeavenStem {
	return o.FromIndex(o.nextIndex(n))
}
