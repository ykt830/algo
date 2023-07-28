package rate_limiter

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type SlidingWindowLimiter struct {
	mutex sync.Mutex

	windows    []int32
	curWindow  int32
	windowsSum int32
	limit      int32
	snippet    time.Duration
}

func NewSlidingWindowLimiter(limit int32, interval time.Duration, windowCnt int32) *SlidingWindowLimiter {
	l := SlidingWindowLimiter{
		windows:    make([]int32, windowCnt),
		curWindow:  0,
		windowsSum: 0,
		limit:      limit,
		snippet:    (interval) / time.Duration(windowCnt),
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Print(err)
			}
		}()

		l.slide()
	}()

	return &l
}

func (l *SlidingWindowLimiter) Allow() bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.windowsSum > l.limit {
		return false
	}

	l.windowsSum++
	l.curWindow++

	return true
}

func (l *SlidingWindowLimiter) slide() {
	for {
		select {
		case <-time.After(l.snippet):
			l.mutex.Lock()
			oldWindow := l.windows[0]
			newWindow := atomic.SwapInt32(&l.curWindow, 0)
			atomic.AddInt32(&l.windowsSum, -oldWindow)
			l.mutex.Unlock()

			l.windows = append(l.windows, newWindow)
			l.windows = l.windows[1:len(l.windows)]
		}
	}
}
