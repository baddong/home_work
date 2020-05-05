package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks
func Run(tasks []Task, n int, m int) error {
	var errCount int
	taskCh := make(chan Task, len(tasks))
	wg := &sync.WaitGroup{}
	wg.Add(n)
	mu := &sync.Mutex{}

	for i := 0; i < n; i++ {
		go func(taskCh <-chan Task, mu *sync.Mutex, errCount *int) {
			defer wg.Done()
			stop := false
			for t := range taskCh {
				err := t()
				mu.Lock()
				if *errCount >= m {
					stop = true
				}
				if err != nil {
					*errCount++
				}
				mu.Unlock()
				if stop {
					return
				}
			}
		}(taskCh, mu, &errCount)
	}

	for _, t := range tasks {
		taskCh <- t
	}

	close(taskCh)
	wg.Wait()

	if errCount >= m {
		return ErrErrorsLimitExceeded
	}
	return nil
}
