package tui

import (
	"fmt"
	"io"
	"os"
	"time"
)

type TimeTracker interface {
	SetWriter(writer io.Writer)
	Start()
	Stop()
}

func NewTrackerFromElapsed(elapsed float64) TimeTracker {
	return &timeTracker{
		elapsedSeconds: elapsed,
		writer:         os.Stdout,
		isRunning:      make(chan bool),
	}
}

func NewTrackerFromTime(from time.Time) TimeTracker {
	now := time.Now()
	elapsedSeconds := now.Sub(from).Seconds()
	return NewTrackerFromElapsed(elapsedSeconds)
}

type timeTracker struct {
	elapsedSeconds float64
	writer         io.Writer
	isRunning      chan bool
}

func (this *timeTracker) SetWriter(writer io.Writer) {
	this.writer = writer
}
func (this *timeTracker) Start() {
	go this.run()
	<-this.isRunning
}

func (this *timeTracker) Stop() {
	this.isRunning <- false
}

func (this *timeTracker) run() {
	for {
		ticker := time.Tick(time.Second)
		fmt.Printf("\r\033[K%s", time.Duration(this.elapsedSeconds)*time.Second)
		select {
		case <-ticker:
			{
				this.elapsedSeconds++
			}
		case r := <-this.isRunning:
			{
				if !r {
					return
				}
			}
		}
	}
}
