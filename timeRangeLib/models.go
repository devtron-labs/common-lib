package timeRangeLib

import (
	"github.com/robfig/cron/v3"
	"time"
)

type TimeRange struct {
	TimeFrom       time.Time
	TimeTo         time.Time
	HourMinuteFrom string
	HourMinuteTo   string
	DayFrom        int
	DayTo          int
	WeekdayFrom    time.Weekday
	WeekdayTo      time.Weekday
	Weekdays       []time.Weekday
	Frequency      Frequency
}

// random values for  for understanding HH:MM format
const hourMinuteFormat = "15:04"

type Frequency string

const (
	Fixed       Frequency = "Fixed"
	Daily       Frequency = "Daily"
	Weekly      Frequency = "Weekly"
	WeeklyRange Frequency = "WeeklyRange"
	Monthly     Frequency = "Monthly"
)

var AllowedFrequencies = []Frequency{Fixed, Daily, Weekly, WeeklyRange, Monthly}

const CRON = cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow
