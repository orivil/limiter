// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package limiter

import "time"

type WaitTimeProvider interface {
	// 根据(失败)次数获得等待时间
	GetWaitTime(times int64) (duration time.Duration)
}
