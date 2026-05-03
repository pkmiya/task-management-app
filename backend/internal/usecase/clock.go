package usecase

import "time"

// Clock は現在時刻を差し替え可能にする（要件 5.4・Phase 03 のテスト容易性）。
type Clock interface {
	Now() time.Time
}

type RealClock struct{}

func (RealClock) Now() time.Time {
	return time.Now().UTC()
}
