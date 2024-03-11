package timeRangeLib

import (
	"fmt"
	"strconv"
	"time"
)

func (tr TimeRange) getCron(lastDayOfMonth int) string {
	// one function, rename to parse hour minute
	minute := getMinute(tr.HourMinuteFrom)
	hour := getHour(tr.HourMinuteFrom)

	switch tr.Frequency {
	case Daily:
		return dailyCron(minute, hour)
	case Weekly:
		return weeklyCron(minute, hour, tr.Weekdays)
	case WeeklyRange:
		return weeklyRangeCron(minute, hour, toString(tr.WeekdayFrom))
	case Monthly:
		return monthlyCron(minute, hour, tr.DayFrom, lastDayOfMonth)
	}
	return ""
}

func dailyCron(minute, hour string) string {
	return fmt.Sprintf("%s %s * * *", minute, hour)
}

func weeklyCron(minute, hour string, weekdays []time.Weekday) string {
	days := weekdaysToString(weekdays)
	return fmt.Sprintf("%s %s * * %s", minute, hour, days)
}

func weeklyRangeCron(minute, hour string, weekdayFrom string) string {
	return fmt.Sprintf("%s %s * * %s", minute, hour, weekdayFrom)
}

func monthlyCron(minute, hour string, dayFrom int, lastDayOfMonth int) string {
	if dayFrom < 0 {
		// move to function which will tell why we add 1
		dayFrom = lastDayOfMonth + 1 + dayFrom
	}
	day := strconv.Itoa(dayFrom)

	return fmt.Sprintf("%s %s %s * *", minute, hour, day)
}
