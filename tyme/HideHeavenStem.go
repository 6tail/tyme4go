package tyme

// HideHeavenStem 藏干（即人元，司令取天干，分野取天干的五行）
type HideHeavenStem struct {
	AbstractCulture
	// 天干
	heavenStem HeavenStem
	// 藏干类型
	hideHeavenStemType HideHeavenStemType
}

func (HideHeavenStem) New(heavenStemName string, hideHeavenStemType HideHeavenStemType) (*HideHeavenStem, error) {
	heavenStem, err := HeavenStem{}.FromName(heavenStemName)
	if err != nil {
		return nil, err
	}
	return &HideHeavenStem{
		heavenStem:         *heavenStem,
		hideHeavenStemType: hideHeavenStemType,
	}, nil
}

func (HideHeavenStem) FromIndex(heavenStemIndex int, hideHeavenStemType HideHeavenStemType) HideHeavenStem {
	return HideHeavenStem{
		heavenStem:         HeavenStem{}.FromIndex(heavenStemIndex),
		hideHeavenStemType: hideHeavenStemType,
	}
}

// GetHeavenStem 天干
func (o HideHeavenStem) GetHeavenStem() HeavenStem {
	return o.heavenStem
}

// GetType 藏干类型
func (o HideHeavenStem) GetType() HideHeavenStemType {
	return o.hideHeavenStemType
}

func (o HideHeavenStem) GetName() string {
	return o.heavenStem.GetName()
}

func (o HideHeavenStem) String() string {
	return o.GetName()
}
