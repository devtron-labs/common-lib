package scheduler

import (
	"testing"
	"time"
)

/*
cases on frequency
case 1 : FIXED : if all fields are empty other than TimeFrom and TimeTo
then here I have add validation that TimeFrom must be less than TimeTo and TimeFrom TimeTo both should be present.
case 2 : DAILY : HourMinuteFrom and HourMinuteTo must pe present
case 3 : WEEKLY :  Weekdays must be present and HourMinuteFrom and HourMinuteTo must be present
case 4 : WEEKLY_RANGE : WeekdayFrom must be present , HourMinuteFrom and HourMinuteTo must be present
case 5 : MONTHLY :  DayFrom and  DayTo must me present , HourMinuteFrom and HourMinuteTo must be present
*/
/*
cases on field validation
DayFrom and DayTo must be from 1 to 31 as i month max have 31 days
HourMinuteFrom and HourMinuteTo must HH:MM here HH must be from 0 to 23 and MM form 0 to 59
*/
//todo have to add test case for hour and min it means that duration must like 5hrs 30 min
func TestGetDurationAndGetCron(t *testing.T) {
	//Test case 1: DAILY frequency
	timeRange1 := TimeRange{
		HourMinuteFrom: "09:00",
		HourMinuteTo:   "03:00",
		Frequency:      DAILY,
	}
	wantDuration1 := 18 * time.Hour
	wantCron1 := "00 09 * * *"
	// Test case 2: DAILY frequency
	timeRange2 := TimeRange{
		HourMinuteFrom: "12:00",
		HourMinuteTo:   "14:00",
		Frequency:      DAILY,
	}
	wantDuration2 := 2 * time.Hour
	wantCron2 := "00 12 * * *"
	timeRange3 := TimeRange{
		HourMinuteFrom: "14:00",
		HourMinuteTo:   "12:00",
		Frequency:      DAILY,
	}
	wantDuration3 := 22 * time.Hour
	wantCron3 := "00 14 * * *"
	timeRange4 := TimeRange{
		HourMinuteFrom: "09:00",
		HourMinuteTo:   "12:00",
		Weekdays:       []time.Weekday{time.Monday, time.Tuesday, time.Wednesday},
		Frequency:      WEEKLY,
	}
	wantDuration4 := 3 * time.Hour
	wantCron4 := "00 09 * * 1,2,3"

	//Test case 3: WEEKLY frequency
	timeRange5 := TimeRange{
		HourMinuteFrom: "09:00",
		HourMinuteTo:   "00:00",
		Frequency:      WEEKLY,
		Weekdays:       []time.Weekday{time.Monday, time.Wednesday, time.Friday},
	}
	wantDuration5 := 15 * time.Hour
	wantCron5 := "00 09 * * 1,3,5"
	timeRange6 := TimeRange{
		HourMinuteFrom: "17:00",
		HourMinuteTo:   "19:30",
		Frequency:      WEEKLY,
		Weekdays:       []time.Weekday{time.Monday, time.Wednesday, time.Friday},
	}
	wantDuration6 := 2*time.Hour + 30*time.Minute
	wantCron6 := "00 17 * * 1,3,5"
	timeRange7 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "01:00",
		Frequency:      WEEKLY_RANGE,
		WeekdayTo:      1,
		WeekdayFrom:    4,
	}
	wantDuration7 := 86 * time.Hour
	wantCron7 := "00 17 1 * *"
	timeRange8 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "13:00",
		Frequency:      WEEKLY_RANGE,
		WeekdayTo:      4,
		WeekdayFrom:    1,
	}
	wantDuration8 := 74 * time.Hour
	wantCron8 := "00 11 * * Monday"
	timeRange9 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "12:00",
		Frequency:      WEEKLY_RANGE,
		WeekdayTo:      9,
		WeekdayFrom:    1,
	}
	wantDuration9 := 74 * time.Hour
	wantCron9 := "00 11 * * Monday"
	timeRange10 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "12:00",
		Frequency:      WEEKLY_RANGE,
		WeekdayTo:      1,
		WeekdayFrom:    9,
	}
	wantDuration10 := 0 * time.Second
	wantCron10 := "00 11 * * %!Weekday(9)"
	timeRange11 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "12:00",
		DayFrom:        1,
		DayTo:          3,
		Frequency:      MONTHLY,
	}
	wantDuration11 := 49 * time.Hour
	wantCron11 := "00 11 1 * *"
	timeRange12 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "00:00",
		DayFrom:        1,
		DayTo:          3,
		Frequency:      MONTHLY,
	}
	wantDuration12 := 37 * time.Hour
	wantCron12 := "00 11 1 * *"
	timeRange13 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "02:00",
		DayFrom:        1,
		DayTo:          3,
		Frequency:      MONTHLY,
	}
	wantDuration13 := 39 * time.Hour
	wantCron13 := "00 11 1 * *"
	timeRange14 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "02:00",
		DayFrom:        27,
		DayTo:          -2,
		Frequency:      MONTHLY,
	}
	wantDuration14 := 39 * time.Hour
	wantCron14 := "00 11 1 * *"
	timeRange15 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "02:00",
		DayFrom:        27,
		DayTo:          -4,
		Frequency:      MONTHLY,
	}
	wantDuration15 := 0 * time.Hour
	wantCron15 := "00 11 27 * *"
	timeRange16 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "12:00",
		DayFrom:        27,
		DayTo:          -4,
		Frequency:      MONTHLY,
	}
	wantDuration16 := 1 * time.Hour
	wantCron16 := "00 11 27 * *"
	timeRange17 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "11:00",
		DayFrom:        27,
		DayTo:          -4,
		Frequency:      MONTHLY,
	}
	wantDuration17 := 0 * time.Hour
	wantCron17 := "00 11 27 * *"
	timeRange18 := TimeRange{
		HourMinuteFrom: "11:00",
		HourMinuteTo:   "11:00",
		DayFrom:        27,
		DayTo:          -5,
		Frequency:      MONTHLY,
	}
	wantDuration18 := 696 * time.Hour
	wantCron18 := "00 11 27 * *"
	//todo have to write test case for all negative handling
	tests := []struct {
		timeRange    TimeRange
		wantDuration time.Duration
		wantCron     string
	}{
		{timeRange1, wantDuration1, wantCron1},
		{timeRange2, wantDuration2, wantCron2},
		{timeRange3, wantDuration3, wantCron3},
		{timeRange4, wantDuration4, wantCron4},
		{timeRange5, wantDuration5, wantCron5},
		{timeRange6, wantDuration6, wantCron6},
		{timeRange7, wantDuration7, wantCron7},
		{timeRange8, wantDuration8, wantCron8},
		{timeRange9, wantDuration9, wantCron9},
		{timeRange10, wantDuration10, wantCron10},
		{timeRange11, wantDuration11, wantCron11},
		{timeRange12, wantDuration12, wantCron12},
		{timeRange13, wantDuration13, wantCron13},
		{timeRange14, wantDuration14, wantCron14},
		{timeRange15, wantDuration15, wantCron15},
		{timeRange16, wantDuration16, wantCron16},
		{timeRange17, wantDuration17, wantCron17},
		{timeRange18, wantDuration18, wantCron18},
	}

	for i, test := range tests {
		// Test getDuration method
		gotDuration, err := test.timeRange.getDuration(4, 2024)
		if err != nil {
			t.Errorf("Test case %d: getDuration() = %v, want %v", i+1, gotDuration, test.wantDuration)
		}
		if gotDuration != test.wantDuration {
			t.Errorf("Test case %d: getDuration() = %v, want %v", i+1, gotDuration, test.wantDuration)
		}

		// Test getCron method
		gotCron := test.timeRange.getCron()
		if gotCron != test.wantCron {
			t.Errorf("Test case %d: getCron() = %v, want %v", i+1, gotCron, test.wantCron)
		}

	}

}
