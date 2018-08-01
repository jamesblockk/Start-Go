package bigline

import (
	"fmt"
	"time"

	"github.com/golang-collections/go-datastructures/queue"
)

func (line *BigLine) Start() {
	go line.checkSubExpire()
	go line.checkMasterExpire()
}

func (line *BigLine) checkSubExpire() {
	for {

		if line.Sub.Queue.Len() >= 1 {
			line.mergeToMaster(line.Sub.Queue, "mergeToMaster")
		}

		time.Sleep(line.SubExpireTime)
	}
}

func (line *BigLine) checkMasterExpire() {
	for {
		line.pop()
		time.Sleep(line.PopFrequency)
	}
}

func (line *BigLine) Push(items ...interface{}) {
	line.Sub.Queue.Put(items)

	if line.Sub.Queue.Len() > line.MaxSubQueue {
		line.mergeToMaster(line.Sub.Queue, "mergeToMaster")
	}
}

func (line *BigLine) mergeToMaster(q *queue.Queue, comment string) {

	fmt.Println("go")
	data := make(chan interface{})
	isLoad := make(chan bool)

	go func() {
		for {
			select {
			case msg := <-data:
				line.already(msg)
				fmt.Println("OK", comment)
			}
			<-isLoad
		}
	}()

	if q.Len() >= 1 {
		b, _ := q.Get(line.MaxSubQueue)
		data <- b
		isLoad <- true
		return
	}

	fmt.Println(" The Queue length == 0 Or waiting time < second ", comment)
}

func (line *BigLine) already(in interface{}) {
	line.Master.Put(in)
}

func (line *BigLine) pop() {
	if line.Master.Len() < 1 {
		fmt.Println("BigLine.Len() < 1 ")
		return
	}

	m, err := line.Master.Get(1)

	if _, ok := line.Delegate.(BigLineProtocol); ok {
		line.Delegate.BigLinePop(m, err)
	}

	fmt.Println("bigQ", m)
	// line.Time = time.Now()
}
