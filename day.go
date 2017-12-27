package cron

const (
	Sunday = 1 << iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	EveryDay = 127
	WeekDays = 62
	WeekEnds = 65
)

type Day int8

func (d Day) Has(q Day) bool {
	return (d & q) > 0
}

type Event struct {
	Day Day
}
