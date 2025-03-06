package schedule

import (
	"sync"
	"sync/atomic"
	"time"
)

// Task holds information about the running task and can be used to stop running tasks.
type Task struct {
	stop          chan struct{}
	nextExecution time.Time
	startedAt     time.Time
	stopped       int32 // 0 means active, 1 means stopped
	once          sync.Once
}

// newTask creates a new Task.
func newTask() *Task {
	return &Task{
		stop:      make(chan struct{}),
		startedAt: time.Now(),
	}
}

// StartedAt returns the time when the scheduler was started.
func (s *Task) StartedAt() time.Time {
	return s.startedAt
}

// NextExecutionTime returns the time when the next execution will happen.
func (s *Task) NextExecutionTime() time.Time {
	return s.nextExecution
}

// ExecutesIn returns the duration until the next execution.
func (s *Task) ExecutesIn() time.Duration {
	return time.Until(s.nextExecution)
}

// IsActive returns true if the scheduler is active.
func (s *Task) IsActive() bool {
	return atomic.LoadInt32(&s.stopped) == 0
}

// Wait blocks until the scheduler is stopped.
// After and At will stop automatically after the task is executed.
func (s *Task) Wait() {
	<-s.stop
}

// Stop stops the scheduler.
func (s *Task) Stop() {
	s.once.Do(func() {
		atomic.StoreInt32(&s.stopped, 1)
		close(s.stop)
	})
}

// After executes the task after the given duration.
// The function is non-blocking. If you want to wait for the task to be executed, use the Task.Wait method.
func After(duration time.Duration, task func()) *Task {
	scheduler := newTask()
	scheduler.nextExecution = time.Now().Add(duration)
	timer := time.NewTimer(duration)

	go func() {
		select {
		case <-timer.C:
			task()
			scheduler.Stop()
		case <-scheduler.stop:
			// If the task is stopped before the timer fires, stop the timer.
			if !timer.Stop() {
				<-timer.C // drain if necessary
			}
			return
		}
	}()

	return scheduler
}

// At executes the task at the given time.
// The function is non-blocking. If you want to wait for the task to be executed, use the Task.Wait method.
func At(t time.Time, task func()) *Task {
	scheduler := newTask()
	scheduler.nextExecution = t
	d := time.Until(t)
	if d < 0 {
		d = 0
	}
	timer := time.NewTimer(d)

	go func() {
		select {
		case <-timer.C:
			task()
			scheduler.Stop()
		case <-scheduler.stop:
			if !timer.Stop() {
				<-timer.C
			}
			return
		}
	}()

	return scheduler
}

// Every executes the task in the given interval, as long as the task function returns true.
// The function is non-blocking. If you want to wait for the task to be executed, use the Task.Wait method.
func Every(interval time.Duration, task func() bool) *Task {
	scheduler := newTask()
	scheduler.nextExecution = time.Now().Add(interval)
	ticker := time.NewTicker(interval)

	go func() {
		for {
			select {
			case <-ticker.C:
				if !task() {
					scheduler.Stop()
					ticker.Stop()
					return
				}
				scheduler.nextExecution = time.Now().Add(interval)
			case <-scheduler.stop:
				ticker.Stop()
				return
			}
		}
	}()

	return scheduler
}
