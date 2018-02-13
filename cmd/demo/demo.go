package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/haleyrc/cron"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var t tracker
	funcs := []func(time.Time){
		say,
		t.sayOnce,
	}

	d := cron.New(ctx, 1*time.Second, funcs...)
	if err := d.Run(); err != nil {
		log.Fatalln(err)
	}
}

type tracker struct {
	ran bool
}

func (t *tracker) sayOnce(_ time.Time) {
	if t.ran {
		return
	}
	fmt.Println("I'm saying it once")
	t.ran = true
}

func say(_ time.Time) {
	fmt.Println("I'm saying something")
}
