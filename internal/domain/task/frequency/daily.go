package frequency

import (
	"fmt"
	"time"
)

type DailyStrategy struct {
	IntervalDays int `json:"interval_days"`
}

func (s *DailyStrategy) IsValid(date time.Time) bool {
	if s.IntervalDays <= 0 {
		return false
	}

	daysSinceEpoche := int(date.Unix() / (24 * 60 * 60))

	return daysSinceEpoche%s.IntervalDays == 0
}

func (s *DailyStrategy) NextOccurrence(from time.Time) time.Time {
	return from.AddDate(0, 0, s.IntervalDays)
}

func (s *DailyStrategy) Description() string {
	if s.IntervalDays == 1 {
		return "Ежедневно"
	}

	dayForm := pluralizeDays(s.IntervalDays)
	return fmt.Sprintf("Каждые %d %s", s.IntervalDays, dayForm)
}

func pluralizeDays(n int) string {
	lastDigit := n % 10

	switch lastDigit {
	case 1:
		return "день"
	case 2, 3, 4:
		return "дня"
	default:
		return "дней"
	}
}

func (s *DailyStrategy) Validate() error {
	if s.IntervalDays < 1 {
		return fmt.Errorf("interval_days must be >= 1")
	}
	return nil
}
