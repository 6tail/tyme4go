package tyme

// Fortune 小运
type Fortune struct {
	AbstractTyme
	// 童限
	childLimit ChildLimit
	// 序号
	index int
}

func (Fortune) FromChildLimit(childLimit ChildLimit, index int) Fortune {
	return Fortune{
		childLimit: childLimit,
		index:      index,
	}
}

// GetAge 年龄
func (o Fortune) GetAge() int {
	return o.childLimit.GetYearCount() + 1 + o.index
}

// GetLunarYear 农历年
func (o Fortune) GetLunarYear() LunarYear {
	return o.childLimit.GetEndTime().GetLunarHour().GetLunarDay().GetLunarMonth().GetLunarYear().Next(o.index)
}

// GetSixtyCycle 干支
func (o Fortune) GetSixtyCycle() SixtyCycle {
	n := o.GetAge()
	if !o.childLimit.IsForward() {
		n = -n
	}
	return o.childLimit.GetEightChar().GetHour().Next(n)
}

func (o Fortune) GetName() string {
	return o.GetSixtyCycle().GetName()
}

func (o Fortune) String() string {
	return o.GetName()
}

func (o Fortune) Next(n int) Fortune {
	return Fortune{}.FromChildLimit(o.childLimit, o.index+n)
}
