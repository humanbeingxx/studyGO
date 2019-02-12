package runner

import (
	"errors"
	"os"
	"os/signal"
	"sync"
	"time"
)

var ErrTimeout = errors.New("timeout")
var ErrInterrupt = errors.New("interupted")

// Runner is a container for concurrent task with timeout
type Runner struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time

	tasks []func()
}

//New makes a runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

//AddTask add task
func (r *Runner) AddTask(tasks ...func()) {
	r.tasks = append(r.tasks, tasks...)
}

//Start runner all tasks
func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	var runnerWaitGroup sync.WaitGroup

	runnerWaitGroup.Add(len(r.tasks))

	for _, task := range r.tasks {
		if r.interrupted() {
			return ErrInterrupt
		}
		taskForRun := task
		go func() {
			defer runnerWaitGroup.Done()
			taskForRun()
		}()
	}

	runnerWaitGroup.Wait()
	return nil
}

func (r *Runner) interrupted() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
