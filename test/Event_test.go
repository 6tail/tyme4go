package test

import (
	"fmt"
	"testing"

	"github.com/6tail/tyme4go/tyme"
)

// initEvents 重置事件数据并初始化所有事件
func initEvents(t *testing.T) {
	// 辅助函数：忽略错误，因为测试数据应有效
	mustUpdate := func(name string, e *tyme.Event, err error) {
		if err != nil {
			t.Fatalf("创建事件 %s 失败: %v", name, err)
		}
		tyme.EventManager{}.UpdateEvent(name, e)
	}

	// 公历现代节日
	e, err := tyme.Event{}.Builder().SolarDay(1, 1, 0).StartYear(1950).Build()
	mustUpdate("公历现代节日:元旦", e, err)
	e, err = tyme.Event{}.Builder().SolarDay(3, 8, 0).StartYear(1950).Build()
	mustUpdate("公历现代节日:妇女节", e, err)
	e, err = tyme.Event{}.Builder().SolarDay(3, 12, 0).StartYear(1979).Build()
	mustUpdate("公历现代节日:植树节", e, err)
	e, err = tyme.Event{}.Builder().SolarDay(5, 1, 0).StartYear(1950).Build()
	mustUpdate("公历现代节日:劳动节", e, err)
	e, err = tyme.Event{}.Builder().SolarDay(5, 4, 0).StartYear(1950).Build()
	mustUpdate("公历现代节日:青年节", e, err)
	e, err = tyme.Event{}.Builder().SolarDay(6, 1, 0).StartYear(1950).Build()
	mustUpdate("公历现代节日:儿童节", e, err)
	e, err = tyme.Event{}.Builder().SolarDay(7, 1, 0).StartYear(1941).Build()
	mustUpdate("公历现代节日:建党节", e, err)
	e, err = tyme.Event{}.Builder().SolarDay(8, 1, 0).StartYear(1933).Build()
	mustUpdate("公历现代节日:建军节", e, err)
	e, err = tyme.Event{}.Builder().SolarDay(9, 10, 0).StartYear(1985).Build()
	mustUpdate("公历现代节日:教师节", e, err)
	e, err = tyme.Event{}.Builder().SolarDay(10, 1, 0).StartYear(1950).Build()
	mustUpdate("公历现代节日:国庆节", e, err)

	// 农历传统节日
	e, err = tyme.Event{}.Builder().LunarDay(1, 1, 0).Build()
	mustUpdate("农历传统节日:春节", e, err)
	e, err = tyme.Event{}.Builder().LunarDay(1, 15, 0).Build()
	mustUpdate("农历传统节日:元宵节", e, err)
	e, err = tyme.Event{}.Builder().LunarDay(2, 2, 0).Build()
	mustUpdate("农历传统节日:龙头节", e, err)
	e, err = tyme.Event{}.Builder().LunarDay(3, 3, 0).Build()
	mustUpdate("农历传统节日:上巳节", e, err)
	e, err = tyme.Event{}.Builder().TermDay(7, 0).Build()
	mustUpdate("农历传统节日:清明节", e, err)
	e, err = tyme.Event{}.Builder().LunarDay(5, 5, 0).Build()
	mustUpdate("农历传统节日:端午节", e, err)
	e, err = tyme.Event{}.Builder().LunarDay(7, 7, 0).Build()
	mustUpdate("农历传统节日:七夕节", e, err)
	e, err = tyme.Event{}.Builder().LunarDay(7, 15, 0).Build()
	mustUpdate("农历传统节日:中元节", e, err)
	e, err = tyme.Event{}.Builder().LunarDay(8, 15, 0).Build()
	mustUpdate("农历传统节日:中秋节", e, err)
	e, err = tyme.Event{}.Builder().LunarDay(9, 9, 0).Build()
	mustUpdate("农历传统节日:重阳节", e, err)
	e, err = tyme.Event{}.Builder().TermDay(24, 0).Build()
	mustUpdate("农历传统节日:冬至节", e, err)
	e, err = tyme.Event{}.Builder().LunarDay(12, 8, 0).Build()
	mustUpdate("农历传统节日:腊八节", e, err)
	e, err = tyme.Event{}.Builder().LunarDay(13, 1, 0).Offset(-1).Build()
	mustUpdate("农历传统节日:除夕", e, err)

	// 其他节日
	e, err = tyme.Event{}.Builder().SolarDay(2, 14, 0).StartYear(270).Build()
	mustUpdate("情人节", e, err)
	e, err = tyme.Event{}.Builder().SolarDay(3, 15, 0).StartYear(1983).Build()
	mustUpdate("国际消费者权益日", e, err)
	e, err = tyme.Event{}.Builder().SolarDay(4, 1, 0).StartYear(1564).Build()
	mustUpdate("愚人节", e, err)
	e, err = tyme.Event{}.Builder().SolarDay(10, 31, 0).StartYear(600).Build()
	mustUpdate("万圣夜", e, err)
	e, err = tyme.Event{}.Builder().SolarDay(11, 1, 0).StartYear(600).Build()
	mustUpdate("万圣节", e, err)
	e, err = tyme.Event{}.Builder().SolarDay(12, 24, 0).StartYear(336).Build()
	mustUpdate("平安夜", e, err)
	e, err = tyme.Event{}.Builder().SolarDay(12, 25, 0).StartYear(336).Build()
	mustUpdate("圣诞节", e, err)

	e, err = tyme.Event{}.Builder().SolarWeek(3, -1, 1).StartYear(1996).Build()
	mustUpdate("全国中小学生安全教育日", e, err)
	e, err = tyme.Event{}.Builder().SolarWeek(5, 2, 0).StartYear(1914).Build()
	mustUpdate("母亲节", e, err)
	e, err = tyme.Event{}.Builder().SolarWeek(6, 3, 0).StartYear(1972).Build()
	mustUpdate("父亲节", e, err)
	e, err = tyme.Event{}.Builder().SolarWeek(11, 4, 4).StartYear(1941).Build()
	mustUpdate("感恩节", e, err)

	// 特殊节日
	e, err = tyme.Event{}.Builder().TermDay(7, -1).Build()
	mustUpdate("寒食节", e, err)
	e, err = tyme.Event{}.Builder().TermHeavenStem(3, 4, 30).Offset(10).Build()
	mustUpdate("春社", e, err)
	e, err = tyme.Event{}.Builder().TermHeavenStem(15, 4, 30).Offset(10).Build()
	mustUpdate("秋社", e, err)

	e, err = tyme.Event{}.Builder().TermHeavenStem(12, 6, 20).Build()
	mustUpdate("入伏", e, err)
	e, err = tyme.Event{}.Builder().TermHeavenStem(12, 6, 30).Build()
	mustUpdate("中伏", e, err)
	e, err = tyme.Event{}.Builder().TermHeavenStem(15, 6, 0).Build()
	mustUpdate("末伏", e, err)

	e, err = tyme.Event{}.Builder().TermHeavenStem(11, 2, 0).Build()
	mustUpdate("入梅", e, err)
	e, err = tyme.Event{}.Builder().TermEarthBranch(13, 7, 0).Build()
	mustUpdate("出梅", e, err)

	// 生日
	e, err = tyme.Event{}.Builder().SolarDay(2, 29, -1).StartYear(2004).Build()
	mustUpdate("公历生日", e, err)
	e, err = tyme.Event{}.Builder().LunarDay(4, 21, 0).StartYear(1986).Build()
	mustUpdate("农历生日", e, err)
}

