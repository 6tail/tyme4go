package tyme

// FetusHeavenStemNames 天干六甲胎神名称
var FetusHeavenStemNames = []string{"门", "碓磨", "厨灶", "仓库", "房床"}

// FetusHeavenStem 天干六甲胎神（《天干六甲胎神歌》甲己之日占在门，乙庚碓磨休移动。丙辛厨灶莫相干，丁壬仓库忌修弄。戊癸房床若移整，犯之孕妇堕孩童。）
type FetusHeavenStem struct {
	LoopTyme
}

func (FetusHeavenStem) New(index int) FetusHeavenStem {
	return FetusHeavenStem{LoopTyme{}.FromIndex(FetusHeavenStemNames, index)}
}

func (o FetusHeavenStem) Next(n int) FetusHeavenStem {
	return FetusHeavenStem{}.New(o.nextIndex(n))
}
