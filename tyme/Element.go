package tyme

// ElementNames 五行名称
var ElementNames = []string{"木", "火", "土", "金", "水"}

// Element 五行
type Element struct {
	LoopTyme
}

func (Element) FromIndex(index int) Element {
	return Element{LoopTyme{}.FromIndex(ElementNames, index)}
}

func (Element) FromName(name string) (*Element, error) {
	p, err := LoopTyme{}.FromName(ElementNames, name)
	if err != nil {
		return nil, err
	}
	return &Element{*p}, nil
}

func (o Element) Next(n int) Element {
	return o.FromIndex(o.nextIndex(n))
}

// GetReinforce 我生者
func (o Element) GetReinforce() Element {
	return o.Next(1)
}

// GetRestrain 我克者
func (o Element) GetRestrain() Element {
	return o.Next(2)
}

// GetReinforced 生我者
func (o Element) GetReinforced() Element {
	return o.Next(-1)
}

// GetRestrained 克我者
func (o Element) GetRestrained() Element {
	return o.Next(-2)
}

// GetDirection 方位
func (o Element) GetDirection() Direction {
	return Direction{}.FromIndex([]int{2, 8, 4, 6, 0}[o.index])
}
