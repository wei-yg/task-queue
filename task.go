package task_queue

import (
	"errors"
	"fmt"
	"time"
)

type Task func() error

type Worker struct {
	Exit  chan bool
	Tasks chan Task
}

func NewWorker() *Worker {
	return NewNoBlockingWorker(0)
}

func NewNoBlockingWorker(channelCount int) *Worker {
	exit := make(chan bool)
	tasks := make(chan Task, channelCount)
	worker := &Worker{
		Exit:  exit,
		Tasks: tasks,
	}
	go worker.run()
	return worker
}

func (w *Worker) run() {
	for {
		select {
		case <-w.Exit:
			return
		case task := <-w.Tasks:
			if err := task(); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
func (w *Worker) Push(task Task) {
	w.Tasks <- task
}
func (w *Worker) Stop() {
	w.Exit <- true
}
func (w *Worker) Sleep(d time.Duration) error {
	timeout := time.After(d)
	for {
		select {
		case <-w.Exit:
			return errors.New("主动退出")
		case <-timeout:
			return nil
		}
	}
}
