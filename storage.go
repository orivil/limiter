// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package limiter

import "time"

type Storage interface {
	// 获得次数及过期时间
	Get(id string) (times int64, expireAt *time.Time, err error)

	// 设置次数及过期时间
	Set(id string, times int64, expireAt *time.Time) error

	// 删除次数或设置为 0 次
	Del(id string) error
}
