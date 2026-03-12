package tyme

// Gender 性别
type Gender int

const (
	// WOMAN 女
	WOMAN Gender = iota
	// MAN 男
	MAN
)

func (o Gender) GetCode() int {
	return int(o)
}

func NewGender(code int) Gender {
	switch code {
	case 1:
		return MAN
	default:
		return WOMAN
	}
}

func (o Gender) GetName() string {
	switch o {
	case WOMAN:
		return "女"
	default:
		return "男"
	}
}

func (o Gender) String() string {
	return o.GetName()
}
