package scheduler

import (
	"strconv"
	"strings"
	"time"
)

func getHour(hourMin string) string {
	return strings.Split(hourMin, ":")[0]
}

func getMinute(hourMin string) string {
	return strings.Split(hourMin, ":")[1]
}

func toString(weekday time.Weekday) string {
	return strconv.Itoa(int(weekday))
}

func intToString(value int) string {
	return strconv.Itoa(value)
}

func weekdaysToString(weekdays []time.Weekday) string {
	days := ""
	for _, day := range weekdays {
		days += toString(day) + ","
	}
	return days[:len(days)-1]
}

func calculateDaysBetweenWeekdays(from, to int) int {
	days := to - from
	if days < 0 {
		days += 7
	}
	return days
}
