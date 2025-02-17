package tyme

import "fmt"

// Numbers 数字
var Numbers = []string{"一", "二", "三", "四", "五", "六", "七", "八", "九", "十", "十一", "十二"}

// KitchenGodSteed 灶马头
type KitchenGodSteed struct {
	// 正月初一的干支
	firstDaySixtyCycle SixtyCycle
}

func NewKitchenGodSteed(lunarYear int) KitchenGodSteed {
	lunarDay, _ := LunarDay{}.FromYmd(lunarYear, 1, 1)
	return KitchenGodSteed{
		firstDaySixtyCycle: lunarDay.GetSixtyCycle(),
	}
}

func (KitchenGodSteed) FromLunarYear(lunarYear int) KitchenGodSteed {
	return NewKitchenGodSteed(lunarYear)
}

// 根据天干获取序数
func (o KitchenGodSteed) byHeavenStem(n int) string {
	return Numbers[o.firstDaySixtyCycle.GetHeavenStem().StepsTo(n)]
}

// 根据地支获取序数
func (o KitchenGodSteed) byEarthBranch(n int) string {
	return Numbers[o.firstDaySixtyCycle.GetEarthBranch().StepsTo(n)]
}

// GetMouse 几鼠偷粮
func (o KitchenGodSteed) GetMouse() string {
	return fmt.Sprintf("%s鼠偷粮", o.byEarthBranch(0))
}

// GetGrass 草子几分
func (o KitchenGodSteed) GetGrass() string {
	return fmt.Sprintf("草子%s分", o.byEarthBranch(0))
}

// GetCattle 几牛耕田
func (o KitchenGodSteed) GetCattle() string {
	return fmt.Sprintf("%s牛耕田", o.byEarthBranch(1))
}

// GetFlower 花收几分
func (o KitchenGodSteed) GetFlower() string {
	return fmt.Sprintf("花收%s分", o.byEarthBranch(3))
}

// GetDragon 几龙治水
func (o KitchenGodSteed) GetDragon() string {
	return fmt.Sprintf("%s龙治水", o.byEarthBranch(4))
}

// GetHorse 几马驮谷
func (o KitchenGodSteed) GetHorse() string {
	return fmt.Sprintf("%s马驮谷", o.byEarthBranch(6))
}

// GetChicken 几鸡抢米
func (o KitchenGodSteed) GetChicken() string {
	return fmt.Sprintf("%s鸡抢米", o.byEarthBranch(9))
}

// GetSilkworm 几姑看蚕
func (o KitchenGodSteed) GetSilkworm() string {
	return fmt.Sprintf("%s姑看蚕", o.byEarthBranch(9))
}

// GetPig 几屠共猪
func (o KitchenGodSteed) GetPig() string {
	return fmt.Sprintf("%s屠共猪", o.byEarthBranch(11))
}

// GetField 甲田几分
func (o KitchenGodSteed) GetField() string {
	return fmt.Sprintf("甲田%s分", o.byHeavenStem(0))
}

// GetCake 几人分饼
func (o KitchenGodSteed) GetCake() string {
	return fmt.Sprintf("%s人分饼", o.byHeavenStem(2))
}

// GetGold 几日得金
func (o KitchenGodSteed) GetGold() string {
	return fmt.Sprintf("%s日得金", o.byHeavenStem(7))
}

// GetPeopleCakes 几人几丙
func (o KitchenGodSteed) GetPeopleCakes() string {
	return fmt.Sprintf("%s人%s丙", o.byEarthBranch(2), o.byHeavenStem(2))
}

// GetPeopleHoes 几人几锄
func (o KitchenGodSteed) GetPeopleHoes() string {
	return fmt.Sprintf("%s人%s锄", o.byEarthBranch(2), o.byHeavenStem(3))
}

// Name 名称
func (o KitchenGodSteed) Name() string {
	return "灶马头"
}
