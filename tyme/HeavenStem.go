package tyme

// HeavenStemNames 天干名称
var HeavenStemNames = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}

// HeavenStem 天干（天元）
type HeavenStem struct {
	LoopTyme
}

func (HeavenStem) FromIndex(index int) HeavenStem {
	return HeavenStem{LoopTyme{}.FromIndex(HeavenStemNames, index)}
}

func (HeavenStem) FromName(name string) (*HeavenStem, error) {
	p, err := LoopTyme{}.FromName(HeavenStemNames, name)
	if err != nil {
		return nil, err
	}
	return &HeavenStem{*p}, nil
}

func (o HeavenStem) Next(n int) HeavenStem {
	return o.FromIndex(o.nextIndex(n))
}

// GetElement 五行
func (o HeavenStem) GetElement() Element {
	return Element{}.FromIndex(o.index / 2)
}

// GetYinYang 阴阳
func (o HeavenStem) GetYinYang() YinYang {
	if o.index%2 == 0 {
		return YANG
	}
	return YIN
}

// GetTenStar 十神（生我者，正印偏印。我生者，伤官食神。克我者，正官七杀。我克者，正财偏财。同我者，劫财比肩。）
func (o HeavenStem) GetTenStar(target HeavenStem) TenStar {
	targetIndex := target.GetIndex()
	offset := targetIndex - o.index
	if o.index%2 != 0 && targetIndex%2 == 0 {
		offset += 2
	}
	return TenStar{}.FromIndex(offset)
}

// GetDirection 方位
func (o HeavenStem) GetDirection() Direction {
	return o.GetElement().GetDirection()
}

// GetJoyDirection 喜神方位（《喜神方位歌》甲己在艮乙庚乾，丙辛坤位喜神安。丁壬只在离宫坐，戊癸原在在巽间。）
func (o HeavenStem) GetJoyDirection() Direction {
	return Direction{}.FromIndex([]int{7, 5, 1, 8, 3}[o.index%5])
}

// GetYangDirection 阳贵神方位（《阳贵神歌》甲戊坤艮位，乙己是坤坎，庚辛居离艮，丙丁兑与乾，震巽属何日，壬癸贵神安。）
func (o HeavenStem) GetYangDirection() Direction {
	return Direction{}.FromIndex([]int{1, 1, 6, 5, 7, 0, 8, 7, 2, 3}[o.index])
}

// GetYinDirection 阴贵神方位（《阴贵神歌》甲戊见牛羊，乙己鼠猴乡，丙丁猪鸡位，壬癸蛇兔藏，庚辛逢虎马，此是贵神方。）
func (o HeavenStem) GetYinDirection() Direction {
	return Direction{}.FromIndex([]int{7, 0, 5, 6, 1, 1, 7, 8, 3, 2}[o.index])
}

// GetWealthDirection 财神方位（《财神方位歌》甲乙东北是财神，丙丁向在西南寻，戊己正北坐方位，庚辛正东去安身，壬癸原来正南坐，便是财神方位真。）
func (o HeavenStem) GetWealthDirection() Direction {
	return Direction{}.FromIndex([]int{7, 1, 0, 2, 8}[o.index/2])
}

// GetMascotDirection 福神方位（《福神方位歌》甲乙东南是福神，丙丁正东是堪宜，戊北己南庚辛坤，壬在乾方癸在西。）
func (o HeavenStem) GetMascotDirection() Direction {
	return Direction{}.FromIndex([]int{3, 3, 2, 2, 0, 8, 1, 1, 5, 6}[o.index])
}

// GetPengZuHeavenStem 天干彭祖百忌
func (o HeavenStem) GetPengZuHeavenStem() PengZuHeavenStem {
	return PengZuHeavenStem{}.FromIndex(o.index)
}

// GetTerrain 地势(长生十二神)
func (o HeavenStem) GetTerrain(earthBranch EarthBranch) Terrain {
	earthBranchIndex := earthBranch.GetIndex()
	offset := earthBranchIndex
	if YANG == o.GetYinYang() {
		offset = -offset
	}
	return Terrain{}.FromIndex([]int{1, 6, 10, 9, 10, 9, 7, 0, 4, 3}[o.index] + offset)
}

// GetCombine 五合（甲己合，乙庚合，丙辛合，丁壬合，戊癸合）
func (o HeavenStem) GetCombine() HeavenStem {
	return o.Next(5)
}

// Combine 合化（甲己合化土，乙庚合化金，丙辛合化水，丁壬合化木，戊癸合化火）,如果无法合化，返回nil
func (o HeavenStem) Combine(target HeavenStem) *Element {
	if o.GetCombine().Equals(target) {
		t := Element{}.FromIndex(o.index + 2)
		return &t
	}
	return nil
}

func (o HeavenStem) Equals(target HeavenStem) bool {
	return o.String() == target.String()
}
