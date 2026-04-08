package frequency

import (
	"time"
)

type Frequency struct {
	Type         FrequencyType `json:"type"`
	IntervalDays *int          `json:"interval_days,omitempty"`
	DayOfMonth   *int          `json:"day_of_month,omitempty"`
	Dates        []int         `json:"dates,omitempty"`
	Parity       *ParityType   `json:"parity,omitempty"`
}

func (f *Frequency) Validate() error {
	if f == nil {
		return nil
	}

	switch f.Type {
	case FrequencyDaily:
		if f.IntervalDays == nil || *f.IntervalDays < 1 {
			return ErrInvalidFrequency
		}

	case FrequencyMonthly:
		if f.DayOfMonth == nil || *f.DayOfMonth < 1 || *f.DayOfMonth > 30 {
			return ErrInvalidFrequency
		}

	case FrequencySpecific:
		if len(f.Dates) == 0 {
			return ErrInvalidFrequency
		}
		for _, d := range f.Dates {
			if d < 1 || d > 31 {
				return ErrInvalidFrequency
			}
		}

	case FrequencyParity:
		if f.Parity == nil || (*f.Parity != ParityEven && *f.Parity != ParityOdd) {
			return ErrInvalidFrequency
		}

	default:
		return ErrInvalidFrequency
	}

	return nil
}

func (f *Frequency) ToStrategy() (FrequencyStrategy, error) {
	if f == nil {
		return nil, nil
	}

	switch f.Type {
	case FrequencyDaily:
		return &DailyStrategy{IntervalDays: *f.IntervalDays}, nil
	case FrequencyMonthly:
		return &MonthlyStrategy{DayOfMonth: *f.DayOfMonth}, nil
	case FrequencySpecific:
		return &SpecificStrategy{Dates: f.Dates}, nil
	case FrequencyParity:
		return &ParityStrategy{Parity: *f.Parity}, nil
	default:
		return nil, ErrInvalidFrequency
	}
}

func (f *Frequency) IsValid(date time.Time) (bool, error) {
	if f == nil {
		return false, nil
	}

	strategy, err := f.ToStrategy()
	if err != nil {
		return false, err
	}

	return strategy.IsValid(date), nil
}

func (f *Frequency) NextOccurrence(from time.Time) (time.Time, error) {
	if f == nil {
		return from, nil
	}

	strategy, err := f.ToStrategy()
	if err != nil {
		return from, err
	}

	return strategy.NextOccurrence(from), nil
}

func (f *Frequency) Description() string {
	if f == nil {
		return "Без периодичности"
	}

	strategy, err := f.ToStrategy()
	if err != nil {
		return "Неизвестная периодичность"
	}

	return strategy.Description()
}
