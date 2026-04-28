package tyme

import "fmt"

// LoopTyme 可轮回的Tyme
type LoopTyme struct {
	AbstractTyme

	// 名称列表
	names []string

	// 索引，从0开始
	index int
}

func (o LoopTyme) GetIndex() int {
	return o.index
}

func (o LoopTyme) GetSize() int {
	return len(o.names)
}

func (o LoopTyme) indexOfName(name string) (int, error) {
	for i, v := range o.names {
		if v == name {
			return i, nil
		}
	}
	return -1, fmt.Errorf("illegal name: " + name)
}

func (o LoopTyme) indexOf(index int) int {
	return o.IndexOf(index, o.GetSize())
}

func (o LoopTyme) nextIndex(n int) int {
	return o.indexOf(o.index + n)
}

func (o LoopTyme) StepsTo(targetIndex int) int {
	return o.indexOf(targetIndex - o.index)
}

func (o LoopTyme) StepsBackTo(targetIndex int) int {
	n := o.GetSize()
	return -((o.index - targetIndex + n) % n)
}

func (o LoopTyme) StepsCloseTo(targetIndex int) int {
	d1 := o.StepsTo(targetIndex)
	d2 := o.StepsBackTo(targetIndex)
	t := d2
	if t < 0 {
		t = -t
	}
	if d1 <= t {
		return d1
	}
	return d2
}

func (LoopTyme) FromIndex(names []string, index int) LoopTyme {
	instance := LoopTyme{names: names}
	instance.index = instance.nextIndex(index)
	return instance
}

func (LoopTyme) FromName(names []string, name string) (*LoopTyme, error) {
	instance := LoopTyme{names: names}
	index, err := instance.indexOfName(name)
	if err != nil {
		return nil, err
	}
	instance.index = index
	return &instance, nil
}

func (o LoopTyme) GetName() string {
	return o.names[o.index]
}

func (o LoopTyme) String() string {
	return o.GetName()
}

func (o LoopTyme) Equals(target LoopTyme) bool {
	return o.String() == target.String()
}
