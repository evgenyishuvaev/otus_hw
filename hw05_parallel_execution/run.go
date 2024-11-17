package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type workState struct {
	wg        sync.WaitGroup
	maxErrors int
	mu        sync.Mutex
	cntErrors int
}

func (w *workState) IncreaseErrors() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.cntErrors++
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

type Task func() error

func Run(tasks []Task, n, m int) error {
	state := workState{
		maxErrors: m,
	}
	defer state.wg.Wait()

	tasksCh := make(chan int)
	defer close(tasksCh)

	worker := func() {
		state.wg.Add(1)
		defer state.wg.Done()

		for indx := range tasksCh {
			err := tasks[indx]()
			if err != nil {
				state.IncreaseErrors()
			}
		}
	}

	for i := 0; i < n; i++ {
		go worker()
	}

	for indx := range tasks {
		if state.IsMaxErrorsRecieved() {
			return ErrErrorsLimitExceeded
		}
		tasksCh <- indx
	}
	return nil
}
