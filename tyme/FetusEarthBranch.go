package tyme

// FetusEarthBranchNames 地支六甲胎神名称
var FetusEarthBranchNames = []string{"碓", "厕", "炉", "门", "栖", "床"}

// FetusEarthBranch 地支六甲胎神（《地支六甲胎神歌》子午二日碓须忌，丑未厕道莫修移。寅申火炉休要动，卯酉大门修当避。辰戌鸡栖巳亥床，犯着六甲身堕胎。）
type FetusEarthBranch struct {
	LoopTyme
}

func (FetusEarthBranch) New(index int) FetusEarthBranch {
	return FetusEarthBranch{LoopTyme{}.FromIndex(FetusEarthBranchNames, index)}
}

func (o FetusEarthBranch) Next(n int) FetusEarthBranch {
	return FetusEarthBranch{}.New(o.nextIndex(n))
}