func Test0(t *testing.T) {
	initEvents(t)

	e := tyme.Event{}.FromName("公历生日")
	if e == nil {
		t.Fatal("公历生日事件未找到")
	}

	d := e.GetSolarDay(2008)
	if d == nil {
		t.Fatal("2008年公历生日不应为nil")
	}
	if d.String() != "2008年2月29日" {
		t.Errorf("期望 2008年2月29日, 得到 %s", d.String())
	}

	// 2005年没有2月29，按设置应倒推1天
	d = e.GetSolarDay(2005)
	if d == nil {
		t.Fatal("2005年公历生日不应为nil")
	}
	if d.String() != "2005年2月28日" {
		t.Errorf("期望 2005年2月28日, 得到 %s", d.String())
	}

	e = tyme.Event{}.FromName("农历生日")
	if e == nil {
		t.Fatal("农历生日事件未找到")
	}
	d = e.GetSolarDay(2026)
	if d == nil {
		t.Fatal("2026年农历生日不应为nil")
	}
	if d.String() != "2026年6月6日" {
		t.Errorf("期望 2026年6月6日, 得到 %s", d.String())
	}
}

func Test1(t *testing.T) {
	initEvents(t)

	e := tyme.Event{}.FromName("公历生日")
	if e == nil {
		t.Fatal("公历生日事件未找到")
	}

	d := e.GetSolarDay(1985)
	if d != nil {
		t.Errorf("1985年公历生日应为nil, 得到 %s", d.String())
	}
}

func Test2(t *testing.T) {
	initEvents(t)

	e := tyme.Event{}.FromName("国际消费者权益日")
	if e == nil {
		t.Fatal("国际消费者权益日事件未找到")
	}
	if e.GetName() != "国际消费者权益日" {
		t.Errorf("名称期望 国际消费者权益日, 得到 %s", e.GetName())
	}

	d := e.GetSolarDay(2026)
	if d == nil {
		t.Fatal("2026年国际消费者权益日不应为nil")
	}
	if d.String() != "2026年3月15日" {
		t.Errorf("期望 2026年3月15日, 得到 %s", d.String())
	}

	events := tyme.Event{}.FromSolarDay(*d)
	// 构建期望的字符串表示，如 "[国际消费者权益日]"
	got := "["
	for i, ev := range events {
		if i > 0 {
			got += ", "
		}
		got += ev.GetName()
	}
	got += "]"
	if got != "[国际消费者权益日]" {
		t.Errorf("事件列表期望 [国际消费者权益日], 得到 %s", got)
	}

	// 打印2026年的所有事件
	fmt.Println("2026年的所有事件:")
	all := tyme.Event{}.All()
	for _, ev := range all {
		sd := ev.GetSolarDay(2026)
		if sd != nil {
			fmt.Printf("%s = %s\n", ev.GetName(), sd.String())
		} else {
			fmt.Printf("%s = 无对应日期\n", ev.GetName())
		}
	}
}

func Test3(t *testing.T) {
	initEvents(t)

	e := tyme.Event{}.FromName("全国中小学生安全教育日")
	if e == nil {
		t.Fatal("全国中小学生安全教育日事件未找到")
	}
	if e.GetName() != "全国中小学生安全教育日" {
		t.Errorf("名称期望 全国中小学生安全教育日, 得到 %s", e.GetName())
	}

	expected := map[int]string{
		2008: "2008年3月31日",
		2015: "2015年3月30日",
		2018: "2018年3月26日",
		2019: "2019年3月25日",
		2021: "2021年3月29日",
		2023: "2023年3月27日",
		2024: "2024年3月25日",
		2025: "2025年3月31日",
	}

	for year, want := range expected {
		d := e.GetSolarDay(year)
		if d == nil {
			t.Errorf("%d年事件应为非nil", year)
			continue
		}
		if d.String() != want {
			t.Errorf("%d年: 期望 %s, 得到 %s", year, want, d.String())
		}
	}
}
