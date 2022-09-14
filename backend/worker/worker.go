package worker

import (
	"fmt"
	"sync"

)

var workerInstance *WorkerPool
var workerMux = new(sync.Mutex)


type WorkerPool struct {
	JobChan chan func()
	// ResultChan chan bool
}
type workers interface {
	worker()
	CreateWorker()
}

func WorkerSingleton() *WorkerPool {
    if workerInstance == nil {
        workerMux.Lock()
        if workerInstance == nil {
            workerInstance = &WorkerPool{
                JobChan: make(chan func(), 10000),
                // ResultChan: make(chan bool, 100),
            }
            defer workerMux.Unlock()
            return  workerInstance
        }
    }
    return workerInstance
}

func(t *WorkerPool) worker(id int) {
    fmt.Println("compare", *t, &t)
			
    for job := range (*t).JobChan {
        
        fmt.Println("worker", id, "started  job")
        job()//do task
        fmt.Println("worker", id, "finished job")
    }
    
    
}
func(t *WorkerPool) CreateWorker(workerAmount int) {
    for w := 1; w <= workerAmount; w++ {
        
        go (*t).worker(w)
    }
}
// func ReadWithSelect(ch chan int) (x int, err error) {
//     timeout := time.NewTimer(time.Microsecond * 500)

//     select {
//     case x = <-ch:
//         return x, nil
//     case <-timeout.C:
//         return 0, errors.New("write time out")
//     }
// }