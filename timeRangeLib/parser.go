package timeRangeLib

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func (tr TimeRange) GetTimeRangeWindow(targetTime time.Time) (nextWindowEdge time.Time, isTimeBetween bool, err error) {
	err = tr.ValidateTimeRange()
	if err != nil {
		return nextWindowEdge, false, err
	}
	windowStart, windowEnd, err := tr.getWindowForTargetTime(targetTime)
	if err != nil {
		return nextWindowEdge, isTimeBetween, err
	}
	if isTimeInBetween(targetTime, windowStart, windowEnd) {
		return windowEnd, true, nil
	}
	return windowStart, false, nil
}

func (tr TimeRange) getWindowForTargetTime(targetTime time.Time) (time.Time, time.Time, error) {

	if tr.Frequency == Fixed {
		windowStart, windowEnd := tr.getWindowForFixedTime(targetTime)
		return windowStart, windowEnd, nil
	}
	return tr.getWindowStartAndEndTime(targetTime)
}

// here target time is required to handle exceptions in monthly
// frequency where current time determines the cron and duration
func (tr TimeRange) getCronScheduleAndDuration(targetTime time.Time) (cron.Schedule, time.Duration, error) {

	evaluator := tr.getTimeRangeExpressionEvaluator(targetTime)
	cronExp := evaluator.getCron()
	parser := cron.NewParser(CRON)
	schedule, err := parser.Parse(cronExp)
	if err != nil {
		return nil, 0, fmt.Errorf("error parsing cron expression %s %v", cronExp, err)
	}
	return schedule, evaluator.getDuration(), nil
}

func (tr TimeRange) getWindowStartAndEndTime(targetTime time.Time) (time.Time, time.Time, error) {

	var windowEnd time.Time
	schedule, duration, err := tr.getCronScheduleAndDuration(targetTime)
	if err != nil {
		return windowEnd, windowEnd, err
	}

	timeMinusDuration := tr.currentTimeMinusWindowDuration(targetTime, duration)
	windowStart := schedule.Next(timeMinusDuration)
	windowEnd = windowStart.Add(duration)

	windowStart, windowEnd = tr.applyStartEndBoundary(windowStart, windowEnd)
	return windowStart, windowEnd, nil
}

func (tr TimeRange) applyStartEndBoundary(windowStart time.Time, windowEnd time.Time) (time.Time, time.Time) {
	if !tr.TimeFrom.IsZero() && windowStart.Before(tr.TimeFrom) {
		windowStart = tr.TimeFrom
	}
	if !tr.TimeTo.IsZero() && windowEnd.After(tr.TimeTo) {
		windowEnd = tr.TimeTo
	}
	return windowStart, windowEnd
}

func (tr TimeRange) currentTimeMinusWindowDuration(targetTime time.Time, duration time.Duration) time.Time {

	prevDuration := tr.getTimeRangeExpressionEvaluator(targetTime).getDurationOfPreviousWindow(duration)
	return targetTime.Add(-1 * prevDuration)
}

func (tr TimeRange) getWindowForFixedTime(targetTime time.Time) (time.Time, time.Time) {
	var windowStartOrEnd time.Time
	if targetTime.After(tr.TimeTo) {
		return windowStartOrEnd, windowStartOrEnd
	}
	return tr.TimeFrom, tr.TimeTo
}
