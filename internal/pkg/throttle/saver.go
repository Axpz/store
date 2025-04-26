package throttle

import (
	"sync/atomic"
	"time"

	"golang.org/x/sync/singleflight"
)

type Saver struct {
	sf          singleflight.Group
	lastRunMs   int64
	minInterval int64
	pending     int32
}

func NewSaver(interval time.Duration) *Saver {
	return &Saver{
		minInterval: int64(interval / time.Millisecond),
	}
}

func (s *Saver) RequestSave(saveFn func() error) error {
	_, err, _ := s.sf.Do("save", func() (any, error) {
		now := time.Now().UnixMilli()
		last := atomic.LoadInt64(&s.lastRunMs)
		if now-last < s.minInterval {
			return nil, nil
		}

		if err := saveFn(); err != nil {
			return nil, err
		}

		atomic.StoreInt64(&s.lastRunMs, now)
		return nil, nil
	})
	return err
}

// RequestMustSave 尝试执行 saveFn。
// - 如果距离上次执行 ≥ interval：立即调用 saveFn，并更新 lastRunMs；
// - 否则：若尚未安排尾调用，则安排一个在剩余时间后执行，并置 pending=1。
func (s *Saver) RequestMustSave(saveFn func() error) error {
	now := time.Now().UnixMilli()
	last := atomic.LoadInt64(&s.lastRunMs)
	// 如果到达间隔，立即执行
	if now-last >= s.minInterval {
		if err := saveFn(); err != nil {
			return err
		}
		atomic.StoreInt64(&s.lastRunMs, time.Now().UnixMilli())
		return nil
	}

	// 间隔内：安排一次尾调用（只安排一次）
	if atomic.CompareAndSwapInt32(&s.pending, 0, 1) {
		wait := time.Duration(s.minInterval-(now-last)) * time.Millisecond
		go func() {
			time.Sleep(wait)
			// 真正执行尾调用
			_ = saveFn()
			atomic.StoreInt64(&s.lastRunMs, time.Now().UnixMilli())
			atomic.StoreInt32(&s.pending, 0)
		}()
	}

	return nil
}
