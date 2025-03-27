package tui

import (
	"fmt"
	"io"
	"os"
	"time"
)

type PrintMode int8

const (
	PRINTMODE_NEWLINE PrintMode = iota
	PRINTMODE_SINGLELINE
)

type TimeTracker interface {
	SetWriter(writer io.Writer)
	Start() chan struct{}
	Stop()
	SetPrintMode(mode PrintMode)
}

func NewTrackerFromElapsed(elapsed float64) TimeTracker {
	return &timeTracker{
		elapsedSeconds: elapsed,
		writer:         os.Stdout,
		isRunning:      make(chan struct{}),
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
	isRunning      chan struct{}
	mode           PrintMode
}

func (this *timeTracker) SetPrintMode(mode PrintMode) {
	this.mode = mode
}
func (this *timeTracker) SetWriter(writer io.Writer) {
	this.writer = writer
}
func (this *timeTracker) Start() chan struct{} {
	go this.run()
	return this.isRunning
}

func (this *timeTracker) Stop() {
	close(this.isRunning)
}

func (this *timeTracker) run() {
	format := "\r\033[K%s"
	if this.mode == PRINTMODE_NEWLINE {
		format = "%s\n"
	}
	for {
		ticker := time.Tick(time.Second)
		fmt.Printf(format, time.Duration(this.elapsedSeconds)*time.Second)
		select {
		case <-ticker:
			{
				this.elapsedSeconds++
			}
		case _, ok := <-this.isRunning:
			{
				if !ok {
					return
				}
			}
		}
	}
}
