/*
Package pool provides utilities for managing parallel execution of tasks.
*/
package pool

// Pool is a structure that maintains pool of goroutines and spreads
// incoming tasks over them.
type Pool struct {
	work chan func()
	sem  chan struct{}
}

// NewPool creates new Pool that manages up to n goroutines.
func NewPool(n int) *Pool {
	return &Pool{
		sem:  make(chan struct{}, n),
		work: make(chan func()),
	}
}

// Exec schedules task to be executed on one of started goroutines.
// If goroutines number is not at the limit, it may spawn new one.
func (p *Pool) Exec(task func()) {
	select {
	// Try to schedule task first if there are some idle workers.
	case p.work <- task:
	default:
		// If there are no idle workers, we may only start new one (if
		// possible), or wait for some worker become idle.
		select {
		case p.sem <- struct{}{}:
			go p.worker(task)
		case p.work <- task:
		}
	}
}

// Close closes all underlying workers and returns when all of them are dead.
func (p *Pool) Close() {
	close(p.work)
	// Wait for all workers are done.
	for i := 0; i < cap(p.sem); i++ {
		p.sem <- struct{}{}
	}
}

func (p *Pool) worker(task func()) {
	defer func() {
		// Pull one item from semaphore to indicate that worker is done.
		<-p.sem
	}()
	var ok bool
	for {
		task()
		if task, ok = <-p.work; !ok {
			return
		}
	}
}
