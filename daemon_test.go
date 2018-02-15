package cron

import (
	"context"
	"testing"
	"time"
)

type daemonSpy struct {
	called int
}

func (s *daemonSpy) run(_ time.Time) {
	s.called++
}

func TestDaemon(t *testing.T) {
	s := &daemonSpy{}
	count := 4
	ctx, cancel := context.WithCancel(context.Background())
	d := New(ctx, 1*time.Second, s.run)

	go func() {
		<-time.After(time.Duration(count) * time.Second)
		cancel()
	}()

	if err := d.Run(); err != nil {
		if err != ErrContextCanceled {
			t.Errorf("unexpected error: %v", err)
		}
	}

	if s.called != count {
		t.Errorf("expected function to be called %d times, but was called %d times", count, s.called)
	}
}
