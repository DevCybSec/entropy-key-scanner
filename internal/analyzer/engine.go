package analyzer

import (
	"fmt"
	"os"
	"sync"
)

type Engine struct {
	WorkerCount int
	Jobs        chan Job
	Results     chan Result
	Wait        *sync.WaitGroup
}

func NewEngine(workerCount int) *Engine {
	return &Engine{
		WorkerCount: workerCount,
		Jobs:        make(chan Job, 100),
		Results:     make(chan Result, 100),
		Wait:        &sync.WaitGroup{},
	}
}

// Start lanza los workers
func (e *Engine) Start() {
	for i := 0; i < e.WorkerCount; i++ {
		e.Wait.Add(1)
		go e.worker(i)
	}
}

// worker es el consumidor que procesará los archivos
func (e *Engine) worker(id int) {
	defer e.Wait.Done()
	fmt.Printf("[Worker %d] Iniciado\n", id)

	for job := range e.Jobs {

		e.processJob(job)

		e.Results <- Result{
			FilePath: job.Path,
			Findings: []string{}, // Lógica pendiente
		}
	}
}

func (e *Engine) processJob(job Job) {
	f, err := os.Open(job.Path)
	if err != nil {
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			return
		}
	}(f)

	// Solo leemos el primer MB para optimizar RAM y velocidad
	const maxRead = 1 * 1024 * 1024
	data := make([]byte, maxRead)
	n, _ := f.Read(data)

	// Solo analizamos lo que realmente se leyó
	e.analyzeFile(job.Path, data[:n])
}
