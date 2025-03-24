package task

import (
	"fmt"
	"time"
)

type TaskSession struct {
	Id       int64
	StartsAt time.Time
	TaskId   int64
	EndsAt   time.Time
}

func CreateTaskSession(taskId int64) *TaskSession {
	return &TaskSession{
		StartsAt: time.Now(),
		EndsAt:   time.Time{},
		TaskId:   taskId,
	}
}
func NewTaskSession(id int64, startTime time.Time, endTime time.Time, taskId int64) *TaskSession {
	var sTime time.Time
	if startTime.IsZero() {
		sTime = time.Now()
	} else {
		sTime = startTime
	}

	return &TaskSession{
		Id:       id,
		StartsAt: sTime,
		EndsAt:   endTime,
		TaskId:   taskId,
	}
}

func (this *TaskSession) isOngoing() bool {
	return this.EndsAt.IsZero()
}
func (this *TaskSession) End() {
	if this.isOngoing() {
		this.EndsAt = time.Now()
	}
}

func (this *TaskSession) String() string {
	endTime := "-"
	if !this.isOngoing() {
		endTime = this.EndsAt.Format(time.RFC850)
	}
	return fmt.Sprintf("#%d {%s, %s} => %d", this.Id, this.StartsAt.Format(time.RFC850), endTime, this.TaskId)
}

func (this *TaskSession) Duration() time.Duration {
	if this.EndsAt.IsZero() {
		return time.Duration(0) * time.Second
	}
	d := this.EndsAt.Sub(this.StartsAt)
	return d
}
