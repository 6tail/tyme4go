package tyme

// FetusMonthNames 逐月胎神名称
var FetusMonthNames = []string{"占房床", "占户窗", "占门堂", "占厨灶", "占房床", "占床仓", "占碓磨", "占厕户", "占门房", "占房床", "占灶炉", "占房床"}

// FetusMonth 逐月胎神（正十二月在床房，二三九十门户中，四六十一灶勿犯，五甲七子八厕凶。）
type FetusMonth struct {
	LoopTyme
}

func (FetusMonth) New(index int) FetusMonth {
	return FetusMonth{LoopTyme{}.FromIndex(FetusMonthNames, index)}
}

// FromLunarMonth 从农历月初始化
func (FetusMonth) FromLunarMonth(lunarMonth LunarMonth) *FetusMonth {
	if lunarMonth.IsLeap() {
		return nil
	}
	m := FetusMonth{}.New(lunarMonth.GetMonth() - 1)
	return &m
}

func (o FetusMonth) Next(n int) FetusMonth {
	return FetusMonth{}.New(o.nextIndex(n))
}
