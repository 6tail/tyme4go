package tyme

import (
	"fmt"
	"regexp"
	"strconv"
)

// SolarFestivalNames 公历现代节日名称
var SolarFestivalNames = []string{"元旦", "三八妇女节", "植树节", "五一劳动节", "五四青年节", "六一儿童节", "建党节", "八一建军节", "教师节", "国庆节"}

// SolarFestivalData 公历现代节日数据
var SolarFestivalData = "@00001011950@01003081950@02003121979@03005011950@04005041950@05006011950@06007011941@07008011933@08009101985@09010011950"

// SolarFestival 公历现代节日
type SolarFestival struct {
	AbstractTyme
	// 类型
	festivalType FestivalType
	// 序号
	index int
	// 公历日
	day SolarDay
	// 名称
	name string
	// 起始年
	startYear int
}

func (SolarFestival) New(festivalType FestivalType, day SolarDay, startYear int, data string) (*SolarFestival, error) {
	index, err := strconv.Atoi(data[1:3])
	if err != nil {
		return nil, err
	}
	return &SolarFestival{
		festivalType: festivalType,
		day:          day,
		startYear:    startYear,
		index:        index,
		name:         SolarFestivalNames[index],
	}, nil
}

func (SolarFestival) FromIndex(year int, index int) (*SolarFestival, error) {
	if index < 0 || index >= len(SolarFestivalNames) {
		return nil, nil
	}
	re, err := regexp.Compile(fmt.Sprintf("@%02d\\d+", index))
	if err != nil {
		return nil, err
	}
	data := re.FindString(SolarFestivalData)
	if data == "" {
		return nil, nil
	}
	t := NewFestivalType(int([]rune(data[3:4])[0] - '0'))
	if t != DAY {
		return nil, nil
	}
	startYear, err := strconv.Atoi(data[8:])
	if err != nil {
		return nil, err
	}
	if year < startYear {
		return nil, nil
	}
	month, err := strconv.Atoi(data[4:6])
	if err != nil {
		return nil, err
	}
	day, err := strconv.Atoi(data[6:8])
	if err != nil {
		return nil, err
	}
	d, _ := SolarDay{}.FromYmd(year, month, day)
	f, err := SolarFestival{}.New(t, *d, startYear, data)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (SolarFestival) FromYmd(year int, month int, day int) (*SolarFestival, error) {
	re, err := regexp.Compile(fmt.Sprintf("@\\d{2}0%02d%02d\\d+", month, day))
	if err != nil {
		return nil, err
	}
	data := re.FindString(SolarFestivalData)
	if data == "" {
		return nil, nil
	}
	startYear, err := strconv.Atoi(data[8:])
	if err != nil {
		return nil, err
	}
	if year < startYear {
		return nil, nil
	}
	d, err := SolarDay{}.FromYmd(year, month, day)
	if err != nil {
		return nil, err
	}
	f, err := SolarFestival{}.New(DAY, *d, startYear, data)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// GetType 节日类型
func (o SolarFestival) GetType() FestivalType {
	return o.festivalType
}

// GetDay 公历日
func (o SolarFestival) GetDay() SolarDay {
	return o.day
}

// GetIndex 索引
func (o SolarFestival) GetIndex() int {
	return o.index
}

func (o SolarFestival) GetName() string {
	return o.name
}

// GetStartYear 起始年
func (o SolarFestival) GetStartYear() int {
	return o.startYear
}

func (o SolarFestival) Next(n int) *SolarFestival {
	size := len(SolarFestivalNames)
	i := o.index + n
	f, _ := SolarFestival{}.FromIndex((o.day.GetYear()*size+i)/size, o.IndexOf(i, size))
	return f
}

func (o SolarFestival) String() string {
	return fmt.Sprintf("%v %v", o.day, o.name)
}
