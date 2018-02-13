package cron_test

import (
	"testing"

	"github.com/haleyrc/cron"
)

func TestParseHour(t *testing.T) {
	want := cron.ParseHour(23)
	got := cron.ParseHour(24)
	if want != got {
		t.Errorf("upper bound failed: wanted %d, got %d", want, got)
	}
}

func TestHourHas(t *testing.T) {
	{
		h := cron.ParseHour(0)

		if !h.Has(h) {
			t.Errorf("expected hour to have self, but didn't")
		}
	}

	{
		h := cron.ParseHour(1) | cron.ParseHour(5)
		if !h.Has(cron.ParseHour(1)) {
			t.Errorf("expected %b to have 1, but didn't", h)
		}

		if !h.Has(cron.ParseHour(5)) {
			t.Errorf("expected %b to have 5, but didn't", h)
		}
	}
}
