package main

import (
	"net/http"
	"time"

	"github.com/jamesblockk/Start-Go/HelloGo/newqueue/control"
)

func main() {

	controlTest := control.Shared()
	controlTest.Init()
	controlTest.Start()

	go func() {
		for i := 0; i <= 1000; i++ {
			controlTest.Push(i)
		}

	}()

	go func() {
		for i := 0; i <= 1000; i++ {
			controlTest.Push(i)
		}

	}()

	go func() {
		for i := 0; i <= 1000; i++ {
			controlTest.Push(i)
		}

	}()

	time.Sleep(1 * time.Second)
	http.ListenAndServe(":8080", nil)
}
