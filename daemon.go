package cron

import (
	"context"
	"errors"
	"time"
)

// ContextCanceled indicates that the client signaled the Daemon to finish.
var ContextCanceled = errors.New("context canceled by parent")

// Daemon is a simple structure for running multiple functions at a given
// interval. It can be stopped by closing the cancel channel, usually by calling
// the genesis context's cancel function. On every tick of the provided interval,
// all of the jobs are called, being passed the current time so that they can
// determine whether to run or not.
type Daemon struct {
	interval time.Duration
	cancel   <-chan struct{}
	jobs     []func(time.Time)
}

// New creates a new Daemon. It listens for a close on the channel provided by
// the context's Done method, otherwise it ticks at the provided interval,
// calling each of the provided jobs in turn.
func New(ctx context.Context, d time.Duration, jobs ...func(time.Time)) *Daemon {
	return &Daemon{
		interval: d,
		cancel:   ctx.Done(),
		jobs:     jobs,
	}
}

// Interval returns the daemon's internal interval value.
func (d *Daemon) Interval() time.Duration { return d.interval }

// Run is the main event loop for the daemon. It performs the work of listening
// on a ticker channel, as well as the cancel channel. If cancel is closed, it
// stops the ticker and return an error signaling an expected cancel. Otherwise,
// if the ticker ticks, it calls all of the Daemon's jobs in turn, passing the
// current time.
func (d *Daemon) Run() error {
	ticker := time.NewTicker(d.interval)
	for {
		select {
		case <-d.cancel:
			ticker.Stop()
			return ContextCanceled
		case <-ticker.C:
			t := time.Now()
			for _, f := range d.jobs {
				f(t)
				time.Sleep(100 * time.Millisecond)
			}
		default:
		}
	}
}
