package frequency

import (
	"fmt"
	"time"
)

type Frequency struct {
	Type         FrequencyType `json:"type"`
	IntervalDays *int          `json:"interval_days,omitempty"`
	DayOfMonth   *int          `json:"day_of_month,omitempty"`
	Dates        []int         `json:"dates,omitempty"`
	Parity       *ParityType   `json:"parity,omitempty"`
	TimeOfDay    *string       `json:"time_of_day,omitempty"`
}

func (f *Frequency) Validate() error {
	if f == nil {
		return nil
	}

	if err := f.ValidateTime(); err != nil {
		return err
	}

	strategy, err := f.ToStrategy()
	if err != nil {
		return err
	}
	if strategy == nil {
		return ErrInvalidFrequency
	}
	return strategy.Validate()
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

func isValidTimeFormat(t string) bool {
	if t == "" {
		return true
	}
	if len(t) != 5 || t[2] != ':' {
		return false
	}
	for i, c := range t {
		if i == 2 {
			continue
		}
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func isValidTimeValue(t string) bool {
	hours := int(t[0]-'0')*10 + int(t[1]-'0')
	minutes := int(t[3]-'0')*10 + int(t[4]-'0')
	return hours >= 0 && hours <= 23 && minutes >= 0 && minutes <= 59
}

func (f *Frequency) ValidateTime() error {
	if f.TimeOfDay == nil {
		return nil
	}

	t := *f.TimeOfDay

	if !isValidTimeFormat(t) {
		return fmt.Errorf("time_of_day must be in HH:MM format (got '%s')", t)
	}

	if !isValidTimeValue(t) {
		return fmt.Errorf("time_of_day must be 00-23:00-59 (got '%s')", t)
	}

	return nil
}
