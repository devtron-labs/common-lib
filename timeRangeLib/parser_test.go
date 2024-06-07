/*
 * Copyright (c) 2024. Devtron Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package timeRangeLib

import (
	"testing"
	"time"
)

func TestGetScheduleSpec_FixedFrequency(t *testing.T) {
	// Define test cases
	testCases := []struct {
		description        string
		timeRange          TimeRange
		targetTime         time.Time
		expectedWindowEdge time.Time
		expectedIsBetween  bool
	}{
		{

			description: "Target time outside the time range",
			timeRange: TimeRange{
				TimeFrom:  time.Date(2024, time.February, 26, 8, 0, 0, 0, time.Local),
				TimeTo:    time.Date(2024, time.February, 26, 17, 0, 0, 0, time.Local),
				Frequency: Fixed,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 18, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time outside the time range",
			timeRange: TimeRange{
				TimeFrom:  time.Date(2024, time.February, 26, 8, 0, 0, 0, time.Local),
				TimeTo:    time.Date(2024, time.February, 26, 7, 0, 0, 0, time.Local),
				Frequency: Fixed,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time before  the time range",
			timeRange: TimeRange{
				TimeFrom:  time.Date(2024, time.February, 26, 8, 0, 0, 0, time.Local),
				TimeTo:    time.Date(2024, time.February, 26, 11, 0, 0, 0, time.Local),
				Frequency: Fixed,
			},
			targetTime:         time.Date(2024, time.February, 26, 7, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.February, 26, 8, 0, 0, 0, time.Local),
			expectedIsBetween:  false,
		},
		{
			description: "Target time outside the time range and TimeFrom is equal to TimeTo",
			timeRange: TimeRange{
				TimeFrom:  time.Date(2024, time.February, 26, 8, 0, 0, 0, time.Local),
				TimeTo:    time.Date(2024, time.February, 26, 8, 0, 0, 0, time.Local),
				Frequency: Fixed,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 07, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time ",
			timeRange: TimeRange{
				TimeFrom:  time.Date(2024, time.February, 26, 8, 0, 0, 0, time.Local),
				TimeTo:    time.Date(2024, time.February, 26, 11, 0, 0, 0, time.Local),
				Frequency: Fixed,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 9, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.February, 26, 11, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time within the time range",
			timeRange: TimeRange{
				HourMinuteFrom: "08:00",
				HourMinuteTo:   "17:00",
				Frequency:      Daily,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 26, 17, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time inside the time range for same start and end",
			timeRange: TimeRange{
				HourMinuteFrom: "08:00",
				HourMinuteTo:   "08:00",
				Frequency:      Daily,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 18, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 27, 8, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time inside the time range for same start and end",
			timeRange: TimeRange{
				HourMinuteFrom: "08:70",
				HourMinuteTo:   "08:00",
				Frequency:      Daily,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 18, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time inside the time range for same start and end",
			timeRange: TimeRange{
				HourMinuteFrom: "12:00ab",
				HourMinuteTo:   "14:00ab",
				Frequency:      Daily,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 18, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time inside the time range for same start and end",
			timeRange: TimeRange{
				HourMinuteFrom: "08:59",
				HourMinuteTo:   "24:00",
				Frequency:      Daily,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 18, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time outside the time range",
			timeRange: TimeRange{
				HourMinuteFrom: "08:00",
				HourMinuteTo:   "07:00",
				Frequency:      Daily,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 18, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 27, 07, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time inside the time range and HourMinuteFrom<HourMinuteTo,expectedWindowEdge=HourMinuteFrom",
			timeRange: TimeRange{
				HourMinuteFrom: "08:00",
				HourMinuteTo:   "09:00",
				Weekdays:       []time.Weekday{time.Monday, time.Wednesday, time.Friday},
				Frequency:      Weekly,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 8, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 26, 9, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time inside the time range and HourMinuteFrom<HourMinuteTo,expectedWindowEdge=HourMinuteTo",
			timeRange: TimeRange{
				HourMinuteFrom: "08:00",
				HourMinuteTo:   "09:00",
				Weekdays:       []time.Weekday{time.Monday, time.Wednesday, time.Friday},
				Frequency:      Weekly,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 9, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 28, 8, 0, 0, 0, time.Local),
			expectedIsBetween:  false,
		},
		{
			description: "Target time before  the time range  and HourMinuteFrom<HourMinuteTo ",
			timeRange: TimeRange{
				HourMinuteFrom: "08:00",
				HourMinuteTo:   "09:00",
				Weekdays:       []time.Weekday{time.Monday, time.Wednesday, time.Friday},
				Frequency:      Weekly,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 7, 30, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 26, 8, 0, 0, 0, time.Local),
			expectedIsBetween:  false,
		},
		{
			description: "Target time after the time range and HourMinuteFrom<HourMinuteTo ",
			timeRange: TimeRange{
				HourMinuteFrom: "08:00",
				HourMinuteTo:   "09:00",
				Weekdays:       []time.Weekday{time.Monday, time.Wednesday, time.Friday},
				Frequency:      Weekly,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 9, 30, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 28, 8, 0, 0, 0, time.Local),
			expectedIsBetween:  false,
		},
		{
			description: "Target time inside  the time range HourMinuteFrom>HourMinuteTo",
			timeRange: TimeRange{
				HourMinuteFrom: "08:00",
				HourMinuteTo:   "07:00",
				Weekdays:       []time.Weekday{time.Monday, time.Wednesday, time.Friday},
				Frequency:      Weekly,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 8, 30, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 27, 7, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time before  the time range HourMinuteFrom>HourMinuteTo",
			timeRange: TimeRange{
				HourMinuteFrom: "08:00",
				HourMinuteTo:   "07:00",
				Weekdays:       []time.Weekday{time.Monday, time.Wednesday, time.Friday},
				Frequency:      Weekly,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 7, 30, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 26, 8, 0, 0, 0, time.Local),
			expectedIsBetween:  false,
		},
		{
			description: "Target time after  the time range HourMinuteFrom>HourMinuteTo",
			timeRange: TimeRange{
				HourMinuteFrom: "08:00",
				HourMinuteTo:   "07:00",
				Weekdays:       []time.Weekday{time.Monday, time.Wednesday, time.Friday},
				Frequency:      Weekly,
			},
			targetTime:         time.Date(2024, time.Month(2), 27, 7, 30, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 28, 8, 0, 0, 0, time.Local),
			expectedIsBetween:  false,
		},
		{
			description: "Target time inside  the time range HourMinuteFrom==HourMinuteTo",
			timeRange: TimeRange{
				HourMinuteFrom: "08:00",
				HourMinuteTo:   "08:00",
				Weekdays:       []time.Weekday{time.Monday, time.Wednesday, time.Friday},
				Frequency:      Weekly,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 8, 30, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 27, 8, 00, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time inside  the time range HourMinuteFrom<HourMinuteTo",
			timeRange: TimeRange{
				HourMinuteFrom: "08:00",
				HourMinuteTo:   "10:00",
				WeekdayFrom:    1,
				WeekdayTo:      3,
				Frequency:      WeeklyRange,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 8, 30, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 28, 10, 00, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time before  the time range HourMinuteFrom<HourMinuteTo",
			timeRange: TimeRange{
				HourMinuteFrom: "08:00",
				HourMinuteTo:   "10:00",
				WeekdayFrom:    1,
				WeekdayTo:      3,
				Frequency:      WeeklyRange,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 7, 30, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 26, 8, 00, 0, 0, time.Local),
			expectedIsBetween:  false,
		},
		{
			description: "Target time after  the time range HourMinuteFrom<HourMinuteTo",
			timeRange: TimeRange{
				HourMinuteFrom: "08:00",
				HourMinuteTo:   "10:00",
				WeekdayFrom:    1,
				WeekdayTo:      3,
				Frequency:      WeeklyRange,
			},
			targetTime:         time.Date(2024, time.Month(2), 28, 10, 30, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(3), 4, 8, 00, 0, 0, time.Local),
			expectedIsBetween:  false,
		},
		{
			description: "Target time inside  the time range HourMinuteFrom>HourMinuteTo",
			timeRange: TimeRange{
				HourMinuteFrom: "08:00",
				HourMinuteTo:   "07:00",
				WeekdayFrom:    1,
				WeekdayTo:      3,
				Frequency:      WeeklyRange,
			},
			targetTime:         time.Date(2024, time.Month(2), 28, 10, 30, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(3), 4, 8, 00, 0, 0, time.Local),
			expectedIsBetween:  false,
		},
		{
			description: "Target time before  the time range HourMinuteFrom>HourMinuteTo",
			timeRange: TimeRange{
				HourMinuteFrom: "08:00",
				HourMinuteTo:   "07:00",
				WeekdayFrom:    1,
				WeekdayTo:      3,
				Frequency:      WeeklyRange,
			},
			targetTime:         time.Date(2024, time.Month(2), 28, 7, 30, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(3), 4, 8, 00, 0, 0, time.Local),
			expectedIsBetween:  false,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo ",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "08:00",
				DayFrom:        26,
				DayTo:          4,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(4), 27, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(5), 4, 8, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo  for dec ",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "08:00",
				DayFrom:        26,
				DayTo:          4,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(12), 27, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2025, time.Month(1), 4, 8, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo  for dec ",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "08:00",
				DayFrom:        26,
				DayTo:          4,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(12), 4, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(12), 26, 9, 0, 0, 0, time.Local),
			expectedIsBetween:  false,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo  for dec ",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "08:00",
				DayFrom:        26,
				DayTo:          4,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2025, time.Month(1), 4, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2025, time.Month(1), 26, 9, 0, 0, 0, time.Local),
			expectedIsBetween:  false,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo  for dec and target time is for next month ",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "08:00",
				DayFrom:        26,
				DayTo:          4,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2025, time.Month(1), 1, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2025, time.Month(1), 4, 8, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo  for dec and target time is for next month ",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "08:00",
				DayFrom:        -2,
				DayTo:          -1,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2025, time.Month(1), 30, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2025, time.Month(1), 31, 8, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo  for dec and target time is for next month ",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "08:00",
				DayFrom:        -1,
				DayTo:          -3,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2025, time.Month(1), 30, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo  for dec and target time is for next month ",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "10:00",
				DayFrom:        -3,
				DayTo:          4,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(3), 4, 8, 59, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(3), 4, 10, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo  for dec and target time is for next month ",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "10:00",
				DayFrom:        -3,
				DayTo:          4,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(3), 4, 9, 59, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(3), 4, 10, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo  for dec and target time is for next month ",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "10:00",
				DayFrom:        -3,
				DayTo:          4,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(3), 4, 10, 1, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(3), 29, 9, 0, 0, 0, time.Local),
			expectedIsBetween:  false,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo  for dec and target time is for next month ",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "10:00",
				DayFrom:        -3,
				DayTo:          4,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(2), 1, 9, 30, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 4, 10, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo  for dec and target time is for next month ",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "10:00",
				DayFrom:        -3,
				DayTo:          4,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(2), 4, 10, 1, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 27, 9, 0, 0, 0, time.Local),
			expectedIsBetween:  false,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo  for dec and target time is for next month ",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "10:00",
				DayFrom:        -3,
				DayTo:          4,
				Frequency:      Monthly,
			},
			//29,30,31,1,2,3,4
			targetTime:         time.Date(2024, time.Month(2), 4, 9, 59, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 4, 10, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo  for dec and target time is for next month ",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "10:00",
				DayFrom:        -3,
				DayTo:          4,
				Frequency:      Monthly,
			},
			//29,30,31,1,2,3,4
			targetTime:         time.Date(2024, time.Month(2), 4, 9, 01, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 4, 10, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo  for dec and target time is for next month ",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "10:00",
				DayFrom:        -3,
				DayTo:          4,
				Frequency:      Monthly,
			},
			//29,30,31,1,2,3,4
			targetTime:         time.Date(2024, time.Month(2), 3, 9, 01, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 4, 10, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo  for dec and target time is for next month ",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "10:00",
				DayFrom:        -3,
				DayTo:          4,
				Frequency:      Monthly,
			},
			//29,30,31,1,2,3,4
			targetTime:         time.Date(2024, time.Month(2), 4, 6, 01, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 4, 10, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},

		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "08:00",
				DayFrom:        26,
				DayTo:          -3,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 8, 59, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(2), 26, 9, 0, 0, 0, time.Local),
			expectedIsBetween:  false,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo and dayTo< -4",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "08:00",
				DayFrom:        26,
				DayTo:          -4,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(5), 26, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo and dayTo< -4",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "11:00",
				DayFrom:        26,
				DayTo:          -4,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(5), 26, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "08:00",
				DayFrom:        26,
				DayTo:          -2,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(5), 26, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(5), 30, 8, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "08:00",
				DayFrom:        26,
				DayTo:          -1,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(5), 26, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(5), 31, 8, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo , both dayFrom and DayTo less than 0",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "08:00",
				DayFrom:        -3,
				DayTo:          -1,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(5), 29, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(5), 31, 8, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo , both dayFrom and DayTo less than 0",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "08:00",
				DayFrom:        -3,
				DayTo:          -4,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(5), 29, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom ==DayTo , both dayFrom and DayTo less than 0",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "08:00",
				DayFrom:        -3,
				DayTo:          -3,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(5), 29, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom <DayTo , both dayFrom and DayTo less than 0",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "08:00",
				DayFrom:        -3,
				DayTo:          -2,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(5), 29, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(5), 30, 8, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom <DayTo , both dayFrom and DayTo less than 0",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "08:00",
				DayFrom:        -3,
				DayTo:          -1,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(5), 29, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(5), 31, 8, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom <DayTo ,  dayFrom less than 0 and dayTo>0",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "08:00",
				DayFrom:        -3,
				DayTo:          24,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(6), 20, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(6), 24, 8, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom <DayTo ,  dayFrom less than 0 and dayTo>0",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00ab",
				HourMinuteTo:   "08:00",
				DayFrom:        -3,
				DayTo:          24,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(6), 20, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom <DayTo ,  dayFrom less than 0 and dayTo>0",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "08:00ab",
				DayFrom:        -3,
				DayTo:          24,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(6), 20, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom <DayTo ,  dayFrom less than 0 and dayTo>0",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "08:00ab",
				Frequency:      Daily,
			},
			targetTime:         time.Date(2024, time.Month(6), 20, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom <DayTo ,  dayFrom less than 0 and dayTo>0",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "08:00ab",
				DayFrom:        -3,
				DayTo:          24,
				Weekdays:       []time.Weekday{time.Monday, time.Wednesday, time.Friday},
				Frequency:      Weekly,
			},
			targetTime:         time.Date(2024, time.Month(6), 20, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},

		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom <DayTo",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "08:00",
				DayFrom:        26,
				DayTo:          27,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(5), 26, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(5), 27, 8, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom <DayTo",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00:00",
				HourMinuteTo:   "08:00",
				DayFrom:        26,
				DayTo:          27,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(5), 26, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom <DayTo,expectedWindowEdge=HourMinuteFrom",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "08:00",
				DayFrom:        26,
				DayTo:          27,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(5), 26, 9, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(5), 27, 8, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom <DayTo,expectedWindowEdge=HourMinuteFrom",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "08:00",
				DayFrom:        26,
				DayTo:          27,
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(5), 26, 8, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(5), 26, 9, 0, 0, 0, time.Local),
			expectedIsBetween:  false,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom <DayTo",
			timeRange: TimeRange{
				TimeFrom: time.Time{},
				TimeTo:   time.Date(2024, time.Month(5), 26, 10, 0, 0, 0, time.Local),

				Frequency: Fixed,
			},
			targetTime:         time.Date(2024, time.Month(5), 26, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom <DayTo",
			timeRange: TimeRange{
				TimeFrom: time.Date(2024, time.Month(5), 26, 10, 0, 0, 0, time.Local),
				TimeTo:   time.Time{},

				Frequency: Fixed,
			},
			targetTime:         time.Date(2024, time.Month(5), 26, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom <DayTo",
			timeRange: TimeRange{
				TimeFrom: time.Date(2024, time.Month(5), 26, 10, 0, 0, 0, time.Local),
				TimeTo:   time.Date(2024, time.Month(5), 26, 10, 0, 0, 0, time.Local),

				Frequency: Fixed,
			},
			targetTime:         time.Date(2024, time.Month(5), 26, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time within the time range",
			timeRange: TimeRange{
				HourMinuteFrom: "",
				HourMinuteTo:   "17:00",
				Frequency:      Daily,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time within the time range",
			timeRange: TimeRange{
				HourMinuteFrom: "17:01",
				HourMinuteTo:   "17:00",
				Weekdays:       []time.Weekday{},
				Frequency:      Weekly,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time within the time range",
			timeRange: TimeRange{
				HourMinuteFrom: "15:00",
				HourMinuteTo:   "17:00",
				WeekdayFrom:    1,
				Frequency:      WeeklyRange,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time within the time range",
			timeRange: TimeRange{
				HourMinuteFrom: "",
				HourMinuteTo:   "17:00",
				Frequency:      Monthly,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time within the time range",
			timeRange: TimeRange{
				HourMinuteFrom: "12:00",
				HourMinuteTo:   "17:00",
				Frequency:      Monthly,
				DayFrom:        0,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time before the time range for edge case february goes to next month",
			timeRange: TimeRange{
				HourMinuteFrom: "12:00",
				HourMinuteTo:   "17:00",
				Frequency:      Monthly,
				DayFrom:        1,
				DayTo:          1,
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(3), 1, 12, 0, 0, 0, time.Local),
			expectedIsBetween:  false,
		},
		{
			description: "Target time within the time range",
			timeRange: TimeRange{
				HourMinuteFrom: "",
				HourMinuteTo:   "17:00",
				Frequency:      "",
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time within the time range",
			timeRange: TimeRange{
				HourMinuteFrom: "17:05",
				HourMinuteTo:   "",
				Frequency:      "",
			},
			targetTime:         time.Date(2024, time.Month(2), 26, 10, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "17:00",
				DayFrom:        26,
				DayTo:          29,
				Frequency:      Monthly,
				TimeFrom:       time.Date(2024, time.Month(5), 26, 10, 0, 0, 0, time.Local),
			},
			targetTime:         time.Date(2024, time.Month(5), 26, 11, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(5), 29, 17, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "17:00",
				DayFrom:        26,
				DayTo:          29,
				Frequency:      Monthly,
				TimeFrom:       time.Date(2024, time.Month(5), 26, 10, 0, 0, 0, time.Local),
				TimeTo:         time.Date(2024, time.Month(5), 29, 13, 0, 0, 0, time.Local),
			},
			targetTime:         time.Date(2024, time.Month(5), 26, 11, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(5), 29, 13, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "17:00",
				DayFrom:        26,
				DayTo:          29,
				Frequency:      Monthly,
				TimeTo:         time.Date(2024, time.Month(5), 29, 13, 0, 0, 0, time.Local),
			},
			targetTime:         time.Date(2024, time.Month(5), 26, 11, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Date(2024, time.Month(5), 29, 13, 0, 0, 0, time.Local),
			expectedIsBetween:  true,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "17:00",
				WeekdayFrom:    9,
				WeekdayTo:      2,
				Frequency:      WeeklyRange,
				TimeTo:         time.Date(2024, time.Month(5), 29, 13, 0, 0, 0, time.Local),
			},
			targetTime:         time.Date(2024, time.Month(5), 26, 11, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "17:00",
				WeekdayFrom:    5,
				WeekdayTo:      11,
				Frequency:      WeeklyRange,
				TimeTo:         time.Date(2024, time.Month(5), 29, 13, 0, 0, 0, time.Local),
			},
			targetTime:         time.Date(2024, time.Month(5), 26, 11, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "17:00ac",
				WeekdayFrom:    5,
				WeekdayTo:      6,
				Frequency:      WeeklyRange,
				TimeTo:         time.Date(2024, time.Month(5), 29, 13, 0, 0, 0, time.Local),
			},
			targetTime:         time.Date(2024, time.Month(5), 26, 11, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
		{
			description: "Target time between  the time range  and HourMinuteFrom>HourMinuteTo and DayFrom >DayTo",
			timeRange: TimeRange{
				HourMinuteFrom: "09:00",
				HourMinuteTo:   "17:00ac",
				WeekdayFrom:    5,
				WeekdayTo:      6,
				Frequency:      "hello",
				TimeTo:         time.Date(2024, time.Month(5), 29, 13, 0, 0, 0, time.Local),
			},
			targetTime:         time.Date(2024, time.Month(5), 26, 11, 0, 0, 0, time.Local),
			expectedWindowEdge: time.Time{},
			expectedIsBetween:  false,
		},
	}
	// Execute test cases
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			nextWindowEdge, isTimeBetween, _ := tc.timeRange.GetTimeRangeWindow(tc.targetTime)
			if nextWindowEdge != tc.expectedWindowEdge || isTimeBetween != tc.expectedIsBetween {
				t.Errorf("Test case failed: %s\nExpected nextWindowEdge: %v, got: %v\nExpected isTimeBetween: %t, got: %t", tc.description, tc.expectedWindowEdge, nextWindowEdge, tc.expectedIsBetween, isTimeBetween)
			}
		})
	}
}
