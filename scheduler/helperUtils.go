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

func getScheduleForFixedTime(targetTime time.Time, timeRange TimeRange) (time.Time, bool) {
	var windowStartOrEnd time.Time
	if targetTime.After(timeRange.TimeTo) {
		return windowStartOrEnd, false
	} else if targetTime.Before(timeRange.TimeFrom) {
		return timeRange.TimeFrom, false
	} else if targetTime.Before(timeRange.TimeTo) && targetTime.After(timeRange.TimeFrom) {
		return timeRange.TimeTo, true
	}
	return windowStartOrEnd, false
}

func (tr TimeRange) isToHourMinuteBeforeWindowEnd(targetTime time.Time) bool {

	currentHourMinute, _ := time.Parse(parseFormat, targetTime.Format(parseFormat))

	parsedHourTo, _ := time.Parse(parseFormat, tr.HourMinuteTo)

	return currentHourMinute.Before(parsedHourTo)
}

func getDaysCount(timeRange TimeRange, targetMonth time.Month, targetYear int) int {

	monthEnd := getLastDayOfMonth(targetYear, targetMonth)
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

//
//func getDaysCount(timeRange TimeRange, targetMonth time.Month, targetYear int) int {
//	var days int
//	var start, end time.Time
//	if timeRange.DayTo < timeRange.DayFrom {
//		if timeRange.DayFrom > 0 {
//			//27 , -2 april , 27, 28, 29
//			//27 , -5 april , 27, 28, 29 .......next month
//			timeRange.DayTo, _ = adjustDaysForMonth(timeRange.DayTo, targetMonth, targetYear)
//			start, end = getStartAndEndTime(timeRange, targetMonth)
//		}
//	} else if timeRange.DayTo > timeRange.DayFrom {
//		//-2 , -1 april 29 ,30
//		if timeRange.DayTo < 0 {
//			var lastDayOfMonth int
//			timeRange.DayFrom, lastDayOfMonth = adjustDaysForMonth(timeRange.DayFrom, targetMonth, targetYear)
//			timeRange.DayTo = lastDayOfMonth + timeRange.DayTo + 1
//			start, end = getStartAndEndTime(timeRange, targetMonth)
//		} else {
//			//-2 , 4  april 29 , 30 , 1, 2,3,4 output 5
//			timeRange.DayFrom, _ = adjustDaysForMonth(timeRange.DayFrom, targetMonth, targetYear)
//			start, end = getStartAndEndTime(timeRange, targetMonth)
//		}
//	}
//	days = int(end.Sub(start).Hours() / 24)
//	return days
//}

func getStartAndEndTime(timeRange TimeRange, targetMonth time.Month) (time.Time, time.Time) {
	start := getStartDate(timeRange, targetMonth)
	end := getEndDate(timeRange, targetMonth)
	if end.Day() < start.Day() && end.Month() == start.Month() && end.Year() == start.Year() {
		end = getEndDate(timeRange, targetMonth+1)
	}
	return start, end
}

func getEndDate(timeRange TimeRange, targetMonth time.Month) time.Time {
	return time.Date(time.Now().Year(), targetMonth, timeRange.DayTo, 0, 0, 0, 0, time.UTC)
}

func getStartDate(timeRange TimeRange, targetMonth time.Month) time.Time {
	return time.Date(time.Now().Year(), targetMonth, timeRange.DayFrom, 0, 0, 0, 0, time.UTC)
}
func adjustDaysForMonth(day int, targetMonth time.Month, targetYear int) (int, int) {
	lastDayOfMonth := getLastDayOfMonth(targetYear, targetMonth)
	if day > 0 {
		return lastDayOfMonth + day, lastDayOfMonth
	}
	return lastDayOfMonth + day + 1, lastDayOfMonth
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
