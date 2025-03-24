package task

type TaskStatus = string

const (
	TASK_NEW     TaskStatus = "new"
	TASK_ONGOING TaskStatus = "ongoing"
	TASK_PAUSED  TaskStatus = "paused"
	TASK_ENDED   TaskStatus = "ended"
)
