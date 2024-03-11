package scheduler

import (
	"fmt"
	"strconv"
	"time"
)

func (tr TimeRange) getCron(lastDayOfMonth int) string {
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
	//day := strconv.Itoa(dayFrom)
	//if dayFrom == -1 {
	//	day = "L"
	//} else if dayFrom <= -2 && dayFrom >= -31 {
	//	day = fmt.Sprintf("L-%s", strconv.Itoa(-dayFrom-1))
	if dayFrom < 0 {
		dayFrom = lastDayOfMonth + 1 + dayFrom
	}
	day := strconv.Itoa(dayFrom)

	return fmt.Sprintf("%s %s %s * *", minute, hour, day)
}

//func (tr TimeRange) getCronExp(lastDayOfMonth int) string {
//	cronExp := tr.getCron(lastDayOfMonth)
//
//	if strings.Contains(cronExp, "L-2") {
//		lastDayOfMonth = lastDayOfMonth - 2
//		cronExp = strings.Replace(cronExp, "L-2", intToString(lastDayOfMonth), -1)
//	} else if strings.Contains(cronExp, "L-1") {
//		lastDayOfMonth = lastDayOfMonth - 1
//		cronExp = strings.Replace(cronExp, "L-1", intToString(lastDayOfMonth), -1)
//	} else {
//		cronExp = strings.Replace(cronExp, "L", intToString(lastDayOfMonth), -1)
//	}
//	return cronExp
//}
