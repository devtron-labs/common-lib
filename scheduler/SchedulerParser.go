package scheduler

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"strconv"
	"strings"
	"time"
)

func (tr TimeRange) isCyclic() bool {
	dayFrom := tr.DayFrom
	dayTo := tr.DayTo
	if dayFrom > 0 && dayTo > 0 && dayTo < dayFrom {
		return true
	} else if dayFrom < 0 && dayTo > 0 {
		return true
	}
	return false
}

func (tr TimeRange) GetScheduleSpec(targetTime time.Time) (nextWindowEdge time.Time, isTimeBetween bool, err error) {
	var windowEnd time.Time
	err = tr.ValidateTimeRange()
	if err != nil {
		return nextWindowEdge, false, err
	}
	if tr.Frequency == FIXED {
		nextWindowEdge, isTimeBetween = getScheduleForFixedTime(targetTime, tr)
		return nextWindowEdge, isTimeBetween, err
	}
	month := targetTime.Month()
	year := targetTime.Year()
	day := targetTime.Day()
	cronExp := tr.getCron()
	if day >= 1 && day < tr.DayTo && tr.isCyclic() {
		if month == 1 {
			month = 12
			year = year - 1
		} else {
			month = month - 1
		}
	}
	lastDayOfMonth := getLastDayOfMonth(year, month, time.Now().Local())
	if strings.Contains(cronExp, "L-2") {
		lastDayOfMonth = lastDayOfMonth - 2
		cronExp = strings.Replace(cronExp, "L-2", intToString(lastDayOfMonth), -1)
	} else if strings.Contains(cronExp, "L-1") {
		lastDayOfMonth = lastDayOfMonth - 1
		cronExp = strings.Replace(cronExp, "L-1", intToString(lastDayOfMonth), -1)
	} else {
		cronExp = strings.Replace(cronExp, "L", intToString(lastDayOfMonth), -1)
	}
	parser := cron.NewParser(CRON)
	schedule, err := parser.Parse(cronExp)
	if err != nil {
		return nextWindowEdge, false, err
	}

	duration, err := tr.getDuration(month, year)
	if err != nil {
		return nextWindowEdge, false, err
	}
	timeMinusDuration := targetTime.Add(-1 * duration)
	windowStart := schedule.Next(timeMinusDuration)
	windowEnd = windowStart.Add(duration)
	if !tr.TimeFrom.IsZero() && windowStart.Before(tr.TimeFrom) {
		windowStart = tr.TimeFrom
	}
	if !tr.TimeTo.IsZero() && windowEnd.After(tr.TimeTo) {
		windowEnd = tr.TimeTo
	}
	if isTimeInBetween(targetTime, windowStart, windowEnd) {
		return windowEnd, true, err
	}
	return windowStart, false, err
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

func isTimeInBetween(timeCurrent, periodStart, periodEnd time.Time) bool {
	return (timeCurrent.After(periodStart) && timeCurrent.Before(periodEnd)) || timeCurrent.Equal(periodStart)
}

func getDurationForHourMinute(timeRange TimeRange) (time.Duration, error) {

	parsedHourFrom, err := time.Parse(parseFormat, timeRange.HourMinuteFrom)
	if err != nil {
		return 0, fmt.Errorf("invalid format for HourMinuteFrom: : %s", err)
	}
	parsedHourTo, err := time.Parse(parseFormat, timeRange.HourMinuteTo)
	if err != nil {
		return 0, fmt.Errorf("invalid format for HourMinuteTo: %s", err)
	}
	if parsedHourTo.Before(parsedHourFrom) {
		parsedHourTo = parsedHourTo.AddDate(0, 0, 1)
	}
	return parsedHourTo.Sub(parsedHourFrom), nil
}

func getDurationBetweenWeekdays(timeRange TimeRange) (time.Duration, error) {
	weekdayFrom := timeRange.WeekdayFrom
	weekdayTo := timeRange.WeekdayTo
	if (weekdayFrom < 0 || weekdayFrom > 6) || (weekdayTo < 0 || weekdayTo > 6) {
		return 0, fmt.Errorf("one or both of the values are outside the range of 0 to 6")
	}
	days := calculateDaysBetweenWeekdays(int(timeRange.WeekdayFrom), int(timeRange.WeekdayTo))
	fromDateTime, err := constructDateTime(timeRange.HourMinuteFrom, 0)
	if err != nil {
		return 0, fmt.Errorf("error in constructing fromDateTime: %s", err)
	}
	toDateTime, err := constructDateTime(timeRange.HourMinuteTo, days)
	if err != nil {
		return 0, fmt.Errorf("error in constructing toDateTime: %s", err)
	}
	return toDateTime.Sub(fromDateTime), nil
}

func constructDateTime(hourMinute string, days int) (time.Time, error) {
	now := time.Now()
	dateTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	fromHour, err := strconv.Atoi(getHour(hourMinute))
	if err != nil {
		return dateTime, err
	}
	fromMinute, err := strconv.Atoi(getMinute(hourMinute))
	if err != nil {
		return dateTime, err
	}
	dateTime = dateTime.Add(time.Duration(fromHour+24*days)*time.Hour + time.Duration(fromMinute)*time.Minute)
	return dateTime, nil
}
func getDurationBetweenWeekDates(timeRange TimeRange, targetMonth time.Month, targetYear int) (time.Duration, error) {

	days := getDaysCountForNegativeDays(timeRange, targetMonth, targetYear)
	if timeRange.DayFrom > 0 && timeRange.DayTo > 0 && timeRange.DayFrom < timeRange.DayTo {
		days = (timeRange.DayTo) - (timeRange.DayFrom)
	}
	fromDateTime, err := constructDateTime(timeRange.HourMinuteFrom, 0)
	if err != nil {
		return 0, fmt.Errorf("error in constructing fromDateTime: %s", err)
	}
	toDateTime, err := constructDateTime(timeRange.HourMinuteTo, days)
	if err != nil {
		return 0, fmt.Errorf("error in constructing toDateTime: %s", err)
	}
	duration := toDateTime.Sub(fromDateTime)
	if duration < 0 {
		return 0, fmt.Errorf("hourMinuteFrom: %s or hourMinuteTo: %s is not valid", timeRange.HourMinuteFrom, timeRange.HourMinuteTo)
	}
	return duration, nil
}

func getDaysCountForNegativeDays(timeRange TimeRange, targetMonth time.Month, targetYear int) int {
	var days int
	var start, end time.Time
	now := time.Now()
	if timeRange.DayTo < timeRange.DayFrom {
		if timeRange.DayFrom > 0 {
			//27 , -2 april , 27, 28, 29
			//27 , -5 april , 27, 28, 29 .......next month
			timeRange.DayTo, _ = adjustDaysForMonth(timeRange.DayTo, targetMonth, targetYear, now)
			start, end = getStartAndEndTime(timeRange, targetMonth, now)
		} else {
			//-2 ,-4 april 29,30,1,.....28 may
			timeRange.DayFrom, _ = adjustDaysForMonth(timeRange.DayFrom, targetMonth, targetYear, now)
			timeRange.DayTo, _ = adjustDaysForMonth(timeRange.DayTo, targetMonth+1, targetYear, now)
			start, end = getStartAndEndTime(timeRange, targetMonth, now)
		}
	} else if timeRange.DayTo > timeRange.DayFrom {
		//-2 , -1 april 29 ,30
		if timeRange.DayTo < 0 {
			var lastDayOfMonth int
			timeRange.DayFrom, lastDayOfMonth = adjustDaysForMonth(timeRange.DayFrom, targetMonth, targetYear, now)
			timeRange.DayTo = lastDayOfMonth + timeRange.DayTo + 1
			start, end = getStartAndEndTime(timeRange, targetMonth, now)
		} else {
			//-2 , 4  april 29 , 30 , 1, 2,3,4 output 5
			timeRange.DayFrom, _ = adjustDaysForMonth(timeRange.DayFrom, targetMonth, targetYear, now)
			start, end = getStartAndEndTime(timeRange, targetMonth, now)
		}
	}
	days = int(end.Sub(start).Hours() / 24)
	return days
}

func getStartAndEndTime(timeRange TimeRange, targetMonth time.Month, now time.Time) (time.Time, time.Time) {
	start := getStartDate(timeRange, targetMonth, now)
	end := getEndDate(timeRange, targetMonth, now)
	if end.Day() < start.Day() && end.Month() == start.Month() && end.Year() == start.Year() {
		end = getEndDate(timeRange, targetMonth+1, now)
	}
	return start, end
}

func getEndDate(timeRange TimeRange, targetMonth time.Month, now time.Time) time.Time {
	return time.Date(now.Year(), targetMonth, timeRange.DayTo, 0, 0, 0, 0, now.Location())
}

func getStartDate(timeRange TimeRange, targetMonth time.Month, now time.Time) time.Time {
	return time.Date(now.Year(), targetMonth, timeRange.DayFrom, 0, 0, 0, 0, now.Location())
}
func adjustDaysForMonth(day int, targetMonth time.Month, targetYear int, now time.Time) (int, int) {
	lastDayOfMonth := getLastDayOfMonth(targetYear, targetMonth, now)
	if day > 0 {
		return lastDayOfMonth + day, lastDayOfMonth
	}
	return lastDayOfMonth + day + 1, lastDayOfMonth
}

func getLastDayOfMonth(targetYear int, targetMonth time.Month, now time.Time) int {
	firstDayOfNextMonth := time.Date(targetYear, targetMonth+1, 1, 0, 0, 0, 0, now.Location())
	lastDayOfMonth := firstDayOfNextMonth.Add(-time.Hour * 24).Day()
	return lastDayOfMonth
}
