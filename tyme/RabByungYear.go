package tyme

import (
	"fmt"
	"strings"
)

// RabByungYear 藏历年(公历1027年为藏历元年，第一饶迥火兔年）
type RabByungYear struct {
	AbstractTyme
	// 饶迥(胜生周)序号，从0开始
	rabByungIndex int
	// 五行索引，从0开始
	elementIndex int
	// 生肖索引，从0开始
	zodiacIndex int
}

func (RabByungYear) Validate(year int) error {
	if year < 1027 || year > 9999 {
		return fmt.Errorf(fmt.Sprintf("illegal rab-byung year: %d", year))
	}
	return nil
}

func NewRabByungYear(rabByungIndex int, elementIndex int, zodiacIndex int) (*RabByungYear, error) {
	if rabByungIndex < 0 || rabByungIndex > 150 {
		return nil, fmt.Errorf("illegal rab-byung index: %d", rabByungIndex)
	}
	if elementIndex < 0 || elementIndex >= len(ElementNames) {
		return nil, fmt.Errorf("illegal element index: %d", elementIndex)
	}
	if zodiacIndex < 0 || zodiacIndex >= len(ZodiacNames) {
		return nil, fmt.Errorf("illegal zodiac index: %d", zodiacIndex)
	}
	return &RabByungYear{rabByungIndex: rabByungIndex, elementIndex: elementIndex, zodiacIndex: zodiacIndex}, nil
}

func (RabByungYear) FromSixtyCycle(rabByungIndex int, sixtyCycle SixtyCycle) (*RabByungYear, error) {
	return NewRabByungYear(rabByungIndex, sixtyCycle.GetHeavenStem().GetElement().index, sixtyCycle.GetEarthBranch().GetZodiac().index)
}

func (RabByungYear) FromElementZodiac(rabByungIndex int, element RabByungElement, zodiac Zodiac) (*RabByungYear, error) {
	return NewRabByungYear(rabByungIndex, element.index, zodiac.index)
}

func (RabByungYear) FromYear(year int) (*RabByungYear, error) {
	err := RabByungYear{}.Validate(year)
	if err != nil {
		return nil, err
	}
	return RabByungYear{}.FromSixtyCycle((year-1024)/60, SixtyCycle{}.FromIndex(year-4))
}

func (y RabByungYear) GetRabByungIndex() int {
	return y.rabByungIndex
}

func (y RabByungYear) GetSixtyCycle() SixtyCycle {
	return SixtyCycle{}.FromIndex(6*(y.elementIndex*2+y.zodiacIndex%2) - 5*y.zodiacIndex)
}

func (y RabByungYear) GetZodiac() Zodiac {
	return Zodiac{}.FromIndex(y.zodiacIndex)
}

func (y RabByungYear) GetElement() RabByungElement {
	return RabByungElement{}.FromIndex(y.elementIndex)
}

func (y RabByungYear) GetName() string {
	digits := []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	units := []string{"", "十", "百"}
	n := y.rabByungIndex + 1
	s := ""
	pos := 0
	for n > 0 {
		digit := n % 10
		if digit > 0 {
			s = digits[digit] + units[pos] + s
		} else if s != "" {
			s = digits[digit] + s
		}
		n /= 10
		pos++
	}
	if strings.HasPrefix(s, "一十") {
		s = s[3:]
	}
	return fmt.Sprintf("第%s饶迥%s%s年", s, y.GetElement().GetName(), y.GetZodiac().GetName())
}

func (y RabByungYear) Next(n int) (*RabByungYear, error) {
	year, err := RabByungYear{}.FromYear(y.GetYear() + n)
	if err != nil {
		return nil, err
	}
	return year, nil
}

func (y RabByungYear) GetYear() int {
	return 1024 + y.rabByungIndex*60 + y.GetSixtyCycle().GetIndex()
}

func (y RabByungYear) GetLeapMonth() int {
	n := 1
	m := 4
	t := 1
	currentYear := y.GetYear()
	for n < currentYear {
		i := m + 31 + t
		n += 2
		m = i - 23
		if i > 35 {
			n += 1
			m -= 12
		}
		t = 1 - t
	}
	if n == currentYear {
		return m
	}
	return 0
}

func (y RabByungYear) GetSolarYear() SolarYear {
	year, _ := SolarYear{}.FromYear(y.GetYear())
	return *year
}

func (y RabByungYear) GetFirstMonth() RabByungMonth {
	m, _ := NewRabByungMonth(y.GetYear(), 1)
	return *m
}

func (y RabByungYear) GetMonthCount() int {
	if y.GetLeapMonth() < 1 {
		return 12
	}
	return 13
}

func (y RabByungYear) GetMonths() []RabByungMonth {
	var months []RabByungMonth
	leapMonth := y.GetLeapMonth()
	year := y.GetYear()
	for i := 1; i < 13; i++ {
		m, _ := NewRabByungMonth(year, i)
		months = append(months, *m)
		if i == leapMonth {
			m, _ = NewRabByungMonth(year, -i)
			months = append(months, *m)
		}
	}
	return months
}

func (y RabByungYear) Equals(target RabByungYear) bool {
	return y.String() == target.String()
}
