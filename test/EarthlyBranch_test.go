package test

import (
	"github.com/6tail/tyme4go/tyme"
	"testing"
)

func TestEarthlyBranch0(t *testing.T) {
	excepted := "子"
	got := tyme.EarthBranch{}.FromIndex(0).GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

// TestEarthlyBranch1 冲
func TestEarthlyBranch1(t *testing.T) {
	d, _ := tyme.EarthBranch{}.FromName("子")
	excepted := "午"
	got := d.GetOpposite().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

// TestEarthlyBranch2 六合
func TestEarthlyBranch2(t *testing.T) {
	d, _ := tyme.EarthBranch{}.FromName("申")
	excepted := "巳"
	got := d.GetCombine().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

// TestEarthlyBranch3 六害
func TestEarthlyBranch3(t *testing.T) {
	d, _ := tyme.EarthBranch{}.FromName("巳")
	excepted := "寅"
	got := d.GetHarm().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

// TestEarthlyBranch4 合化
func TestEarthlyBranch4(t *testing.T) {
	a, _ := tyme.EarthBranch{}.FromName("卯")
	b, _ := tyme.EarthBranch{}.FromName("戌")
	excepted := "火"
	got := a.Combine(*b).GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

// TestEarthlyBranch5 卯子无法合化
func TestEarthlyBranch5(t *testing.T) {
	a, _ := tyme.EarthBranch{}.FromName("卯")
	b, _ := tyme.EarthBranch{}.FromName("子")
	got := a.Combine(*b)
	if nil != got {
		t.Errorf("excepted: %v, got: %v", nil, got)
	}
}
