package test

import (
	"github.com/6tail/tyme4go/tyme"
	"testing"
)

// TestLegalHoliday0 法定假日测试
func TestLegalHoliday0(t *testing.T) {
	d, _ := tyme.LegalHoliday{}.FromYmd(2011, 5, 1)
	excepted := "2011年5月1日 劳动节(休)"
	got := d.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "2011年5月2日 劳动节(休)"
	got = d.Next(1).String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "2011年6月4日 端午节(休)"
	got = d.Next(2).String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "2011年4月30日 劳动节(休)"
	got = d.Next(-1).String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "2011年4月5日 清明节(休)"
	got = d.Next(-2).String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestLegalHoliday1(t *testing.T) {
	d, _ := tyme.LegalHoliday{}.FromYmd(2001, 12, 29)
	excepted := "2001年12月29日 元旦节(班)"
	got := d.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
