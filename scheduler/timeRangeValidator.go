package scheduler

import (
	"errors"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

func (tr TimeRange) ValidateTimeRange() error {
	if !slices.Contains(AllowedFrequencies, tr.Frequency) {
		return errors.New("invalid Frequency type")
	}
	if tr.Frequency != Fixed {
		colonCountFrom := strings.Count(tr.HourMinuteFrom, ":")
		if colonCountFrom != 1 {
			return errors.New("invalid format: must contain exactly one colon")
		}
		colonCountTo := strings.Count(tr.HourMinuteTo, ":")
		if colonCountTo != 1 {
			return errors.New("invalid format: must contain exactly one colon")
		}
		err := validateHourMinute(tr.HourMinuteFrom)
		if err != nil {
			return err
		}
		err = validateHourMinute(tr.HourMinuteTo)
		if err != nil {
			return err
		}

	}
	switch tr.Frequency {
	case Daily:
		if tr.HourMinuteFrom == "" || tr.HourMinuteTo == "" {
			return errors.New("HourMinuteFrom and HourMinuteTo must be present for Daily frequency")
		}
	case Fixed:
		if tr.TimeFrom.IsZero() || tr.TimeTo.IsZero() {
			return errors.New("TimeFrom and TimeTo must be present for Fixed frequency")
		}
		if tr.TimeFrom.After(tr.TimeTo) {
			return errors.New("TimeFrom must be less than TimeTo for Fixed frequency")
		}
		if tr.TimeFrom == tr.TimeTo {
			return errors.New("TimeFrom must not be equal to TimeTo for Fixed frequency")
		}
	case Weekly:
		if len(tr.Weekdays) == 0 {
			return errors.New("weekdays, must be present for Weekly frequency")
		}
	case WeeklyRange:
		if tr.WeekdayFrom == 0 || tr.WeekdayTo == 0 {
			return errors.New("WeekdayFrom, must be present for WeeklyRange frequency")
		}
		if (tr.WeekdayFrom < 0 || tr.WeekdayFrom > 6) || (tr.WeekdayTo < 0 || tr.WeekdayTo > 6) {
			return errors.New("one or both of the values are outside the range of 0 to 6")
		}
	case Monthly:
		if tr.DayFrom == 0 || tr.DayTo == 0 {
			return errors.New("DayFrom, DayTo, must be present for Monthly frequency")
		}
		// this is to prevent overlapping windows crossing to next month for both negatives
		if tr.DayFrom < 0 && tr.DayTo < 0 && tr.DayFrom > tr.DayTo {
			return errors.New("invalid value of DayFrom or DayTo,DayFrom and DayTo is less than zero and  DayFrom > DayTo")
		}
		// this is an edge case where with negative 'to' date results into a date before the 'from' date
		// example: 26,-4 will pe prevented because for February it will become invalid
		// also currently max negative supported is third last day of the month
		if (tr.DayTo < 0 && tr.DayFrom > 0 && tr.DayFrom > 29+tr.DayTo) || tr.DayTo < -3 || tr.DayFrom < -3 {
			return errors.New("invalid value of DayFrom or DayTo")
		}
	}
	return nil
}

func validateHourMinute(hourMinute string) error {
	parts := strings.Split(hourMinute, ":")
	if len(parts) != 2 {
		return errors.New("HourMinute is not valid, should be strictly of format HH:MM")
	}
	hh, err := strconv.Atoi(parts[0])
	if err != nil {
		return errors.New("Hour is not valid" + parts[0])
	}
	if hh > 23 || hh < 0 {
		return errors.New("Hour is not valid" + parts[0])
	}

	mm, err := strconv.Atoi(parts[1])
	if err != nil {
		return errors.New("Hour is not valid" + parts[1])
	}
	if mm > 59 || mm < 0 {
		return errors.New("Hour is not valid" + parts[0])
	}
	return nil
}
