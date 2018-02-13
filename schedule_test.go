package schedule

import (
	"fmt"
	"testing"
	"time"
)

var (
	tz        = time.FixedZone("America/New_York", -17762)
	Sunday    = time.Date(2018, time.Month(02), 11, 0, 0, 0, 0, tz)
	Monday    = time.Date(2018, time.Month(02), 12, 0, 0, 0, 0, tz)
	Tuesday   = time.Date(2018, time.Month(02), 13, 0, 0, 0, 0, tz)
	Wednesday = time.Date(2018, time.Month(02), 14, 0, 0, 0, 0, tz)
	Thursday  = time.Date(2018, time.Month(02), 15, 0, 0, 0, 0, tz)
	Friday    = time.Date(2018, time.Month(02), 16, 0, 0, 0, 0, tz)
	Saturday  = time.Date(2018, time.Month(02), 17, 0, 0, 0, 0, tz)
)

type spy struct {
	wasCalled bool
}

func (s *spy) run() {
	s.wasCalled = true
}

func (s *spy) init() {
	s.wasCalled = false
}

type testcase struct {
	name string
	day  time.Time
	want bool
}

type suite struct {
	f     func(time.Time) bool
	tests []testcase
}

func TestWithSchedule(t *testing.T) {
	suites := []suite{
		suite{
			f: OnSunday,
			tests: []testcase{
				{"Sunday", Sunday, true},
				{"Monday", Monday, false},
				{"Tuesday", Tuesday, false},
				{"Wednesday", Wednesday, false},
				{"Thursday", Thursday, false},
				{"Friday", Friday, false},
				{"Saturday", Saturday, false},
			},
		},
		suite{
			f: OnMonday,
			tests: []testcase{
				{"Sunday", Sunday, false},
				{"Monday", Monday, true},
				{"Tuesday", Tuesday, false},
				{"Wednesday", Wednesday, false},
				{"Thursday", Thursday, false},
				{"Friday", Friday, false},
				{"Saturday", Saturday, false},
			},
		},
		suite{
			f: OnTuesday,
			tests: []testcase{
				{"Sunday", Sunday, false},
				{"Monday", Monday, false},
				{"Tuesday", Tuesday, true},
				{"Wednesday", Wednesday, false},
				{"Thursday", Thursday, false},
				{"Friday", Friday, false},
				{"Saturday", Saturday, false},
			},
		},
		suite{
			f: OnWednesday,
			tests: []testcase{
				{"Sunday", Sunday, false},
				{"Monday", Monday, false},
				{"Tuesday", Tuesday, false},
				{"Wednesday", Wednesday, true},
				{"Thursday", Thursday, false},
				{"Friday", Friday, false},
				{"Saturday", Saturday, false},
			},
		},
		suite{
			f: OnThursday,
			tests: []testcase{
				{"Sunday", Sunday, false},
				{"Monday", Monday, false},
				{"Tuesday", Tuesday, false},
				{"Wednesday", Wednesday, false},
				{"Thursday", Thursday, true},
				{"Friday", Friday, false},
				{"Saturday", Saturday, false},
			},
		},
		suite{
			f: OnFriday,
			tests: []testcase{
				{"Sunday", Sunday, false},
				{"Monday", Monday, false},
				{"Tuesday", Tuesday, false},
				{"Wednesday", Wednesday, false},
				{"Thursday", Thursday, false},
				{"Friday", Friday, true},
				{"Saturday", Saturday, false},
			},
		},
		suite{
			f: OnSaturday,
			tests: []testcase{
				{"Sunday", Sunday, false},
				{"Monday", Monday, false},
				{"Tuesday", Tuesday, false},
				{"Wednesday", Wednesday, false},
				{"Thursday", Thursday, false},
				{"Friday", Friday, false},
				{"Saturday", Saturday, true},
			},
		},
		suite{
			f: OnWeekdays,
			tests: []testcase{
				{"Sunday", Sunday, false},
				{"Monday", Monday, true},
				{"Tuesday", Tuesday, true},
				{"Wednesday", Wednesday, true},
				{"Thursday", Thursday, true},
				{"Friday", Friday, true},
				{"Saturday", Saturday, false},
			},
		},
		suite{
			f: OnWeekends,
			tests: []testcase{
				{"Sunday", Sunday, true},
				{"Monday", Monday, false},
				{"Tuesday", Tuesday, false},
				{"Wednesday", Wednesday, false},
				{"Thursday", Thursday, false},
				{"Friday", Friday, false},
				{"Saturday", Saturday, true},
			},
		},
	}

	for _, s := range suites {
		Suite(t, s)
	}
}

func Suite(t *testing.T, suite suite) {
	s := &spy{}
	f := WithSchedule(s.run, suite.f)
	for _, tst := range suite.tests {
		s.init()
		f(tst.day)
		if tst.want != s.wasCalled {
			if tst.want {
				t.Errorf("wanted function to be called on %v, but wasn't", tst.name)
				continue
			}
			t.Errorf("didn't want function to be called on %v, but was", tst.name)
		}
	}
}

func ExampleWithSchedule_withMemory() {
	Monday := time.Date(2018, time.Month(02), 12, 0, 0, 0, 0, tz)
	NextMonday := time.Date(2018, time.Month(02), 19, 0, 0, 0, 0, tz)
	OnceOnMonday := func(last time.Time, buffer time.Duration) func(time.Time) bool {
		return func(t time.Time) bool {
			bottom := t.Add(time.Duration(-1) * buffer)
			shouldRun := last.Before(bottom) && OnMonday(t)

			if shouldRun {
				last = time.Now()
			}

			return shouldRun
		}
	}

	var last time.Time
	f := WithSchedule(func() { fmt.Println("I ran") }, OnceOnMonday(last, 23*time.Hour))

	f(Monday)
	f(Monday)
	f(NextMonday)

	// Output: I ran
	// I ran
}

func ExampleWithSchedule_didRun() {
	Monday := time.Date(2018, time.Month(02), 12, 0, 0, 0, 0, tz)
	f := WithSchedule(func() { fmt.Println("I ran") }, OnMonday)
	f(Monday)

	// Output: I ran
}

func ExampleWithSchedule_didntRunWrongDay() {
	Tuesday := time.Date(2018, time.Month(02), 13, 0, 0, 0, 0, tz)
	f := WithSchedule(func() { fmt.Println("I ran") }, OnMonday)
	f(Tuesday)

	// Output:
}

func ExampleWithSchedule_withHistory() {
	Monday := time.Date(2018, time.Month(02), 12, 0, 0, 0, 0, tz)
	OnceOnMonday := func(last time.Time, buffer time.Duration) func(time.Time) bool {
		return func(t time.Time) bool {
			bottom := t.Add(time.Duration(-1) * buffer)
			shouldRun := last.Before(bottom) && OnMonday(t)

			if shouldRun {
				last = time.Now()
			}

			return shouldRun
		}
	}

	var last time.Time
	f := WithSchedule(func() { fmt.Println("I ran") }, OnceOnMonday(last, 23*time.Hour))
	f(Monday)

	// Output: I ran
}
