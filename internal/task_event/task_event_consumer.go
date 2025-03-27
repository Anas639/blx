package taskevent

import (
	"fmt"

	"github.com/anas639/blx/internal/event"
	"github.com/anas639/blx/internal/services"
	"github.com/anas639/blx/internal/task"
	"github.com/anas639/blx/internal/tui"
)

type TaskEventConsumer interface {
	Start()
	Wait()
}

func NewTaskEventConsumer(taskService *services.TaskService, payload chan event.EventPayload) TaskEventConsumer {
	return &eventConsumer{
		payload:     payload,
		done:        make(chan struct{}),
		taskService: taskService,
	}
}

type eventConsumer struct {
	payload     chan event.EventPayload
	done        chan struct{}
	taskService *services.TaskService
	timer       tui.TimeTracker
}

func (this *eventConsumer) stopTimer() {
	if this.timer != nil {
		this.timer.Stop()
		this.timer = nil
	}
}

func (this *eventConsumer) checkActiveTask() error {
	activeTask, err := this.taskService.GetLastActiveTask()
	if err != nil {
		return err
	}

	this.stopTimer()

	this.timer = tui.NewTrackerFromElapsed(activeTask.GetElapsedTime(task.TIMER_MODE_TASK).Seconds())
	this.timer.SetPrintMode(tui.PRINTMODE_NEWLINE)
	this.timer.Start()
	return nil
}
func (this *eventConsumer) Start() {
	go func() {
		err := this.checkActiveTask()
		if err != nil {
			fmt.Println("No Task")
		}
		for {
			pl, ok := <-this.payload
			if !ok {
				break
			}
			if pl.Err() != nil {
				fmt.Println(pl.Err().Error())
				continue
			}
			switch pl.Type {
			case event.EVENT_START:
				{
					this.checkActiveTask()
				}
			case event.EVENT_PAUSE:
				{
					this.stopTimer()
					err := this.checkActiveTask()
					if err != nil {
						fmt.Println("Task Paused")
					}
				}
			case event.EVENT_END:
				{
					this.stopTimer()
					err := this.checkActiveTask()
					if err != nil {
						fmt.Println("Task Ended")
					}
				}

			}
		}
		close(this.done)
	}()
}

func (this *eventConsumer) Wait() {
	<-this.done
}
