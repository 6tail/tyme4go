package test

import (
	"testing"

	"github.com/6tail/tyme4go/tyme"
)

// TestElement0 五行测试
func TestElement0(t *testing.T) {
	// 地支寅的五行为木
	d, _ := tyme.EarthBranch{}.FromName("寅")
	excepted := "木"
	got := d.GetElement().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	// 地支寅的五行(木)生火
	excepted = "火"
	got = d.GetElement().GetReinforce().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
