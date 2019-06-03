package pool

import (
	"sync"
)

type Task struct {
	Proc func()
}

type worker struct {
	jobQueue chan Task
	workerPool chan chan Task
}


type Dispatcher struct {
	workerPool chan chan Task
	workerSize int
	once sync.Once
}


func newWorker(p chan chan Task) *worker {
	return &worker{jobQueue : make(chan Task), workerPool: p}
}


func (w *worker) run() {
	for {
		w.workerPool <- w.jobQueue
		select {
			case task := <- w.jobQueue:
				task.Proc()
		}
	}
}

func NewDispatcher(workerSize int) *Dispatcher {
	return &Dispatcher{workerSize:workerSize}
}

func (d *Dispatcher) Run() {
	d.once.Do(func () {
		workerPool := make(chan chan Task, d.workerSize)
		d.workerPool = workerPool
		for i := 0; i < d.workerSize; i++ {
			worker := newWorker(d.workerPool)
			go worker.run()
		}
	})
}

func (d *Dispatcher) Schedule(t Task) {
	w := <- d.workerPool
	w <- t
}
