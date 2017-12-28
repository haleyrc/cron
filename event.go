package cron

import "time"

// Event represents a single scheduled event. An event can run at any number of
// times, and will run for every combination of times in the three bit fields.
type Event struct {
	Day    Day
	Hour   Hour
	Minute Minute
}

// NewEvent creates an event for the specified days, with the specified hour and
// minute. Currently you can specify multiple days, but only a single value for
// hour and minute.
//
// TODO (RCH): Change this to accept bit fields for hour and minute cleanly.
func NewEvent(d Day, h uint32, m uint64) *Event {
	return &Event{Day: d, Hour: ParseHour(h), Minute: ParseMinute(m)}
}

// Trigger returns true if the event should run given the time t and false o.w.
func (evt *Event) Trigger(t time.Time) bool {
	d := ParseDay(t.Weekday())
	h := ParseHour(uint32(t.Hour()))
	m := ParseMinute(uint64(t.Minute()))
	return evt.Day.Has(d) && evt.Hour.Has(h) && evt.Minute.Has(m)
}
