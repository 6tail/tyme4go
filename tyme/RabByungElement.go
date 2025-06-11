package tyme

import (
	"strings"
)

// RabByungElement 藏历五行
type RabByungElement struct {
	Element
}

func (RabByungElement) FromIndex(index int) RabByungElement {
	return RabByungElement{Element{}.FromIndex(index)}
}

func (RabByungElement) FromName(name string) (*RabByungElement, error) {
	p, err := Element{}.FromName(strings.ReplaceAll(name, "铁", "金"))
	if err != nil {
		return nil, err
	}
	return &RabByungElement{*p}, nil
}

func (o RabByungElement) GetName() string {
	return strings.ReplaceAll(o.Element.GetName(), "金", "铁")
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
