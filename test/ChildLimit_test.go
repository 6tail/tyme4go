package test

import (
	"testing"

	"github.com/6tail/tyme4go/tyme"
)

func TestChildLimitGregorianReform(t *testing.T) {
	birth, err := tyme.SolarTime{}.FromYmdHms(1572, 6, 27, 2, 49, 0)
	if err != nil {
		t.Fatal(err)
	}
	childLimit := tyme.ChildLimit{}.FromSolarTime(*birth, tyme.WOMAN)
	if got, want := childLimit.GetEndTime().String(), "1582年10月19日 13:29:00"; got != want {
		t.Errorf("end time: got %s, want %s", got, want)
	}
	if childLimit.IsForward() {
		t.Error("expected backward calculation for a woman born in a yang year")
	}
	if got, want := childLimit.GetStartDecadeFortune().GetStartAge(), 11; got != want {
		t.Errorf("start age: got %d, want %d", got, want)
	}
}

func TestChildLimitProvidersGregorianReform(t *testing.T) {
	birth, err := tyme.SolarTime{}.FromYmdHms(1572, 6, 27, 2, 49, 0)
	if err != nil {
		t.Fatal(err)
	}
	term, err := tyme.SolarTerm{}.FromName(1572, "芒种")
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name     string
		provider tyme.IChildLimitProvider
		want     string
	}{
		{"default", tyme.DefaultChildLimitProvider{}, "1582年10月19日 13:29:00"},
		{"lunar_sect_1", tyme.LunarSect1ChildLimitProvider{}, "1582年10月17日 02:49:00"},
		{"lunar_sect_2", tyme.LunarSect2ChildLimitProvider{}, "1582年10月19日 12:49:00"},
		{"china_95", tyme.China95ChildLimitProvider{}, "1582年10月19日 02:49:00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := tt.provider.GetInfo(*birth, *term)
			if got := info.GetEndTime().String(); got != tt.want {
				t.Errorf("end time: got %s, want %s", got, tt.want)
			}
		})
	}
}
