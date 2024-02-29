package scheduler

import (
	"testing"
	"time"
)

func TestGetScheduleSpec_FixedFrequency(t *testing.T) {
	//impl := &ScheduleParserImpl{}

	// Define test cases
	testCases := []struct {
		description        string
		timeRange          TimeRange
		targetTime         time.Time
		expectedWindowEdge time.Time
		expectedIsBetween  bool
	}{
		{
			description: "Target time within the time range",
			timeRange: TimeRange{
				HourMinuteFrom: "08:00",
				HourMinuteTo:   "17:00",
				Frequency:      DAILY,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 26, 17, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time outside the time range",
			timeRange: TimeRange{
				HourMinuteFrom: "08:00",
				HourMinuteTo:   "17:00",
				Frequency:      DAILY,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 18, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 27, 8, 0, 0, 0, time.Local),
			expectedIsBetween:  false,
		},
		{

			description: "Target time outside the time range",
			timeRange: TimeRange{
				TimeFrom:  time.Date(2024, time.February, 26, 8, 0, 0, 0, time.Local),
				TimeTo:    time.Date(2024, time.February, 26, 17, 0, 0, 0, time.Local),
				Frequency: FIXED,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 18, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(0001, time.January, 01, 0, 0, 0, 0, time.UTC),
			expectedIsBetween:  false,
		},
		{
			description: "Target time outside the time range",
			timeRange: TimeRange{
				TimeFrom:  time.Date(2024, time.February, 26, 8, 0, 0, 0, time.Local),
				TimeTo:    time.Date(2024, time.February, 26, 17, 0, 0, 0, time.Local),
				Frequency: FIXED,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 26, 17, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time outside the time range",
			timeRange: TimeRange{

				Frequency: WEEKLY,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 26, 17, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
	}

	// Execute test cases
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			nextWindowEdge, isTimeBetween := tc.timeRange.GetScheduleSpec(tc.targetTime)
			if nextWindowEdge != tc.expectedWindowEdge || isTimeBetween != tc.expectedIsBetween {
				t.Errorf("Test case failed: %s\nExpected nextWindowEdge: %v, got: %v\nExpected isTimeBetween: %t, got: %t", tc.description, tc.expectedWindowEdge, nextWindowEdge, tc.expectedIsBetween, isTimeBetween)
			}
		})
	}
}
