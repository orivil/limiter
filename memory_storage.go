// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package limiter

import (
	"sync"
	"time"
)

type MemoryStorage struct {
	times map[string]int64
	exps  map[string]*time.Time
	mu    sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		times: map[string]int64{},
		exps:  map[string]*time.Time{},
		mu:    sync.Mutex{},
	}
}

func (m *MemoryStorage) Get(id string) (times int64, expireAt *time.Time, err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	times = m.times[id]
	expireAt = m.exps[id]
	return
}

func (m *MemoryStorage) Set(id string, times int64, expireAt *time.Time) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.times[id] = times
	m.exps[id] = expireAt
	return nil
}

func (m *MemoryStorage) Del(id string) error {
	m.mu.Lock()
	delete(m.times, id)
	delete(m.exps, id)
	m.mu.Unlock()
	return nil
}
