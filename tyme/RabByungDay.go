package tyme

import "fmt"

var RabByungDayNames = []string{"初一", "初二", "初三", "初四", "初五", "初六", "初七", "初八", "初九", "初十", "十一", "十二", "十三", "十四", "十五", "十六", "十七", "十八", "十九", "二十", "廿一", "廿二", "廿三", "廿四", "廿五", "廿六", "廿七", "廿八", "廿九", "三十"}

// RabByungDay 藏历日，仅支持藏历1950年十二月初一（公历1951年1月8日）至藏历2050年十二月三十（公历2051年2月11日）
type RabByungDay struct {
	AbstractTyme
	// 藏历月
	month RabByungMonth
	day   int
	leap  bool
}

func NewRabByungDay(month RabByungMonth, day int) (*RabByungDay, error) {
	if day == 0 || day < -30 || day > 30 {
		return nil, fmt.Errorf("illegal day %d in %s", day, month)
	}
	leap := day < 0
	d := day
	if leap {
		d = -d
	}
	missDays := month.GetMissDays()
	leapDays := month.GetLeapDays()
	missDayMap := make(map[int]bool)
	leapDayMap := make(map[int]bool)
	for _, d := range missDays {
		missDayMap[d] = true
	}
	for _, d := range leapDays {
		leapDayMap[d] = true
	}
	if leap {
		if !leapDayMap[d] {
			return nil, fmt.Errorf("illegal leap day %d in %s", d, month)
		}
	} else {
		if _, exists := missDayMap[d]; exists {
			return nil, fmt.Errorf("illegal day %d in %s", d, month)
		}
	}
	return &RabByungDay{month: month, day: d, leap: leap}, nil
}

func (RabByungDay) FromYmd(year, month, day int) (*RabByungDay, error) {
	m, err := RabByungMonth{}.FromYm(year, month)
	if err != nil {
		return nil, err
	}
	return NewRabByungDay(*m, day)
}

func (RabByungDay) FromElementZodiac(rabByungIndex int, element RabByungElement, zodiac Zodiac, month, day int) (*RabByungDay, error) {
	m, err := RabByungMonth{}.FromElementZodiac(rabByungIndex, element, zodiac, month)
	if err != nil {
		return nil, err
	}
	return NewRabByungDay(*m, day)
}

func (RabByungDay) FromSolarDay(solarDay SolarDay) (*RabByungDay, error) {
	start, _ := SolarDay{}.FromYmd(1951, 1, 8)
	days := solarDay.Subtract(*start)
	m, err := RabByungMonth{}.FromYm(1950, 12)
	if err != nil {
		return nil, err
	}
	count := m.GetDayCount()
	for days >= count {
		days -= count
		m, err = m.Next(1)
		if err != nil {
			return nil, err
		}
		count = m.GetDayCount()
	}
	day := days + 1
	for _, d := range m.GetSpecialDays() {
		if d < 0 {
			if day >= -d {
				day++
			}
		} else if d > 0 {
			if day == d+1 {
				day = -d
				break
			} else if day > d+1 {
				day--
			}
		}
	}
	return NewRabByungDay(*m, day)
}

func (o RabByungDay) GetRabByungMonth() RabByungMonth {
	return o.month
}

func (o RabByungDay) GetYear() int {
	return o.month.GetYear()
}

func (o RabByungDay) GetMonth() int {
	return o.month.GetMonthWithLeap()
}

func (o RabByungDay) GetDay() int {
	return o.day
}

func (o RabByungDay) IsLeap() bool {
	return o.leap
}

func (o RabByungDay) GetDayWithLeap() int {
	if o.leap {
		return -o.day
	}
	return o.day
}

func (o RabByungDay) GetName() string {
	name := RabByungDayNames[o.day-1]
	if o.leap {
		return "闰" + name
	}
	return name
}

func (o RabByungDay) String() string {
	return fmt.Sprintf("%s%s", o.month, o.GetName())
}

func (o RabByungDay) Subtract(target RabByungDay) int {
	return o.GetSolarDay().Subtract(target.GetSolarDay())
}

func (o RabByungDay) GetSolarDay() SolarDay {
	start, _ := SolarDay{}.FromYmd(1951, 1, 7)
	m, _ := RabByungMonth{}.FromYm(1950, 12)
	n := 0
	for !o.month.Equals(*m) {
		n += m.GetDayCount()
		m, _ = m.Next(1)
	}
	t := o.day
	for _, sd := range m.GetSpecialDays() {
		if sd < 0 {
			if t > -sd {
				t--
			}
		} else if sd > 0 {
			if t > sd {
				t++
			}
		}
	}
	if o.leap {
		t++
	}
	return start.Next(n + t)
}

func (o RabByungDay) Next(n int) (*RabByungDay, error) {
	return RabByungDay{}.FromSolarDay(o.GetSolarDay().Next(n))
}

func (o RabByungDay) Equals(target RabByungDay) bool {
	return o.String() == target.String()
}
