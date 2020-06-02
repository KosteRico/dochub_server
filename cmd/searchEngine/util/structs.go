package util

import (
	"sync/atomic"
	"time"
)

type MaxGoroutinesLimiter struct {
	counter         int32
	goroutinesLimit int32
}

func NewMaxGoroutinesLimiter(limit int32) *MaxGoroutinesLimiter {
	return &MaxGoroutinesLimiter{
		counter:         0,
		goroutinesLimit: limit,
	}
}

func (mgl *MaxGoroutinesLimiter) Increment() {
	atomic.AddInt32(&mgl.counter, 1)
}

func (mgl *MaxGoroutinesLimiter) Decrement() {
	atomic.AddInt32(&mgl.counter, -1)
}

func (mgl *MaxGoroutinesLimiter) Counter() int32 {
	return atomic.LoadInt32(&mgl.counter)
}

func (mgl *MaxGoroutinesLimiter) Wait() {
	for mgl.Counter() >= mgl.goroutinesLimit {
		time.Sleep(1 * time.Nanosecond)
	}
}
