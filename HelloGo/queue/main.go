package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/jamesblockk/Start-Go/HelloGo/queue/queue"
)

// reference https://blog.dubbelboer.com/2015/04/25/go-faster-queue.html

func main() {
	t := time.Tick(time.Second)
	q := queue.New()

	for i := 1; i > 0; i++ {
		q.Push(i)

		if q.Len() > 10 {
			q.Pop()
		}

		timeCount(t, q)
	}

}

func timeCount(t <-chan time.Time, q *queue.Squeue) {
	select {
	case <-t:
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("cap: %d, len: %d, used: %d\n", q.Cap(), q.Len(), m.Alloc)
	default:
	}
}
