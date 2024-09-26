package tyme

import "fmt"

// PengZu 彭祖百忌
type PengZu struct {
	AbstractCulture,
	pengZuHeavenStem PengZuHeavenStem
	pengZuEarthBranch PengZuEarthBranch
}

func (PengZu) FromSixtyCycle(sixtyCycle SixtyCycle) PengZu {
	return PengZu{
		pengZuHeavenStem:  PengZuHeavenStem{}.FromIndex(sixtyCycle.GetHeavenStem().GetIndex()),
		pengZuEarthBranch: PengZuEarthBranch{}.FromIndex(sixtyCycle.GetEarthBranch().GetIndex()),
	}
}

func (o PengZu) GetName() string {
	return fmt.Sprintf("%v %v", o.pengZuHeavenStem, o.pengZuEarthBranch)
}

func (o PengZu) String() string {
	return o.GetName()
}

// GetPengZuHeavenStem 天干彭祖百忌
func (o PengZu) GetPengZuHeavenStem() PengZuHeavenStem {
	return o.pengZuHeavenStem
}

// GetPengZuEarthBranch 地支彭祖百忌
func (o PengZu) GetPengZuEarthBranch() PengZuEarthBranch {
	return o.pengZuEarthBranch
}
