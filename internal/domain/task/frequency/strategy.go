package frequency

import "time"

type FrequencyStrategy interface {
	IsValid(date time.Time) bool
	NextOccurrence(from time.Time) time.Time
	Description() string
	Validate() error
}
