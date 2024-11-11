package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type workState struct {
	wg        sync.WaitGroup
	maxJobs   int
	maxErrors int
	mu        sync.Mutex
	jobInWork int
	cntErrors int
}

func (w *workState) IncreaseJobs() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.jobInWork++
}

func (w *workState) ReduceJobs() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.jobInWork--
}

func (w *workState) IncreaseErrors() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.cntErrors++
}

func (w *workState) ReduceErrors() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.cntErrors--
}

func (w *workState) JobsInWork() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.jobInWork
}

func (w *workState) CntErrors() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.cntErrors
}

func (w *workState) IsMaxErrorsRecieved() bool {
	// если указанно максимальное кол-во ошибок 0 или меньше, игнорируем кол-во ошибок
	if w.maxErrors <= 0 {
		return false
	}
	cntErrors := w.CntErrors()
	return cntErrors == w.maxErrors || cntErrors >= w.maxErrors
}

func (w *workState) IsMaxJobsRunnig() bool {
	return w.JobsInWork() == w.maxJobs
}

type Task func() error

func Run(tasks []Task, n, m int) error {
	wg := sync.WaitGroup{}
	tasksCh := make(chan Task, n)
	quitCh := make(chan struct{})
	state := workState{
		maxJobs:   n,
		maxErrors: m,
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		Consume(tasksCh, quitCh, &state)
	}()

	for _, task := range tasks {
		if state.IsMaxErrorsRecieved() {
			quitCh <- struct{}{}
			return ErrErrorsLimitExceeded
		}
		tasksCh <- task
	}
	close(tasksCh)
	wg.Wait()
	return nil
}

func Consume(in <-chan Task, quitCh <-chan struct{}, state *workState) {
	for task := range in {
		select {
		case <-quitCh:
			return
		default:
		}

		// проверяем кол-во заупущенных тасок
		if state.IsMaxJobsRunnig() {
			state.wg.Wait()
		}

		state.wg.Add(1)
		state.IncreaseJobs()
		go func() {
			defer state.wg.Done()
			defer state.ReduceJobs()
			if state.IsMaxErrorsRecieved() {
				return
			}
			err := task()
			if err != nil {
				state.IncreaseErrors()
			}
		}()
	}
	state.wg.Wait()
}
