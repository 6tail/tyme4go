package tyme

// DecadeFortune 大运（10年1大运）
type DecadeFortune struct {
	AbstractTyme
	// 童限
	childLimit ChildLimit
	// 序号
	index int
}

func (DecadeFortune) FromChildLimit(childLimit ChildLimit, index int) DecadeFortune {
	return DecadeFortune{
		childLimit: childLimit,
		index:      index,
	}
}

// GetStartAge 开始年龄
func (o DecadeFortune) GetStartAge() int {
	return o.childLimit.GetYearCount() + 1 + o.index*10
}

// GetEndAge 结束年龄
func (o DecadeFortune) GetEndAge() int {
	return o.GetStartAge() + 9
}

// GetStartLunarYear 开始农历年
func (o DecadeFortune) GetStartLunarYear() LunarYear {
	return o.childLimit.GetEndTime().GetLunarHour().GetLunarDay().GetLunarMonth().GetLunarYear().Next(o.index * 10)
}

// GetEndLunarYear 结束农历年
func (o DecadeFortune) GetEndLunarYear() LunarYear {
	return o.GetStartLunarYear().Next(9)
}

// GetSixtyCycle 干支
func (o DecadeFortune) GetSixtyCycle() SixtyCycle {
	n := o.index + 1
	if !o.childLimit.IsForward() {
		n = -n
	}
	return o.childLimit.GetEightChar().GetMonth().Next(n)
}

func (o DecadeFortune) GetName() string {
	return o.GetSixtyCycle().GetName()
}

func (o DecadeFortune) String() string {
	return o.GetName()
}

func (o DecadeFortune) Next(n int) DecadeFortune {
	return DecadeFortune{}.FromChildLimit(o.childLimit, o.index+n)
}

// GetStartFortune 开始小运
func (o DecadeFortune) GetStartFortune() Fortune {
	return Fortune{}.FromChildLimit(o.childLimit, o.index*10)
}
