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
	// 干支
	sixtyCycle SixtyCycle
}

func NewRabByungYear(rabByungIndex int, sixtyCycle SixtyCycle) (*RabByungYear, error) {
	if rabByungIndex < 0 || rabByungIndex > 150 {
		return nil, fmt.Errorf("illegal rab-byung index: %d", rabByungIndex)
	}
	return &RabByungYear{rabByungIndex: rabByungIndex, sixtyCycle: sixtyCycle}, nil
}

func (RabByungYear) FromSixtyCycle(rabByungIndex int, sixtyCycle SixtyCycle) (*RabByungYear, error) {
	return NewRabByungYear(rabByungIndex, sixtyCycle)
}

func (RabByungYear) FromElementZodiac(rabByungIndex int, element RabByungElement, zodiac Zodiac) (*RabByungYear, error) {
	for i := 0; i < 60; i++ {
		sc := SixtyCycle{}.FromIndex(i)
		if sc.GetEarthBranch().GetZodiac().Equals(zodiac) && sc.GetHeavenStem().GetElement().GetIndex() == element.GetIndex() {
			return NewRabByungYear(rabByungIndex, sc)
		}
	}
	return nil, fmt.Errorf("no matching sixty cycle for element %s and zodiac %s", element.GetName(), zodiac.GetName())
}

func (RabByungYear) FromYear(year int) (*RabByungYear, error) {
	return NewRabByungYear((year-1024)/60, SixtyCycle{}.FromIndex(year-4))
}

func (y RabByungYear) GetRabByungIndex() int {
	return y.rabByungIndex
}

func (y RabByungYear) GetSixtyCycle() SixtyCycle {
	return y.sixtyCycle
}

func (y RabByungYear) GetZodiac() Zodiac {
	return y.sixtyCycle.GetEarthBranch().GetZodiac()
}

func (y RabByungYear) GetElement() RabByungElement {
	return RabByungElement{}.FromIndex(y.sixtyCycle.GetHeavenStem().GetElement().GetIndex())
}

func (y *RabByungYear) GetName() string {
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
	return 1024 + y.rabByungIndex*60 + y.sixtyCycle.GetIndex()
}

func (y RabByungYear) GetLeapMonth() int {
	yVal := 1
	m := 4
	t := 0
	currentYear := y.GetYear()
	for yVal < currentYear {
		i := m - 1
		if t%2 == 0 {
			i += 33
		} else {
			i += 32
		}
		yVal = (yVal*12 + i) / 12
		m = i%12 + 1
		t++
	}
	if yVal == currentYear {
		return m
	}
	return 0
}

func (y RabByungYear) GetSolarYear() SolarYear {
	year, _ := SolarYear{}.FromYear(y.GetYear())
	return *year
}

func (y RabByungYear) GetFirstMonth() RabByungMonth {
	m, _ := NewRabByungMonth(y, 1)
	return *m
}

func (y RabByungYear) GetMonthCount() int {
	if y.GetLeapMonth() == 0 {
		return 12
	}
	return 13
}

func (y RabByungYear) GetMonths() []RabByungMonth {
	var months []RabByungMonth
	leapMonth := y.GetLeapMonth()
	for i := 1; i <= 12; i++ {
		m, _ := NewRabByungMonth(y, i)
		months = append(months, *m)
		if i == leapMonth {
			lm, _ := NewRabByungMonth(y, -i)
			months = append(months, *lm)
		}
	}
	return months
}

func (o RabByungYear) Equals(target RabByungYear) bool {
	return o.String() == target.String()
}
