package main

import (
	"fmt"
	"time"

	"github.com/jamesblockk/Start-Go/HelloGo/iteratorGO/iteractor"
)

type DDD struct {
	A int
}

func main() {

	defaultForLoop()

	// var a ddd

	// i := 0
	var datas []interface{}
	i := 0

	for i < 100 {
		aa := DDD{A: i}
		datas = append(datas, aa)
		i++
	}

	t1 := time.Now() //   測試運行時間用的
	iteractor.Loop(datas)

	elapsed := time.Since(t1)    //   測試運時間用的
	fmt.Println("time", elapsed) //印出時間
}

func defaultForLoop() {
	t1 := time.Now() //   測試運行時間用的
	i := 0
	var datas []interface{}

	for i < 100 {
		aa := DDD{A: i}
		datas = append(datas, aa)
		i++
	}

	elapsed := time.Since(t1)    //   測試運時間用的
	fmt.Println("time", elapsed) //印出時間
}

func iteractorLoop() {
	// t1 := time.Now() //   測試運行時間用的
	// i := 0

	// elapsed := time.Since(t1)    //   測試運時間用的
	// fmt.Println("time", elapsed) //印出時間
}
