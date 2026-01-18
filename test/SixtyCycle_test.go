package test

import (
	"testing"

	"github.com/6tail/tyme4go/tyme"
)

// TestSixtyCycle0 六十甲子测试
func TestSixtyCycle0(t *testing.T) {
	excepted := "丁丑"
	got := tyme.SixtyCycle{}.FromIndex(13).GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSixtyCycle1(t *testing.T) {
	o, _ := tyme.SixtyCycle{}.FromName("辛酉")
	excepted := "石榴木"
	got := o.GetSound().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
