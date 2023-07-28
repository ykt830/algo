package rate_limiter

import (
	"errors"
	"time"
)

// ChannelLimiter 利用 channel 进行限流，适用于对并发量有严格要求的场景
type ChannelLimiter struct {
	ch      chan struct{}
	timeout int64
}

var TimeoutErr = errors.New("wait timeout")

func NewChannelLimiter(limit int, timeout int64) *ChannelLimiter {
	return &ChannelLimiter{
		ch:      make(chan struct{}, limit),
		timeout: timeout,
	}
}

func (l *ChannelLimiter) Allow() error {
	timer := time.After(time.Duration(l.timeout) * time.Second)
	select {
	case <-timer:
		return TimeoutErr
	case l.ch <- struct{}{}:
	}

	return nil
}

func (l *ChannelLimiter) Release() {
	<-l.ch
}
