package handlers

import (
	"time"

	taskdomain "example.com/taskservice/internal/domain/task"
	taskfrequency "example.com/taskservice/internal/domain/task/frequency"
)

type taskMutationDTO struct {
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
	Status      taskdomain.Status        `json:"status"`
	Frequency   *taskfrequency.Frequency `json:"frequency,omitempty"`
	TimeOfDay   *string                  `json:"time_of_day,omitempty"`
}

type taskDTO struct {
	ID                   int64                    `json:"id"`
	Title                string                   `json:"title"`
	Description          string                   `json:"description"`
	Status               taskdomain.Status        `json:"status"`
	Frequency            *taskfrequency.Frequency `json:"frequency,omitempty"`
	FrequencyDescription string
	TimeOfDay            *string   `json:"time_of_day,omitempty"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

func newTaskDTO(task *taskdomain.Task) taskDTO {
	dto := taskDTO{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		Frequency:   task.Frequency,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}

	if task.Frequency != nil {
		dto.FrequencyDescription = task.Frequency.Description()
		dto.TimeOfDay = task.Frequency.TimeOfDay
	}

	return dto
}
