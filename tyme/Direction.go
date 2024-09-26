package tyme

// DirectionNames 方位名称，依据后天八卦排序（0坎北, 1坤西南, 2震东, 3巽东南, 4中, 5乾西北, 6兑西, 7艮东北, 8离南）
var DirectionNames = []string{"北", "西南", "东", "东南", "中", "西北", "西", "东北", "南"}

// Direction 方位
type Direction struct {
	LoopTyme
}

func (Direction) FromIndex(index int) Direction {
	return Direction{LoopTyme{}.FromIndex(DirectionNames, index)}
}

func (Direction) FromName(name string) (*Direction, error) {
	p, err := LoopTyme{}.FromName(DirectionNames, name)
	if err != nil {
		return nil, err
	}
	return &Direction{*p}, nil
}

func (o Direction) Next(n int) Direction {
	return o.FromIndex(o.nextIndex(n))
}

// GetLand 九野
func (o Direction) GetLand() Land {
	return Land{}.FromIndex(o.index)
}

// GetElement 五行
func (o Direction) GetElement() Element {
	return Element{}.FromIndex([]int{4, 2, 0, 0, 2, 3, 3, 2, 1}[o.index])
}
