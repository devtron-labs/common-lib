package scheduler

import (
	"strconv"
	"time"
)

func (tr TimeRange) isMonthOverlapping() bool {
	dayFrom := tr.DayFrom
	dayTo := tr.DayTo
	if dayFrom > 0 && dayTo > 0 && dayTo < dayFrom {
		return true
	} else if dayFrom < 0 && dayTo > 0 {
		return true
	}
	return false
}

func (tr TimeRange) isToHourMinuteBeforeWindowEnd(targetTime time.Time) bool {

	currentHourMinute, _ := time.Parse(parseFormat, targetTime.Format(parseFormat))

	parsedHourTo, _ := time.Parse(parseFormat, tr.HourMinuteTo)

	return currentHourMinute.Before(parsedHourTo)
}

func (tr TimeRange) compareHourMinute() bool {
	parseHourFrom, _ := time.Parse(parseFormat, tr.HourMinuteFrom)
	parsedHourTo, _ := time.Parse(parseFormat, tr.HourMinuteTo)
	return parsedHourTo.Before(parseHourFrom)
}

func getDaysCount(timeRange TimeRange, monthEnd int) int {

	windowEndDay := timeRange.DayTo
	if windowEndDay < 0 {
		windowEndDay = monthEnd + 1 + windowEndDay
	}

	windowStartDay := timeRange.DayFrom
	if windowStartDay < 0 {
		windowStartDay = monthEnd + 1 + windowStartDay
	}

	totalDays := windowEndDay - windowStartDay
	if timeRange.isMonthOverlapping() {
		totalDays = totalDays + monthEnd
	}
	return totalDays
}

func getLastDayOfMonth(targetYear int, targetMonth time.Month) int {
	firstDayOfNextMonth := time.Date(targetYear, targetMonth+1, 1, 0, 0, 0, 0, time.UTC)
	lastDayOfMonth := firstDayOfNextMonth.Add(-time.Hour * 24).Day()
	return lastDayOfMonth
}

func constructDateTime(hourMinute string, days int) time.Time {
	now := time.Now()
	dateTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	fromHour, _ := strconv.Atoi(getHour(hourMinute))
	fromMinute, _ := strconv.Atoi(getMinute(hourMinute))

	dateTime = dateTime.Add(time.Duration(fromHour+24*days)*time.Hour + time.Duration(fromMinute)*time.Minute)
	return dateTime
}
