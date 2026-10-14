package tyme

import "testing"

func TestChildLimitDateArithmetic(t *testing.T) {
	tests := []struct {
		name  string
		birth [6]int
		add   [6]int
		want  string
	}{
		{"reproduced_panic", [6]int{1572, 6, 27, 2, 49, 0}, [6]int{10, 3, 12, 10, 40, 0}, "1582年10月19日 13:29:00"},
		{"silent_wrong_day", [6]int{1582, 9, 27, 0, 0, 0}, [6]int{0, 0, 18, 0, 0, 0}, "1582年10月25日 00:00:00"},
		{"second_crosses_gap", [6]int{1582, 10, 4, 23, 59, 59}, [6]int{0, 0, 0, 0, 0, 1}, "1582年10月15日 00:00:00"},
		{"day_after_gap", [6]int{1582, 10, 15, 2, 49, 0}, [6]int{0, 0, 1, 0, 0, 0}, "1582年10月16日 02:49:00"},
		{"last_day_of_reform_month", [6]int{1582, 10, 31, 23, 59, 59}, [6]int{0, 0, 0, 0, 0, 1}, "1582年11月1日 00:00:00"},
		{"year_into_reform_month", [6]int{1581, 10, 15, 2, 49, 0}, [6]int{1, 0, 0, 0, 0, 0}, "1582年10月15日 02:49:00"},
		{"year_out_of_reform_month", [6]int{1582, 10, 15, 2, 49, 0}, [6]int{1, 0, 0, 0, 0, 0}, "1583年10月15日 02:49:00"},
		{"month_into_reform_month", [6]int{1582, 9, 15, 2, 49, 0}, [6]int{0, 1, 0, 0, 0, 0}, "1582年10月15日 02:49:00"},
		{"month_out_of_reform_month", [6]int{1582, 10, 15, 2, 49, 0}, [6]int{0, 1, 0, 0, 0, 0}, "1582年11月15日 02:49:00"},
		{"year_lands_in_gap", [6]int{1581, 10, 9, 2, 49, 0}, [6]int{1, 0, 0, 0, 0, 0}, "1582年10月19日 02:49:00"},
		{"month_lands_in_gap", [6]int{1582, 9, 9, 2, 49, 0}, [6]int{0, 1, 0, 0, 0, 0}, "1582年10月19日 02:49:00"},
		{"leap_day_overflow", [6]int{2000, 2, 29, 8, 30, 0}, [6]int{1, 0, 0, 0, 0, 0}, "2001年3月1日 08:30:00"},
		{"month_end_overflow", [6]int{2001, 1, 31, 8, 30, 0}, [6]int{0, 1, 0, 0, 0, 0}, "2001年3月3日 08:30:00"},
		{"julian_leap_year", [6]int{1500, 2, 28, 0, 0, 0}, [6]int{0, 0, 1, 0, 0, 0}, "1500年2月29日 00:00:00"},
		{"gregorian_common_year", [6]int{1900, 2, 28, 0, 0, 0}, [6]int{0, 0, 1, 0, 0, 0}, "1900年3月1日 00:00:00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, a := tt.birth, tt.add
			birth, err := SolarTime{}.FromYmdHms(b[0], b[1], b[2], b[3], b[4], b[5])
			if err != nil {
				t.Fatal(err)
			}
			info := AbstractChildLimitProvider{}.next(*birth, a[0], a[1], a[2], a[3], a[4], a[5])
			if got := info.GetEndTime().String(); got != tt.want {
				t.Errorf("end time: got %s, want %s", got, tt.want)
			}
			if got := info.GetStartTime().String(); got != birth.String() {
				t.Errorf("start time: got %s, want %s", got, birth.String())
			}
			gotCounts := [5]int{info.GetYearCount(), info.GetMonthCount(), info.GetDayCount(), info.GetHourCount(), info.GetMinuteCount()}
			wantCounts := [5]int{a[0], a[1], a[2], a[3], a[4]}
			if gotCounts != wantCounts {
				t.Errorf("duration counts: got %v, want %v", gotCounts, wantCounts)
			}
		})
	}
}

func TestChildLimitZeroOffsetPreservesReformMonth(t *testing.T) {
	for day := 1; day <= 31; day++ {
		if day >= 5 && day <= 14 {
			continue
		}
		birth, err := SolarTime{}.FromYmdHms(1582, 10, day, 12, 34, 56)
		if err != nil {
			t.Fatal(err)
		}
		info := AbstractChildLimitProvider{}.next(*birth, 0, 0, 0, 0, 0, 0)
		if got := info.GetEndTime().String(); got != birth.String() {
			t.Errorf("zero offset: got %s, want %s", got, birth.String())
		}
	}
}
