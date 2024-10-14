package tyme

// MinorRenNames 小六壬名称
var MinorRenNames = []string{"大安", "留连", "速喜", "赤口", "小吉", "空亡"}

// MinorRen 小六壬
type MinorRen struct {
	LoopTyme
}

func (MinorRen) FromIndex(index int) MinorRen {
	return MinorRen{LoopTyme{}.FromIndex(MinorRenNames, index)}
}

func (MinorRen) FromName(name string) (*MinorRen, error) {
	p, err := LoopTyme{}.FromName(MinorRenNames, name)
	if err != nil {
		return nil, err
	}
	return &MinorRen{*p}, nil
}

func (o MinorRen) Next(n int) MinorRen {
	return o.FromIndex(o.nextIndex(n))
}

// GetLuck 吉凶
func (o MinorRen) GetLuck() Luck {
	return Luck{}.FromIndex(o.index % 2)
}

// GetElement 五行
func (o MinorRen) GetElement() Element {
	return Element{}.FromIndex([]int{0, 4, 1, 3, 0, 2}[o.index])
}
