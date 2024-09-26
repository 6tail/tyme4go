package tyme

// TwentyEightStarNames 二十八宿名称
var TwentyEightStarNames = []string{"角", "亢", "氐", "房", "心", "尾", "箕", "斗", "牛", "女", "虚", "危", "室", "壁", "奎", "娄", "胃", "昴", "毕", "觜", "参", "井", "鬼", "柳", "星", "张", "翼", "轸"}

// TwentyEightStar 二十八宿
type TwentyEightStar struct {
	LoopTyme
}

func (TwentyEightStar) FromIndex(index int) TwentyEightStar {
	return TwentyEightStar{LoopTyme{}.FromIndex(TwentyEightStarNames, index)}
}

func (TwentyEightStar) FromName(name string) (*TwentyEightStar, error) {
	p, err := LoopTyme{}.FromName(TwentyEightStarNames, name)
	if err != nil {
		return nil, err
	}
	return &TwentyEightStar{*p}, nil
}

func (o TwentyEightStar) Next(n int) TwentyEightStar {
	return o.FromIndex(o.nextIndex(n))
}

// GetSevenStar 七曜
func (o TwentyEightStar) GetSevenStar() SevenStar {
	return SevenStar{}.FromIndex(o.index%7 + 4)
}

// GetLand 九野
func (o TwentyEightStar) GetLand() Land {
	return Land{}.FromIndex([]int{4, 4, 4, 2, 2, 2, 7, 7, 7, 0, 0, 0, 0, 5, 5, 5, 6, 6, 6, 1, 1, 1, 8, 8, 8, 3, 3, 3}[o.index])
}

// GetZone 宫
func (o TwentyEightStar) GetZone() Zone {
	return Zone{}.FromIndex(o.index / 7)
}

// GetAnimal 动物
func (o TwentyEightStar) GetAnimal() Animal {
	return Animal{}.FromIndex(o.index)
}

// GetLuck 吉凶
func (o TwentyEightStar) GetLuck() Luck {
	return Luck{}.FromIndex([]int{0, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 1, 1, 0, 1, 0}[o.index])
}
