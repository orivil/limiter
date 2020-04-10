// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package limiter

import "time"

var Now = time.Now

type TimesLimiter struct {
	store Storage
	wt    WaitTimeProvider
}

func NewTimesLimiter(wt WaitTimeProvider, storage Storage) *TimesLimiter {
	return &TimesLimiter{wt: wt, store: storage}
}

func (ts *TimesLimiter) SetFailed(id string) (wait time.Duration, err error) {
	var (
		times    int64
		expireAt *time.Time
	)
	times, expireAt, err = ts.store.Get(id)
	if err != nil {
		return 0, err
	}
	now := Now()
	if expireAt != nil && expireAt.Before(now) {
		err = ts.store.Del(id)
		if err != nil {
			return 0, err
		}
		times = 0
	}
	times++
	wait = ts.wt.GetWaitTime(times)
	if wait > 0 {
		now = now.Add(wait)
		expireAt = &now
	} else {
		expireAt = nil
	}
	err = ts.store.Set(id, times, expireAt)
	return wait, err
}

func (ts *TimesLimiter) SetSuccess(id string) (err error) {
	return ts.store.Del(id)
}
