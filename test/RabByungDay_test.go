package test

import (
	"testing"

	"github.com/6tail/tyme4go/tyme"
)

// TestRabByungDay0 藏历日测试
func TestRabByungDay0(t *testing.T) {
	solar, _ := tyme.SolarDay{}.FromYmd(1951, 1, 8)
	d, _ := solar.GetRabByungDay()
	excepted := "第十六饶迥铁虎年十二月初一"
	got := d.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	e, _ := tyme.RabByungElement{}.FromName("铁")
	z, _ := tyme.Zodiac{}.FromName("虎")
	d, _ = tyme.RabByungDay{}.FromElementZodiac(15, *e, *z, 12, 1)
	excepted = "1951年1月8日"
	got = d.GetSolarDay().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

// TestRabByungDay1 藏历日测试
func TestRabByungDay1(t *testing.T) {
	solar, _ := tyme.SolarDay{}.FromYmd(2051, 2, 11)
	d, err := solar.GetRabByungDay()
	if err != nil {
		t.Errorf("err: %s", err)
	}
	excepted := "第十八饶迥铁马年十二月三十"
	got := d.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	e, _ := tyme.RabByungElement{}.FromName("铁")
	z, _ := tyme.Zodiac{}.FromName("马")
	d, _ = tyme.RabByungDay{}.FromElementZodiac(17, *e, *z, 12, 30)
	excepted = "2051年2月11日"
	got = d.GetSolarDay().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

// TestRabByungDay2 藏历日测试
func TestRabByungDay2(t *testing.T) {
	solar, _ := tyme.SolarDay{}.FromYmd(2025, 4, 23)
	d, _ := solar.GetRabByungDay()
	excepted := "第十七饶迥木蛇年二月廿五"
	got := d.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	e, _ := tyme.RabByungElement{}.FromName("木")
	z, _ := tyme.Zodiac{}.FromName("蛇")
	d, _ = tyme.RabByungDay{}.FromElementZodiac(16, *e, *z, 2, 25)
	excepted = "2025年4月23日"
	got = d.GetSolarDay().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

// TestRabByungDay3 藏历日测试
func TestRabByungDay3(t *testing.T) {
	solar, _ := tyme.SolarDay{}.FromYmd(1951, 2, 8)
	d, _ := solar.GetRabByungDay()
	excepted := "第十六饶迥铁兔年正月初二"
	got := d.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	e, _ := tyme.RabByungElement{}.FromName("铁")
	z, _ := tyme.Zodiac{}.FromName("兔")
	d, _ = tyme.RabByungDay{}.FromElementZodiac(15, *e, *z, 1, 2)
	excepted = "1951年2月8日"
	got = d.GetSolarDay().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

// TestRabByungDay4 藏历日测试
func TestRabByungDay4(t *testing.T) {
	solar, _ := tyme.SolarDay{}.FromYmd(1951, 1, 24)
	d, _ := solar.GetRabByungDay()
	excepted := "第十六饶迥铁虎年十二月闰十六"
	got := d.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	e, _ := tyme.RabByungElement{}.FromName("铁")
	z, _ := tyme.Zodiac{}.FromName("虎")
	d, _ = tyme.RabByungDay{}.FromElementZodiac(15, *e, *z, 12, -16)
	excepted = "1951年1月24日"
	got = d.GetSolarDay().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

// TestRabByungDay5 藏历日测试
func TestRabByungDay5(t *testing.T) {
	solar, _ := tyme.SolarDay{}.FromYmd(1961, 6, 24)
	d, _ := solar.GetRabByungDay()
	excepted := "第十六饶迥铁牛年五月十一"
	got := d.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	e, _ := tyme.RabByungElement{}.FromName("铁")
	z, _ := tyme.Zodiac{}.FromName("牛")
	d, _ = tyme.RabByungDay{}.FromElementZodiac(15, *e, *z, 5, 11)
	excepted = "1961年6月24日"
	got = d.GetSolarDay().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

// TestRabByungDay6 藏历日测试
func TestRabByungDay6(t *testing.T) {
	solar, _ := tyme.SolarDay{}.FromYmd(1952, 2, 23)
	d, _ := solar.GetRabByungDay()
	excepted := "第十六饶迥铁兔年十二月廿八"
	got := d.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	e, _ := tyme.RabByungElement{}.FromName("铁")
	z, _ := tyme.Zodiac{}.FromName("兔")
	d, _ = tyme.RabByungDay{}.FromElementZodiac(15, *e, *z, 12, 28)
	excepted = "1952年2月23日"
	got = d.GetSolarDay().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

// TestRabByungDay7 藏历日测试
func TestRabByungDay7(t *testing.T) {
	solar, _ := tyme.SolarDay{}.FromYmd(2025, 4, 26)
	d, _ := solar.GetRabByungDay()
	excepted := "第十七饶迥木蛇年二月廿九"
	got := d.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

// TestRabByungDay8 藏历日测试
func TestRabByungDay8(t *testing.T) {
	solar, _ := tyme.SolarDay{}.FromYmd(2025, 4, 25)
	d, _ := solar.GetRabByungDay()
	excepted := "第十七饶迥木蛇年二月廿七"
	got := d.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
