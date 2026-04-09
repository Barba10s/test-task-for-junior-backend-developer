package frequency

import (
	"fmt"
	"time"
)

type ParityType string

const (
	ParityEven ParityType = "even"
	ParityOdd  ParityType = "odd"
)

type ParityStrategy struct {
	Parity ParityType `json:"parity"`
}

func (s *ParityStrategy) IsValid(date time.Time) bool {
	if s.Parity != ParityEven && s.Parity != ParityOdd {
		return false
	}

	day := date.Day()

	isEven := day%2 == 0

	if s.Parity == ParityEven {
		return isEven
	}
	return !isEven
}

func (s *ParityStrategy) NextOccurrence(from time.Time) time.Time {
	next := from.AddDate(0, 0, 1)
	next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())

	for i := 0; i < 31; i++ {
		if s.IsValid(next) {
			return next
		}
		next = next.AddDate(0, 0, 1)
	}

	return from
}

func (s *ParityStrategy) Description() string {
	switch s.Parity {
	case ParityEven:
		return "По чётным дням месяца"
	case ParityOdd:
		return "По нечётным дням месяца"
	default:
		return "Неизвестная периодичность"
	}
}

func (s *ParityStrategy) Validate() error {
	if s.Parity != ParityEven && s.Parity != ParityOdd {
		return fmt.Errorf("parity must be 'even' or 'odd'")
	}
	return nil
}
