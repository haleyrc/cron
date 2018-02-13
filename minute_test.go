package cron_test

import (
	"testing"

	"github.com/haleyrc/cron"
)

func TestParseMinute(t *testing.T) {
	want := cron.ParseMinute(59)
	got := cron.ParseMinute(60)
	if want != got {
		t.Errorf("upper bound failed: wanted %d, got %d", want, got)
	}
}

func TestMinuteHas(t *testing.T) {
	{
		h := cron.ParseMinute(0)

		if !h.Has(h) {
			t.Errorf("expected minute to have self, but didn't")
		}
	}

	{
		h := cron.ParseMinute(1) | cron.ParseMinute(5)
		if !h.Has(cron.ParseMinute(1)) {
			t.Errorf("expected %b to have 1, but didn't", h)
		}

		if !h.Has(cron.ParseMinute(5)) {
			t.Errorf("expected %b to have 5, but didn't", h)
		}
	}
}
