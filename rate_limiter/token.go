package rate_limiter

import "time"

type TokenLimiter struct {
	interval time.Duration
	bucket   chan struct{}
}

func NewTokenLimiter(limit int, interval time.Duration) *TokenLimiter {
	l := TokenLimiter{
		interval: interval,
		bucket:   make(chan struct{}, limit),
	}

	for i := 0; i < limit-1; i++ {
		l.bucket <- struct{}{}
	}

	go l.addToken()

	return &l
}

func (l *TokenLimiter) Allow() bool {
	select {
	case <-l.bucket:
		return true
	default:
		return false
	}
}

func (l *TokenLimiter) addToken() {
	for {
		select {
		case <-time.After(l.interval):
			l.bucket <- struct{}{}
		}
	}
}
