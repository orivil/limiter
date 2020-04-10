// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package times_limiter

import (
	"github.com/go-redis/redis"
	"sync"
)

type Storage interface {
	// 增加一次, 并返回总次数, 默认为 0 次
	Incr(id string) (times int64, err error)

	// 删除次数或设置为 0 次
	Del(id string) error
}

type MemoryStorage struct {
	times map[string]int64
	mu    sync.Mutex
}

type RedisStorage struct {
	client *redis.Client
}

func (r *RedisStorage) Incr(id string) (times int64, err error) {
	return r.client.Incr(id).Result()
}

func (r *RedisStorage) Del(id string) error {
	return r.client.Del(id).Err()
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		times: make(map[string]int64, 10),
		mu:    sync.Mutex{},
	}
}

func NewRedisStorage(client *redis.Client) *RedisStorage {
	return &RedisStorage{client: client}
}

func (m *MemoryStorage) Incr(id string) (times int64, err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.times[id]++
	return m.times[id], nil
}

func (m *MemoryStorage) Del(id string) error {
	m.mu.Lock()
	delete(m.times, id)
	m.mu.Unlock()
	return nil
}
