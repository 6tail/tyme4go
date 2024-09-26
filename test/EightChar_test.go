package test

import (
	"github.com/6tail/tyme4go/tyme"
	"testing"
)

// TestEightChar0 八字测试
func TestEightChar0(t *testing.T) {
	// 使用元亨利贞的计算方式
	// tyme.ChildLimitInfoProvider = tyme.China95ChildLimitProvider{}

	eightChar, _ := tyme.EightChar{}.New("癸卯", "辛酉", "己亥", "癸酉")
	taiYuan := eightChar.GetFetalOrigin()

	excepted := "壬子"
	got := taiYuan.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "桑柘木"
	got = taiYuan.GetSound().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestEightChar1(t *testing.T) {
	time, _ := tyme.SolarTime{}.FromYmdHms(2005, 12, 23, 8, 37, 0)
	eightChar := time.GetLunarHour().GetEightChar()

	excepted := "乙酉"
	got := eightChar.GetYear().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "戊子"
	got = eightChar.GetMonth().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "辛巳"
	got = eightChar.GetDay().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "壬辰"
	got = eightChar.GetHour().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestEightChar3(t *testing.T) {
	time, _ := tyme.SolarTime{}.FromYmdHms(2022, 3, 9, 20, 51, 0)
	childLimit := tyme.ChildLimit{}.FromSolarTime(*time, tyme.MAN)

	excepted := 8
	got := childLimit.GetYearCount()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = 9
	got = childLimit.GetMonthCount()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = 2
	got = childLimit.GetDayCount()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = 10
	got = childLimit.GetHourCount()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = 26
	got = childLimit.GetMinuteCount()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted1 := "2030年12月12日 07:17:00"
	got1 := childLimit.GetEndTime().String()
	if excepted1 != got1 {
		t.Errorf("excepted: %v, got: %v", excepted1, got1)
	}
}
