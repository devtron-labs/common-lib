package scheduler

import (
	"github.com/robfig/cron/v3"
	"strings"
	"time"
)

func (tr TimeRange) GetScheduleSpec(targetTime time.Time) (nextWindowEdge time.Time, isTimeBetween bool, err error) {
	err = tr.ValidateTimeRange()
	if err != nil {
		return nextWindowEdge, false, err
	}
	if tr.Frequency == Fixed {
		nextWindowEdge, isTimeBetween = getScheduleForFixedTime(targetTime, tr)
		return nextWindowEdge, isTimeBetween, err
	}
	month, year := tr.getMonthAndYearForPreviousWindow(targetTime)
	cronExp := tr.getCronExp(year, month)
	parser := cron.NewParser(CRON)
	schedule, err := parser.Parse(cronExp)
	if err != nil {
		return nextWindowEdge, false, err
	}
	duration := tr.getDuration(month, year)

	windowStart, windowEnd := tr.getWindowStartAndEndTime(targetTime, duration, schedule)
	if isTimeInBetween(targetTime, windowStart, windowEnd) {
		return windowEnd, true, err
	}
	return windowStart, false, err
}

func (tr TimeRange) getWindowStartAndEndTime(targetTime time.Time, duration time.Duration, schedule cron.Schedule) (time.Time, time.Time) {
	var windowEnd time.Time

	prevDuration := duration
	if tr.isMonthOverlapping() {
		diff := getLastDayOfMonth(targetTime.Year(), targetTime.Month()) - getLastDayOfMonth(targetTime.Year(), targetTime.Month()-1)
		prevDuration = duration - time.Duration(diff)*time.Hour*24
	}

	timeMinusDuration := targetTime.Add(-1 * prevDuration)
	windowStart := schedule.Next(timeMinusDuration)
	windowEnd = windowStart.Add(duration)
	if !tr.TimeFrom.IsZero() && windowStart.Before(tr.TimeFrom) {
		windowStart = tr.TimeFrom
	}
	if !tr.TimeTo.IsZero() && windowEnd.After(tr.TimeTo) {
		windowEnd = tr.TimeTo
	}
	return windowStart, windowEnd
}

func (tr TimeRange) getCronExp(year int, month time.Month) string {
	cronExp := tr.getCron()
	lastDayOfMonth := getLastDayOfMonth(year, month)
	if strings.Contains(cronExp, "L-2") {
		lastDayOfMonth = lastDayOfMonth - 2
		cronExp = strings.Replace(cronExp, "L-2", intToString(lastDayOfMonth), -1)
	} else if strings.Contains(cronExp, "L-1") {
		lastDayOfMonth = lastDayOfMonth - 1
		cronExp = strings.Replace(cronExp, "L-1", intToString(lastDayOfMonth), -1)
	} else {
		cronExp = strings.Replace(cronExp, "L", intToString(lastDayOfMonth), -1)
	}
	return cronExp
}

// this will determine if the relevant year and month for the last window happens
// in the same month or previous month
func (tr TimeRange) getMonthAndYearForPreviousWindow(targetTime time.Time) (time.Month, int) {
	month := targetTime.Month()
	year := targetTime.Year()
	day := targetTime.Day()

	if tr.isMonthOverlapping() && tr.checkForOverlappingWindow(targetTime, day) {
		if month == 1 {
			month = 12
			year = year - 1
		} else {
			month = month - 1
		}
	}
	return month, year
}

func (tr TimeRange) checkForOverlappingWindow(targetTime time.Time, day int) bool {
	// for an overlapping window if the current time is on the latter part of the overlap then
	// we use the last month for calculation.

	if day < 1 {
		return false
	}
	return day < tr.DayTo || (day == tr.DayTo && tr.isToHourMinuteBeforeWindowEnd(targetTime))
}

func isTimeInBetween(timeCurrent, periodStart, periodEnd time.Time) bool {
	return (timeCurrent.After(periodStart) && timeCurrent.Before(periodEnd)) || timeCurrent.Equal(periodStart)
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
