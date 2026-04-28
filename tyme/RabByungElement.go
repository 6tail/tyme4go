package tyme

// RabByungElementNames 藏历五行名称
var RabByungElementNames = []string{"木", "火", "土", "铁", "水"}

// RabByungElement 藏历五行
type RabByungElement struct {
	Element
}

func (RabByungElement) FromIndex(index int) RabByungElement {
	return RabByungElement{Element{LoopTyme{}.FromIndex(RabByungElementNames, index)}}
}

func (RabByungElement) FromName(name string) (*RabByungElement, error) {
	p, err := LoopTyme{}.FromName(RabByungElementNames, name)
	if err != nil {
		return nil, err
	}
	return &RabByungElement{Element{*p}}, nil
}

func (o RabByungElement) Next(n int) RabByungElement {
	return o.FromIndex(o.nextIndex(n))
}

// GetReinforce 我生者
func (o RabByungElement) GetReinforce() RabByungElement {
	return o.Next(1)
}

// GetRestrain 我克者
func (o RabByungElement) GetRestrain() RabByungElement {
	return o.Next(2)
}

// GetReinforced 生我者
func (o RabByungElement) GetReinforced() RabByungElement {
	return o.Next(-1)
}

// GetRestrained 克我者
func (o RabByungElement) GetRestrained() RabByungElement {
	return o.Next(-2)
}
