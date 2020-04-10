// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package limiter

import (
	"math"
	"time"
)

type Options struct {
	// 第 1 次触发等待的等待时间, 随后每次按指数递增
	Wait time.Duration
	// 第 1 次触发等待的失败次数
	StartLimitTimes int64
}

func (ops *Options) GetWaitTime(times int64) (duration time.Duration) {
	if times >= ops.StartLimitTimes {
		z := times - ops.StartLimitTimes
		if z >= 63 {
			return math.MaxInt64
		}
		return ops.Wait << z
	}
	return 0
}
