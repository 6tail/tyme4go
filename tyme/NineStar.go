package tyme

import "fmt"

// NineStarNames 九星名称
var NineStarNames = []string{"一", "二", "三", "四", "五", "六", "七", "八", "九"}

// NineStarColors 九星颜色
var NineStarColors = []string{"白", "黑", "碧", "绿", "黄", "白", "赤", "白", "紫"}

// NineStar 九星
type NineStar struct {
	LoopTyme
}

func (NineStar) FromIndex(index int) NineStar {
	return NineStar{LoopTyme{}.FromIndex(NineStarNames, index)}
}

func (NineStar) FromName(name string) (*NineStar, error) {
	p, err := LoopTyme{}.FromName(NineStarNames, name)
	if err != nil {
		return nil, err
	}
	return &NineStar{*p}, nil
}

func (o NineStar) Next(n int) NineStar {
	return o.FromIndex(o.nextIndex(n))
}

// GetColor 颜色
func (o NineStar) GetColor() string {
	return NineStarColors[o.index]
}

// GetElement 五行
func (o NineStar) GetElement() Element {
	return o.GetDirection().GetElement()
}

// GetDipper 北斗九星
func (o NineStar) GetDipper() Dipper {
	return Dipper{}.FromIndex(o.index)
}

// GetDirection 方位
func (o NineStar) GetDirection() Direction {
	return Direction{}.FromIndex(o.index)
}

func (o NineStar) String() string {
	return fmt.Sprintf("%v%v%v", o.GetName(), o.GetColor(), o.GetElement())
}
