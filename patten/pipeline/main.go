package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jamesblockk/start-go/patten/pipeline/pipe"
)

// Lesson From https://segmentfault.com/a/1190000014788594

func loadCheckpoint() int {
	return 1
}

func extractReviewsFromA(check *int, num int) ([]string, error) {
	// fmt.Println("extractReviewsFromA")

	return []string{"3453", "1123"}, nil
}

func jobA(datas []string) error {
	// fmt.Println("jobA", datas)
	return nil
}

func jobB(datas []string) error {
	// fmt.Println("jobB", datas)

	return nil
}

func jobC(datas []string) error {
	// fmt.Println("jobC", datas)

	return nil
}

func saveCheckpoint(checkpoint int) error {
	// fmt.Println("saveCheckpoint", checkpoint)
	return nil
}

func main() {
	t1 := time.Now() //   測試運行時間用的

	i := 0
	for i < 100000 {
		startPipeLine()
		i++
	}
	elapsed := time.Since(t1)    //   測試運時間用的
	fmt.Println("time", elapsed) //印出時間
}

func startPipeLine() {
	checkpoint := loadCheckpoint()

	//工序(1)在pipeline外执行，最后一个工序是保存checkpoint
	pipeline := pp.NewPipeline(8, 32, 2, 1)
	for {
		//(1)
		//加载100条数据，并修改变量checkpoint
		//data是数组，每个元素是一条评论，之后的联表、NLP都直接修改data里的每条记录。
		data, err := extractReviewsFromA(&checkpoint, 100)
		if err != nil {
			log.Print(err)
			break
		}
		curCheckpoint := checkpoint

		ok := pipeline.Async(func() error {
			//(2)
			return jobA(data)
		}, func() error {
			//(3)
			return jobB(data)
		}, func() error {
			//(4)
			return jobC(data)
		}, func() error {
			//(5)保存checkpoint
			// log.Print("done:", curCheckpoint)
			return saveCheckpoint(curCheckpoint)
		})

		if !ok {
			break
		}

		if len(data) < 100 {
			break
		} //done

		// fmt.Println(len(data))
	}
	err := pipeline.Wait()
	if err != nil {
		log.Print(err)
	}
}
