package cron_test

import "testing"
import "github.com/haleyrc/cron"

func TestHasDayIdentity(t *testing.T) {
	tests := []struct {
		Title string
		Sub   cron.Day
		Obj   cron.Day
		Want  bool
	}{
		{Title: "Sunday", Sub: cron.Sunday, Obj: cron.Sunday, Want: true},
		{Title: "Monday", Sub: cron.Monday, Obj: cron.Monday, Want: true},
		{Title: "Tuesday", Sub: cron.Tuesday, Obj: cron.Tuesday, Want: true},
		{Title: "Wednesday", Sub: cron.Wednesday, Obj: cron.Wednesday, Want: true},
		{Title: "Thursday", Sub: cron.Thursday, Obj: cron.Thursday, Want: true},
		{Title: "Friday", Sub: cron.Friday, Obj: cron.Friday, Want: true},
		{Title: "Saturday", Sub: cron.Saturday, Obj: cron.Saturday, Want: true},
		{Title: "Sunday", Sub: 1, Obj: cron.Sunday, Want: true},
		{Title: "Monday", Sub: 2, Obj: cron.Monday, Want: true},
		{Title: "Tuesday", Sub: 4, Obj: cron.Tuesday, Want: true},
		{Title: "Wednesday", Sub: 8, Obj: cron.Wednesday, Want: true},
		{Title: "Thursday", Sub: 16, Obj: cron.Thursday, Want: true},
		{Title: "Friday", Sub: 32, Obj: cron.Friday, Want: true},
		{Title: "Saturday", Sub: 64, Obj: cron.Saturday, Want: true},
	}

	for _, test := range tests {
		got := test.Sub.Has(test.Obj)
		if test.Want != got {
			t.Errorf("%s: comparison failed: %d doesn't have %d", test.Title, test.Sub, test.Obj)
		}
	}
}

func TestHasDayCombination(t *testing.T) {
	tests := []struct {
		Title string
		Sub   cron.Day
		Objs  []cron.Day
		Want  bool
	}{
		{
			Title: "Sunday+Monday Has",
			Sub:   cron.Sunday | cron.Monday,
			Objs:  []cron.Day{cron.Sunday, cron.Monday},
			Want:  true,
		},
		{
			Title: "Sunday+Monday Doesn't Have",
			Sub:   cron.Sunday | cron.Monday,
			Objs:  []cron.Day{cron.Tuesday, cron.Wednesday, cron.Thursday, cron.Friday, cron.Saturday},
			Want:  false,
		},
		{
			Title: "Saturday+Sunday Has",
			Sub:   cron.Saturday | cron.Sunday,
			Objs:  []cron.Day{cron.Sunday, cron.Saturday},
			Want:  true,
		},
		{
			Title: "Saturday+Sunday Doesn't Have",
			Sub:   cron.Saturday | cron.Sunday,
			Objs:  []cron.Day{cron.Monday, cron.Tuesday, cron.Wednesday, cron.Thursday, cron.Friday},
			Want:  false,
		},
		{
			Title: "All Days",
			Sub:   cron.EveryDay,
			Objs:  []cron.Day{cron.Sunday, cron.Monday, cron.Tuesday, cron.Wednesday, cron.Thursday, cron.Friday, cron.Saturday},
			Want:  true,
		},
		{
			Title: "Week Days Has",
			Sub:   cron.WeekDays,
			Objs:  []cron.Day{cron.Monday, cron.Tuesday, cron.Wednesday, cron.Thursday, cron.Friday},
			Want:  true,
		},
		{
			Title: "Week Days Doesn't Have",
			Sub:   cron.WeekDays,
			Objs:  []cron.Day{cron.Sunday, cron.Saturday},
			Want:  false,
		},
		{
			Title: "Week Ends Has",
			Sub:   cron.WeekEnds,
			Objs:  []cron.Day{cron.Sunday, cron.Saturday},
			Want:  true,
		},
		{
			Title: "Week Ends Has",
			Sub:   cron.WeekEnds,
			Objs:  []cron.Day{cron.Monday, cron.Tuesday, cron.Wednesday, cron.Thursday, cron.Friday},
			Want:  false,
		},
	}

	for _, test := range tests {
		for _, obj := range test.Objs {
			got := test.Sub.Has(obj)
			if test.Want != got {
				t.Errorf("%s: comparison failed: %d doesn't have %d", test.Title, test.Sub, obj)
			}
		}
	}
}
