package frequency

import (
	"fmt"
	"time"
)

type MonthlyStrategy struct {
	DayOfMonth int `json:"day_of_month"`
}

func (s *MonthlyStrategy) IsValid(date time.Time) bool {
	if s.DayOfMonth < 1 || s.DayOfMonth > 30 {
		return false
	}

	day := date.Day()

	return day == s.DayOfMonth
}

func (s *MonthlyStrategy) NextOccurrence(from time.Time) time.Time {
	year := from.Year()
	month := from.Month()

	next := time.Date(year, month, s.DayOfMonth, 0, 0, 0, 0, from.Location())

	if !next.After(from) {
		next = next.AddDate(0, 1, 0)
	}

	return next
}

func (s *MonthlyStrategy) Description() string {
	return fmt.Sprintf("Каждое %d-е число месяца", s.DayOfMonth)
}

func (s *MonthlyStrategy) Validate() error {
	if s.DayOfMonth < 1 || s.DayOfMonth > 30 {
		return fmt.Errorf("day_of_month must be between 1 and 30")
	}
	return nil
}
