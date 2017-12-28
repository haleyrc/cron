package cron

import "time"

const (
	// Sunday    000 0001
	Sunday = 1 << iota
	// Monday    000 0010
	Monday
	// Tuesday   000 0100
	Tuesday
	// Wednesday 000 1000
	Wednesday
	// Thursday  001 0000
	Thursday
	// Friday    010 0000
	Friday
	// Saturday  100 0000
	Saturday
)

const (
	// EveryDay is the mask containing all the days in the bit field:      111 1111.
	EveryDay = 0x7F
	// WeekDays is the mask containing only the weekdays in the bit field: 011 1110.
	WeekDays = 0x3E
	// WeekEnds is the mask containing only the weekends in teh bit field: 100 0001.
	WeekEnds = 0x41
)

// Day represents a bit field where the LSB represents Sunday, and the 7th bit
// represents Saturday.
type Day int8

// Has returns true if any of the bits specified by q is set in d.
func (d Day) Has(q Day) bool {
	return (d & q) > 0
}

// ParseDay converts weekdays from the time package to cron days as a bit field.
func ParseDay(d time.Weekday) Day {
	return Day(1 << uint(d))
}
