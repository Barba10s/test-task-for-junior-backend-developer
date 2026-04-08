package frequency

import "errors"

type FrequencyType string

const (
	FrequencyDaily    FrequencyType = "daily"
	FrequencyMonthly  FrequencyType = "monthly"
	FrequencySpecific FrequencyType = "specific"
	FrequencyParity   FrequencyType = "parity"
)

var ErrInvalidFrequency = errors.New("invalid frequency settings")
