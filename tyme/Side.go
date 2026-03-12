package tyme

// Side 内外
type Side int

const (
	// IN 内
	IN Side = iota
	// OUT 外
	OUT
)

func (o Side) GetCode() int {
	return int(o)
}

func NewSide(code int) Side {
	switch code {
	case 1:
		return OUT
	default:
		return IN
	}
}

func (o Side) GetName() string {
	switch o {
	case IN:
		return "内"
	default:
		return "外"
	}
}

func (o Side) String() string {
	return o.GetName()
}
