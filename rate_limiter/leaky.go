package rate_limiter

import "time"

type Worker interface {
	work()
}

type LeakyLimiter struct {
	interval time.Duration

	bucket chan Worker
}

func NewLeakyLimiter(limit int32, interval time.Duration) *LeakyLimiter {
	return &LeakyLimiter{
		interval: interval,
		bucket:   make(chan Worker, limit),
	}
}

func (l *LeakyLimiter) AddTask(worker Worker) bool {
	select {
	case l.bucket <- worker:
		return true
	default:
		return false
	}
}

func (l *LeakyLimiter) Run() {
	for {
		select {
		case <-time.After(l.interval):
			worker := <-l.bucket
			worker.work()
		}
	}

}
