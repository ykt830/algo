package rate_limiter

import (
	"sync/atomic"
	"time"
)

type CountLimiter struct {
	interval time.Duration
	limit    int32
	counter  int32
}

func NewCountLimiter(interval time.Duration, limit int32) *CountLimiter {
	l := CountLimiter{
		interval: interval,
		limit:    limit,
		counter:  0,
	}

	go l.daemon()

	return &l
}

func (l *CountLimiter) Allow() bool {
	return l.AllowN(1)
}

func (l *CountLimiter) AllowN(n int32) bool {
	return atomic.AddInt32(&l.counter, n) < l.limit
}

func (l *CountLimiter) daemon() {
	for {
		select {
		case <-time.After(l.interval):
			atomic.StoreInt32(&l.counter, 0)
		}
	}
}
