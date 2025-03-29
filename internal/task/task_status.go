package task

type TaskStatus = string

const (
	TASK_NEW     TaskStatus = "new"
	TASK_ONGOING TaskStatus = "ongoing"
	TASK_PAUSED  TaskStatus = "paused"
	TASK_ENDED   TaskStatus = "ended"
)

func StatusesFromSlice(input []string) []TaskStatus {
	values := map[string]TaskStatus{
		"new":     TASK_NEW,
		"ongoing": TASK_ONGOING,
		"paused":  TASK_PAUSED,
		"ended":   TASK_ENDED,
	}

	existingStatuses := make(map[TaskStatus]bool, 4)
	statuses := []TaskStatus{}
	for _, status := range input {
		st, ok := values[status]
		if !ok {
			continue
		}
		exists := existingStatuses[st]
		if exists {
			continue
		}
		existingStatuses[st] = true
		statuses = append(statuses, st)
	}
	return statuses
}

func AllStatuses() []TaskStatus {
	return []TaskStatus{
		TASK_NEW, TASK_ONGOING, TASK_PAUSED, TASK_ENDED,
	}
}
