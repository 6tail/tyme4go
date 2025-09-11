package tyme

import "strings"

// FetusDay 逐日胎神
type FetusDay struct {
	AbstractCulture
	fetusHeavenStem  FetusHeavenStem
	fetusEarthBranch FetusEarthBranch
	side             Side
	direction        Direction
}

func (FetusDay) New(sixtyCycle SixtyCycle) FetusDay {
	fetusHeavenStem := FetusHeavenStem{}.New(sixtyCycle.GetHeavenStem().GetIndex() % 5)
	fetusEarthBranch := FetusEarthBranch{}.New(sixtyCycle.GetEarthBranch().GetIndex() % 6)
	index := []int{3, 3, 8, 8, 8, 8, 8, 1, 1, 1, 1, 1, 1, 6, 6, 6, 6, 6, 5, 5, 5, 5, 5, 5, 0, 0, 0, 0, 0, -9, -9, -9, -9, -9, -5, -5, -1, -1, -1, -3, -7, -7, -7, -7, -5, 7, 7, 7, 7, 7, 7, 2, 2, 2, 2, 2, 3, 3, 3, 3}[sixtyCycle.GetIndex()]
	side := OUT
	if index < 0 {
		side = IN
	}
	direction := Direction{}.FromIndex(index)

	return FetusDay{
		fetusHeavenStem:  fetusHeavenStem,
		fetusEarthBranch: fetusEarthBranch,
		side:             side,
		direction:        direction,
	}
}

func (FetusDay) FromLunarDay(lunarDay LunarDay) FetusDay {
	return FetusDay{}.New(lunarDay.GetSixtyCycle())
}

func (FetusDay) FromSixtyCycleDay(sixtyCycleDay SixtyCycleDay) FetusDay {
	return FetusDay{}.New(sixtyCycleDay.GetSixtyCycle())
}

// GetSide 内外
func (o FetusDay) GetSide() Side {
	return o.side
}

// GetDirection 方位
func (o FetusDay) GetDirection() Direction {
	return o.direction
}

// GetFetusHeavenStem 天干六甲胎神
func (o FetusDay) GetFetusHeavenStem() FetusHeavenStem {
	return o.fetusHeavenStem
}

// GetFetusEarthBranch 地支六甲胎神
func (o FetusDay) GetFetusEarthBranch() FetusEarthBranch {
	return o.fetusEarthBranch
}

func (o FetusDay) GetName() string {
	s := o.fetusHeavenStem.GetName() + o.fetusEarthBranch.GetName()
	if "门门" == s {
		s = "占大门"
	} else if "碓磨碓" == s {
		s = "占碓磨"
	} else if "房床床" == s {
		s = "占房床"
	} else if strings.HasPrefix(s, "门") {
		s = "占" + s
	}

	s += " "

	if IN == o.side {
		s += "房"
	}
	s += o.side.GetName()

	directionName := o.direction.GetName()
	if OUT == o.side && strings.Contains("北南西东", directionName) {
		s += "正"
	}
	s += directionName
	return s
}

func (o FetusDay) String() string {
	return o.GetName()
}
