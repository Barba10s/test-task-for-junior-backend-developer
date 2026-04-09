package task

import "errors"

var (
	ErrInvalidInput = errors.New("invalid task input")
	ErrTimeConflict = errors.New("time slot conflict")
)
