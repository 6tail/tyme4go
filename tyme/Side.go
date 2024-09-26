package tyme

// Side 内外
type Side int

const (
	// IN 内
	IN Side = iota
	// OUT 外
	OUT
)

func (d Side) GetCode() int {
	return int(d)
}

func NewSide(code int) Side {
	switch code {
	case 1:
		return OUT
	default:
		return IN
	}
}

func (d Side) GetName() string {
	switch d {
	case IN:
		return "内"
	default:
		return "外"
	}
}

func (d Side) String() string {
	return d.GetName()
}
