package test

import (
	"github.com/6tail/tyme4go/tyme"
	"testing"
)

// TestRabByungYear0 藏历年测试
func TestRabByungYear0(t *testing.T) {
	e, _ := tyme.RabByungElement{}.FromName("火")
	z, _ := tyme.Zodiac{}.FromName("兔")
	y, _ := tyme.RabByungYear{}.FromElementZodiac(0, *e, *z)
	excepted := "第一饶迥火兔年"
	got := y.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "1027年"
	got = y.GetSolarYear().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "丁卯"
	got = y.GetSixtyCycle().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted1 := 10
	got1 := y.GetLeapMonth()
	if excepted1 != got1 {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

// TestRabByungYear1 藏历年测试
func TestRabByungYear1(t *testing.T) {
	y, _ := tyme.RabByungYear{}.FromYear(1027)
	excepted := "第一饶迥火兔年"
	got := y.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

// TestRabByungYear2 藏历年测试
func TestRabByungYear2(t *testing.T) {
	y, _ := tyme.RabByungYear{}.FromYear(2010)
	excepted := "第十七饶迥铁虎年"
	got := y.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

// TestRabByungYear3 藏历年测试
func TestRabByungYear3(t *testing.T) {
	y, _ := tyme.RabByungYear{}.FromYear(2043)
	excepted := 5
	got := y.GetLeapMonth()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	y, _ = tyme.RabByungYear{}.FromYear(2044)
	excepted = 0
	got = y.GetLeapMonth()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

// TestRabByungYear4 藏历年测试
func TestRabByungYear4(t *testing.T) {
	y, _ := tyme.RabByungYear{}.FromYear(1961)
	excepted := "第十六饶迥铁牛年"
	got := y.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
