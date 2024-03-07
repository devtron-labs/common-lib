package scheduler

import (
	"errors"
	"golang.org/x/exp/slices"
	"strings"
	"unicode"
)

func (tr TimeRange) ValidateTimeRange() error {
	if !slices.Contains(AllowedFrequencies, tr.Frequency) {
		return errors.New("invalid Frequency type")
	}
	if tr.Frequency != FIXED {
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
	case DAILY:
		if tr.HourMinuteFrom == "" || tr.HourMinuteTo == "" {
			return errors.New("HourMinuteFrom and HourMinuteTo must be present for DAILY frequency")
		}
	case FIXED:
		if tr.TimeFrom.IsZero() || tr.TimeTo.IsZero() {
			return errors.New("TimeFrom and TimeTo must be present for FIXED frequency")
		}
		if tr.TimeFrom.After(tr.TimeTo) {
			return errors.New("TimeFrom must be less than TimeTo for FIXED frequency")
		}
		if tr.TimeFrom == tr.TimeTo {
			return errors.New("TimeFrom must not be equal to TimeTo for FIXED frequency")
		}
	case WEEKLY:
		if len(tr.Weekdays) == 0 {
			return errors.New("weekdays, must be present for WEEKLY frequency")
		}
	case WEEKLY_RANGE:
		if tr.WeekdayFrom == 0 || tr.WeekdayTo == 0 {
			return errors.New("WeekdayFrom, must be present for WEEKLY_RANGE frequency")
		}
		if (tr.WeekdayFrom < 0 || tr.WeekdayFrom > 6) || (tr.WeekdayTo < 0 || tr.WeekdayTo > 6) {
			return errors.New("one or both of the values are outside the range of 0 to 6")
		}
	case MONTHLY:
		if tr.DayFrom == 0 || tr.DayTo == 0 {
			return errors.New("DayFrom, DayTo, must be present for MONTHLY frequency")
		}
		if tr.DayFrom < 0 && tr.DayTo < 0 && tr.DayFrom > tr.DayTo {
			return errors.New("invalid value of DayFrom or DayTo,DayFrom and DayTo is less than zero and  DayFrom > DayTo")
		}
		if (tr.DayTo < 0 && tr.DayFrom > 0 && tr.DayFrom > 29+tr.DayTo) || tr.DayTo < -3 || tr.DayFrom < -3 {
			return errors.New("invalid value of DayFrom or DayTo")
		}
		if tr.DayFrom == tr.DayTo {
			return errors.New("invalid value , DayFrom must not be equal to DayTo")
		}
	}
	return nil
}

func validateHourMinute(HourMinute string) error {
	for _, char := range HourMinute {
		if !unicode.IsDigit(char) && char != ':' {
			return errors.New("HourMinute is not valid,must not contain character other than integer and colon(:)")
		}
	}
	return nil
}
