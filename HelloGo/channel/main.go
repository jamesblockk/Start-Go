package main

import (
	"fmt"
	"time"
)

func main() {

	execTime(10000, func() {
		fmt.Println("2")
	})
}

func execTime(count int, funcT func()) {
	t1 := time.Now() //   測試運行時間用的
	i := 0
	for i < count {
		funcT()
		i++
	}

	elapsed := time.Since(t1)    //   測試運時間用的
	fmt.Println("time", elapsed) //印出時間
}
