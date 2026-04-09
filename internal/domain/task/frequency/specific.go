package frequency

import (
	"fmt"
	"slices"
	"sort"
	"strings"
	"time"
)

type SpecificStrategy struct {
	Dates []int `json:"dates"`
}

func (s *SpecificStrategy) IsValid(date time.Time) bool {
	if len(s.Dates) == 0 {
		return false
	}

	day := date.Day()

	return slices.Contains(s.Dates, day)
}

func (s *SpecificStrategy) NextOccurrence(from time.Time) time.Time {
	if len(s.Dates) == 0 {
		return from
	}

	sorted := make([]int, len(s.Dates))
	copy(sorted, s.Dates)
	sort.Ints(sorted)

	currentDay := from.Day()
	currentMonth := from.Month()
	currentYear := from.Year()

	for _, day := range sorted {
		if day > currentDay {
			return time.Date(currentYear, currentMonth, day, 0, 0, 0, 0, from.Location())
		}
	}

	nextMonth := currentMonth + 1
	nextYear := currentYear
	if nextMonth > 12 {
		nextMonth = 1
		nextYear++
	}

	return time.Date(nextYear, nextMonth, sorted[0], 0, 0, 0, 0, from.Location())
}

func (s *SpecificStrategy) Description() string {
	if len(s.Dates) == 0 {
		return "Без периодичности"
	}

	sorted := make([]int, len(s.Dates))
	copy(sorted, s.Dates)
	sort.Ints(sorted)

	parts := make([]string, len(sorted))
	for i, d := range sorted {
		parts[i] = fmt.Sprintf("%d", d)
	}

	return fmt.Sprintf("Числа месяца: %s", strings.Join(parts, ", "))
}

func (s *SpecificStrategy) Validate() error {
	if len(s.Dates) == 0 {
		return fmt.Errorf("dates array cannot be empty")
	}
	if len(s.Dates) > 31 {
		return fmt.Errorf("maximum 31 dates allowed (got %d)", len(s.Dates))
	}

	seen := make(map[int]bool)
	for _, day := range s.Dates {
		if day < 1 || day > 31 {
			return fmt.Errorf("each date must be between 1 and 31 (got %d)", day)
		}
		if seen[day] {
			return fmt.Errorf("duplicate date: %d", day)
		}
		seen[day] = true
	}
	return nil
}
