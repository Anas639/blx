package task

import (
	"testing"
)

func TestStartNew(t *testing.T) {
	task := NewTask(0, "test")
	_, error := task.Start()
	if error != nil {
		t.Errorf("Expected the task to start without errors, but instead it ended wit error %s", error.Error())
		return
	}
	ongoing := task.Status == TASK_ONGOING
	if !ongoing {
		t.Errorf("Expected the task status to be %s, but instead got %s", TASK_ONGOING, task.Status)
		return
	}
	sessionLength := len(task.sessions)
	if sessionLength != 1 {
		t.Errorf("Expected the task sessions to have a length of %d, but instead got a length of %d", 1, sessionLength)
		return
	}
}

func TestStartOngoing(t *testing.T) {
	task := NewTask(0, "test")
	task.SetStatus(TASK_ONGOING)
	session := CreateTaskSession(0)
	task.sessions = append(task.sessions, session)
	_, error := task.Start()
	if error == nil {
		t.Errorf("Expected the task to throw an error because it's already started, but instead it started normaly ")
		return
	}
}

func TestStartPaused(t *testing.T) {
	task := NewTask(0, "test")
	task.SetStatus(TASK_PAUSED)
	session := CreateTaskSession(0)
	session.End()
	task.sessions = append(task.sessions, session)
	_, error := task.Start()
	if error != nil {
		t.Errorf("Expected the task to start without errors, but instead it ended wit error %s", error.Error())
		return
	}
	ongoing := task.Status == TASK_ONGOING
	if !ongoing {
		t.Errorf("Expected the task status to be %s, but instead got %s", TASK_ONGOING, task.Status)
		return
	}
	sessionLength := len(task.sessions)
	if sessionLength != 2 {
		t.Errorf("Expected the task sessions to have a length of %d, but instead got a length of %d", 2, sessionLength)
		return
	}
}

func TestStartEnded(t *testing.T) {
	task := NewTask(0, "test")
	task.SetStatus(TASK_ENDED)
	_, error := task.Start()
	if error == nil {
		t.Errorf("Expected the task to throw an error, but instad it started without any error!")
		return
	}
}

func TestPauseNew(t *testing.T) {
	task := NewTask(0, "test")
	task.SetStatus(TASK_NEW)
	_, err := task.Pause()
	if err == nil {
		t.Errorf("Expected the task to throw an error, but instad it paused without any error!")
		return
	}
}

func TestPauseOnGoing(t *testing.T) {
	task := NewTask(0, "test")
	task.SetStatus(TASK_ONGOING)
	session := CreateTaskSession(0)
	task.sessions = append(task.sessions, session)
	_, err := task.Pause()
	if err != nil {
		t.Errorf("Expected the task to pause normally, but instead it ended with errors %s", err.Error())
		return
	}
	if task.Status != TASK_PAUSED {
		t.Errorf("Expected the task status to be %s, but intead got %s", TASK_PAUSED, task.Status)
		return
	}
	sessionLength := len(task.sessions)
	if sessionLength != 1 {
		t.Errorf("Expected the task sessions to have a length of %d, but instead got a length of %d", 1, sessionLength)
		return
	}
	if session.isOngoing() {
		t.Errorf("Expected the last session to be ended but instead it was still ongoing")
		return
	}
}

func TestPausePaused(t *testing.T) {
	task := NewTask(0, "test")
	task.SetStatus(TASK_PAUSED)
	session := CreateTaskSession(0)
	session.End()
	task.sessions = append(task.sessions, session)
	_, err := task.Pause()
	if err == nil {
		t.Errorf("Expected the task not to pause because it's already paused, but instad it paused normally")
		return
	}
}

func TestPauseEnded(t *testing.T) {
	task := NewTask(0, "test")
	task.SetStatus(TASK_ENDED)
	session := CreateTaskSession(0)
	session.End()
	task.sessions = append(task.sessions, session)
	_, err := task.Pause()
	if err == nil {
		t.Errorf("Expected the task not to pause because it's already ended, but instad it paused normally")
		return
	}
}

func TestEndNew(t *testing.T) {
	task := NewTask(0, "test")
	task.SetStatus(TASK_NEW)
	_, error := task.End()
	if error == nil {
		t.Errorf("Expected the task to end with an error because it does not contain any sessions yet")
		return
	}
}

func TestEndOngoing(t *testing.T) {
	task := NewTask(0, "test")
	task.SetStatus(TASK_ONGOING)
	session := CreateTaskSession(0)
	task.sessions = append(task.sessions, session)
	_, error := task.End()
	if error != nil {
		t.Errorf("Expected the task to end successfully, but instead it ended with error %s", error.Error())
		return
	}
	if task.Status != TASK_ENDED {
		t.Errorf("Expected the task status to be %s, but intead got %s", TASK_ENDED, task.Status)
		return
	}
	sessionLength := len(task.sessions)
	if sessionLength != 1 {
		t.Errorf("Expected the task sessions to have a length of %d, but instead got a length of %d", 1, sessionLength)
		return
	}
	if session.isOngoing() {
		t.Errorf("Expected the last session to be ended but instead it was still ongoing")
		return
	}
}

func TestEndPaused(t *testing.T) {
	task := NewTask(0, "test")
	task.SetStatus(TASK_PAUSED)
	session := CreateTaskSession(0)
	session.End()
	task.sessions = append(task.sessions, session)
	_, error := task.End()
	if error != nil {
		t.Errorf("Expected the task to end successfully, but instead it ended with error %s", error.Error())
		return
	}
	if task.Status != TASK_ENDED {
		t.Errorf("Expected the task status to be %s, but intead got %s", TASK_ENDED, task.Status)
		return
	}
	sessionLength := len(task.sessions)
	if sessionLength != 1 {
		t.Errorf("Expected the task sessions to have a length of %d, but instead got a length of %d", 1, sessionLength)
		return
	}
	if session.isOngoing() {
		t.Errorf("Expected the last session to be ended but instead it was still ongoing")
		return
	}
}

func TestEndEnded(t *testing.T) {
	task := NewTask(0, "test")
	task.SetStatus(TASK_ENDED)
	session := CreateTaskSession(0)
	session.End()
	task.sessions = append(task.sessions, session)
	_, error := task.End()
	if error == nil {
		t.Errorf("Expected the task to end with errors because it's already ended, but instad it ende normally")
		return
	}
}
