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

func (o AbstractCulture) GetName() string {
	return ""
}

func (o AbstractCulture) String() string {
	return o.GetName()
}

func (o AbstractCulture) Equals(target AbstractCulture) bool {
	return o.String() == target.String()
}
