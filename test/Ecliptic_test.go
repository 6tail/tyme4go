package test

import (
	"github.com/6tail/tyme4go/tyme"
	"testing"
)

// TestEcliptic0 黄道黑道十二神测试
func TestEcliptic0(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2023, 10, 30)
	star := d.GetLunarDay().GetTwelveStar()
	excepted := "天德"
	got := star.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "黄道"
	got = star.GetEcliptic().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "吉"
	got = star.GetEcliptic().GetLuck().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
