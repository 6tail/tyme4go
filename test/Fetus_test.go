package test

import (
	"testing"

	"github.com/6tail/tyme4go/tyme"
)

// TestFetus0 胎神测试
func TestFetus0(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2021, 11, 13)
	excepted := "碓磨厕 外东南"
	got := d.GetLunarDay().GetFetusDay().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
