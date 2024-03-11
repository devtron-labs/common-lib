package scheduler

import (
	"time"
)

func (tr TimeRange) getDuration(monthEnd int) time.Duration {
	switch tr.Frequency {
	case Daily, Weekly:
		return getDurationForHourMinute(tr)
	case WeeklyRange:
		return getDurationBetweenWeekdays(tr)
	case Monthly:
		return getDurationBetweenWeekDates(tr, monthEnd)
	}
	return 0
}

func getDurationForHourMinute(timeRange TimeRange) time.Duration {

	parsedHourFrom, _ := time.Parse(parseFormat, timeRange.HourMinuteFrom)
	parsedHourTo, _ := time.Parse(parseFormat, timeRange.HourMinuteTo)
	if parsedHourTo.Before(parsedHourFrom) || parsedHourTo.Equal(parsedHourFrom) {
		parsedHourTo = parsedHourTo.AddDate(0, 0, 1)
	}
	return parsedHourTo.Sub(parsedHourFrom)
}

func getDurationBetweenWeekdays(timeRange TimeRange) time.Duration {
	days := calculateDaysBetweenWeekdays(int(timeRange.WeekdayFrom), int(timeRange.WeekdayTo))

	fromDateTime := constructDateTime(timeRange.HourMinuteFrom, 0)
	toDateTime := constructDateTime(timeRange.HourMinuteTo, days)
	return toDateTime.Sub(fromDateTime)
}

func getDurationBetweenWeekDates(timeRange TimeRange, monthEnd int) time.Duration {

	days := getDaysCount(timeRange, monthEnd)
	//if timeRange.DayFrom > 0 && timeRange.DayTo > 0 && timeRange.DayFrom < timeRange.DayTo {
	//	days = (timeRange.DayTo) - (timeRange.DayFrom)
	//}
	fromDateTime := constructDateTime(timeRange.HourMinuteFrom, 0)
	toDateTime := constructDateTime(timeRange.HourMinuteTo, days)

	return toDateTime.Sub(fromDateTime)
}
