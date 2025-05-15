package tyme

import (
	"fmt"
	"regexp"
	"strconv"
)

// LunarFestivalNames 农历传统节日名称
var LunarFestivalNames = []string{"春节", "元宵节", "龙头节", "上巳节", "清明节", "端午节", "七夕节", "中元节", "中秋节", "重阳节", "冬至节", "腊八节", "除夕"}

// LunarFestivalData 农历传统节日数据
var LunarFestivalData = "@0000101@0100115@0200202@0300303@04107@0500505@0600707@0700715@0800815@0900909@10124@1101208@122"

// LunarFestival 农历传统节日（依据国家标准《农历的编算和颁行》GB/T 33661-2017）
type LunarFestival struct {
	AbstractTyme
	// 类型
	festivalType FestivalType
	// 序号
	index int
	// 农历日
	day LunarDay
	// 节气
	solarTerm *SolarTerm
	// 名称
	name string
}

func (LunarFestival) New(festivalType FestivalType, day LunarDay, solarTerm *SolarTerm, data string) (*LunarFestival, error) {
	index, err := strconv.Atoi(data[1:3])
	if err != nil {
		return nil, err
	}
	return &LunarFestival{
		festivalType: festivalType,
		day:          day,
		solarTerm:    solarTerm,
		index:        index,
		name:         LunarFestivalNames[index],
	}, nil
}

func (LunarFestival) FromIndex(year int, index int) (*LunarFestival, error) {
	if index < 0 || index >= len(LunarFestivalNames) {
		return nil, fmt.Errorf(fmt.Sprintf("illegal index: %d", index))
	}
	re, err := regexp.Compile(fmt.Sprintf("@%02d\\d+", index))
	if err != nil {
		return nil, err
	}
	data := re.FindString(LunarFestivalData)
	if data == "" {
		return nil, nil
	}
	t := NewFestivalType(int([]rune(data[3:4])[0] - '0'))
	switch t {
	case DAY:
		month, err := strconv.Atoi(data[4:6])
		if err != nil {
			return nil, err
		}
		day, err := strconv.Atoi(data[6:8])
		if err != nil {
			return nil, err
		}
		d, _ := LunarDay{}.FromYmd(year, month, day)
		f, err := LunarFestival{}.New(t, *d, nil, data)
		if err != nil {
			return nil, err
		}
		return f, nil
	case TERM:
		i, err := strconv.Atoi(data[4:])
		if err != nil {
			return nil, err
		}
		solarTerm := SolarTerm{}.FromIndex(year, i)
		f, err := LunarFestival{}.New(t, solarTerm.GetJulianDay().GetSolarDay().GetLunarDay(), &solarTerm, data)
		if err != nil {
			return nil, err
		}
		return f, nil
	case EVE:
		d, _ := LunarDay{}.FromYmd(year+1, 1, 1)
		f, err := LunarFestival{}.New(t, d.Next(-1), nil, data)
		if err != nil {
			return nil, err
		}
		return f, nil
	}
	return nil, nil
}

func (LunarFestival) FromYmd(year int, month int, day int) (*LunarFestival, error) {
	re, err := regexp.Compile(fmt.Sprintf("@\\d{2}0%02d%02d", month, day))
	if err != nil {
		return nil, err
	}
	data := re.FindString(LunarFestivalData)
	if data != "" {
		d, err := LunarDay{}.FromYmd(year, month, day)
		if err != nil {
			return nil, err
		}
		f, err := LunarFestival{}.New(DAY, *d, nil, data)
		if err != nil {
			return nil, err
		}
		return f, nil
	}
	re, err = regexp.Compile("@\\d{2}1\\d{2}")
	if err != nil {
		return nil, err
	}
	arr := re.FindAllString(LunarFestivalData, -1)
	for _, data := range arr {
		i, err := strconv.Atoi(data[4:])
		if err != nil {
			return nil, err
		}
		solarTerm := SolarTerm{}.FromIndex(year, i)
		d := solarTerm.GetJulianDay().GetSolarDay().GetLunarDay()
		if d.GetYear() == year && d.GetMonth() == month && d.GetDay() == day {
			f, err := LunarFestival{}.New(TERM, d, &solarTerm, data)
			if err != nil {
				return nil, err
			}
			return f, nil
		}
	}

	re, err = regexp.Compile("@\\d{2}2")
	if err != nil {
		return nil, err
	}
	data = re.FindString(LunarFestivalData)
	if data == "" {
		return nil, nil
	}
	d, err := LunarDay{}.FromYmd(year, month, day)
	if err != nil {
		return nil, err
	}
	nextDay := d.Next(1)
	if nextDay.GetMonth() == 1 && nextDay.GetDay() == 1 {
		f, err := LunarFestival{}.New(EVE, *d, nil, data)
		if err != nil {
			return nil, err
		}
		return f, nil
	}
	return nil, nil
}

// GetType 节日类型
func (o LunarFestival) GetType() FestivalType {
	return o.festivalType
}

// GetDay 农历日
func (o LunarFestival) GetDay() LunarDay {
	return o.day
}

// GetIndex 索引
func (o LunarFestival) GetIndex() int {
	return o.index
}

func (o LunarFestival) GetName() string {
	return o.name
}

// GetSolarTerm 节气
func (o LunarFestival) GetSolarTerm() *SolarTerm {
	return o.solarTerm
}

func (o LunarFestival) Next(n int) *LunarFestival {
	size := len(LunarFestivalNames)
	i := o.index + n
	f, _ := LunarFestival{}.FromIndex((o.day.GetYear()*size+i)/size, o.IndexOf(i, size))
	return f
}

func (o LunarFestival) String() string {
	return fmt.Sprintf("%v %v", o.day, o.name)
}
