// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package limiter_test

import (
	"github.com/orivil/limiter"
	"testing"
	"time"
)

var opt = limiter.Options{
	Wait:            60 * time.Second,
	StartLimitTimes: 5,
}

var waitDurations = []time.Duration{
	0, 0, 0, 0, 60, 120, 240, 480, 960,
}

func TestNewTimesLimiter(t *testing.T) {
	storage := limiter.NewMemoryStorage()
	timesLimiter := limiter.NewTimesLimiter(&opt, storage)
	for i, duration := range waitDurations {
		duration = duration * time.Second
		wd, err := timesLimiter.SetFailed("tony")
		if err != nil {
			panic(err)
		}
		if wd != duration {
			t.Errorf("index %d: need: %d, got: %d\n", i, duration, wd)
		}
		wd, err = timesLimiter.SetFailed("money")
		if err != nil {
			panic(err)
		}
		if wd != duration {
		}
	}
	limiter.Now = func() time.Time {
		return time.Now().Add(961 * time.Second)
	}
	wd, err := timesLimiter.SetFailed("tony")
	if err != nil {
		panic(err)
	}
	if wd != 0 {
		t.Errorf("need: %d, got: %d\n", 0, wd)
	}
}
