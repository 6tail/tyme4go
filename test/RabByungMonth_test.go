package test

import (
	"testing"

	"github.com/6tail/tyme4go/tyme"
)

// TestRabByungMonth0 藏历月测试
func TestRabByungMonth0(t *testing.T) {
	m, _ := tyme.RabByungMonth{}.FromYm(1950, 12)
	excepted := "第十六饶迥铁虎年十二月"
	got := m.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
