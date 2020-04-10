// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package times_limiter

type WaitTimeProvider interface {
	// 根据(失败)次数获得等待时间
	GetWaitTime(times int64) (duration int64)
}

type TimesLimiter struct {
	store Storage
	wt    WaitTimeProvider
}

func NewTimesLimiter(wt WaitTimeProvider, storage Storage) *TimesLimiter {
	return &TimesLimiter{wt: wt, store: storage}
}

func (ts *TimesLimiter) SetFailed(id string) (waitDuration int64, err error) {
	var times int64
	times, err = ts.store.Incr(id)
	if err != nil {
		return 0, err
	}
	return ts.wt.GetWaitTime(times), nil
}

func (ts *TimesLimiter) SetSuccess(id string) (err error) {
	return ts.store.Del(id)
}
