package tyme

// Gender 性别
type Gender int

const (
	// WOMAN 女
	WOMAN Gender = iota
	// MAN 男
	MAN
)

func (d Gender) GetCode() int {
	return int(d)
}

func NewGender(code int) Gender {
	switch code {
	case 1:
		return MAN
	default:
		return WOMAN
	}
}

func (d Gender) GetName() string {
	switch d {
	case WOMAN:
		return "女"
	default:
		return "男"
	}
}

func (d Gender) String() string {
	return d.GetName()
}
