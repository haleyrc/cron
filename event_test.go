package cron_test

import (
	"testing"
	"time"

	"github.com/haleyrc/cron"
)

func TestEventIsNow(t *testing.T) {
	l, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Errorf("could not load location: %v", err)
		t.FailNow()
	}

	// Sunday
	sun := time.Date(2017, time.Month(12), 24, 20, 16, 0, 0, l)
	// Monday
	mon := time.Date(2017, time.Month(12), 25, 20, 16, 0, 0, l)
	// Tuesday
	tue := time.Date(2017, time.Month(12), 26, 20, 16, 0, 0, l)
	// Wednesday
	wed := time.Date(2017, time.Month(12), 27, 20, 16, 0, 0, l)
	// Thursday
	thurs := time.Date(2017, time.Month(12), 28, 20, 16, 0, 0, l)
	// Friday
	fri := time.Date(2017, time.Month(12), 29, 20, 16, 0, 0, l)
	// Saturday
	sat := time.Date(2017, time.Month(12), 30, 20, 16, 0, 0, l)

	evt := cron.NewEvent(cron.Wednesday|cron.Friday, 20, 16)
	tests := []struct {
		evt  *cron.Event
		Now  time.Time
		Want bool
	}{
		{evt, sun, false},
		{evt, mon, false},
		{evt, tue, false},
		{evt, wed, true},
		{evt, thurs, false},
		{evt, fri, true},
		{evt, sat, false},
		// Right day and minute, wrong hour
		{cron.NewEvent(cron.Wednesday, 19, 16), wed, false},
		// Right day and hour, wrong minute
		{cron.NewEvent(cron.Wednesday, 20, 17), wed, false},
	}

	for _, test := range tests {
		got := test.evt.Trigger(test.Now)
		if test.Want != got {
			if test.Want {
				t.Errorf("expected event to trigger, but didn't. event was %v, now was %s", test.evt, test.Now)
			} else {
				t.Errorf("expected event to not trigger, but did. event was %v, now was %s", test.evt, test.Now)
			}
		}
	}
}
