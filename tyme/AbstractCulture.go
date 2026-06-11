package tyme

// AbstractCulture 传统文化抽象
type AbstractCulture struct {
	Culture
}

// IndexOf 转换为不超范围的索引
func (o AbstractCulture) IndexOf(index int, size int) int {
	i := index % size
	if i < 0 {
		i += size
	}
	return i
}

// FloorDiv 向下取整的除法（向负无穷取整）
func (o AbstractCulture) FloorDiv(x int, y int) int {
	r := x / y
	if (x^y) < 0 && (r*y != x) {
		r--
	}
	return r
}

func (o AbstractCulture) GetName() string {
	return ""
}

func (o AbstractCulture) String() string {
	return o.GetName()
}

func (o AbstractCulture) Equals(target AbstractCulture) bool {
	return o.String() == target.String()
}
