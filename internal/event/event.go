package event

import "fmt"

const (
	EVENT_START byte = iota
	EVENT_PAUSE
	EVENT_END
)

type EventPayload struct {
	Type   byte
	TaskId int64
	err    error
}

func NewPayload(t byte, id int64) EventPayload {
	return EventPayload{
		TaskId: id,
		Type:   t,
		err:    nil,
	}
}

func NewPayloadError(err error) EventPayload {
	return EventPayload{
		TaskId: 0,
		Type:   0,
		err:    err,
	}
}

func (this *EventPayload) String() string {
	if this.Err() != nil {
		return this.Err().Error()
	}
	var typeStr string
	switch this.Type {
	case EVENT_START:
		{
			typeStr = "start"
		}
	case EVENT_PAUSE:
		{
			typeStr = "pause"
		}
	case EVENT_END:
		{
			typeStr = "end"
		}
	default:
		{
			typeStr = "unkown"

		}
	}
	return fmt.Sprintf("%s task #%d", typeStr, this.TaskId)
}

func (this *EventPayload) Err() error {
	return this.err
}

type EventBroadcaster interface {
	SendEvent(event EventPayload) error
}

type EventListener interface {
	Listen() (chan EventPayload, error)
	Close() error
}
