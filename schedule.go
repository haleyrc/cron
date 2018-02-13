package schedule

import (
	"time"
)

// WithSchedule runs the provided function if shouldRun returns true.
func WithSchedule(f func(), shouldRun func(t time.Time) bool) func(time.Time) {
	return func(t time.Time) {
		if shouldRun(t) {
			f()
		}
	}
}

// OnWeekdays returns true if the current day is Monday-Friday.
func OnWeekdays(t time.Time) bool {
	wd := t.Weekday()
	if wd == time.Sunday || wd == time.Saturday {
		return false
	}
	return true
}

// OnWeekends returns true if the current day is Saturday or Sunday.
func OnWeekends(t time.Time) bool {
	return !OnWeekdays(t)
}

// OnSunday returns true if the current day is Sunday.
func OnSunday(t time.Time) bool {
	if t.Weekday() == time.Sunday {
		return true
	}
	return false
}

// OnMonday returns true if the current day is Monday.
func OnMonday(t time.Time) bool {
	if t.Weekday() == time.Monday {
		return true
	}
	return false
}

// OnTuesday returns true if the current day is Tuesday.
func OnTuesday(t time.Time) bool {
	if t.Weekday() == time.Tuesday {
		return true
	}
	return false
}

// OnWednesday returns true if the current day is Wednesday.
func OnWednesday(t time.Time) bool {
	if t.Weekday() == time.Wednesday {
		return true
	}
	return false
}

// OnThursday returns true if the current day is Thursday.
func OnThursday(t time.Time) bool {
	if t.Weekday() == time.Thursday {
		return true
	}
	return false
}

// OnFriday returns true if the current day is Friday.
func OnFriday(t time.Time) bool {
	if t.Weekday() == time.Friday {
		return true
	}
	return false
}

// OnSaturday returns true if the current day is Saturday.
func OnSaturday(t time.Time) bool {
	if t.Weekday() == time.Saturday {
		return true
	}
	return false
}
