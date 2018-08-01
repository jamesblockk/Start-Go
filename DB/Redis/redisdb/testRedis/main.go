package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/jamesblockk/Start-Go/DB/Redis/redisdb"
)

func main() {
	if err := redisdb.Shared().Init(); err != nil {
		log.Fatal(err)
		return
	}

	type data struct {
		S  string            `json:"s"`
		Aa map[string]string `json:"d"`
	}

	m := map[string]string{}
	m["asda"] = "dvvv"

	d := data{S: "dsfsd ", Aa: m}
	b, er := json.Marshal(d)
	if er != nil {
		fmt.Print(er)
	}
	fmt.Println(string(b))

	rd := redisdb.Conn()
	ExecTime(100000, func() {
		// if err := redisdb.Set("sdd", b); err != nil {
		// 	fmt.Print(err)
		// }
		// rd.Incr("zxccc")
		rd.Get("sdd")

	})

	rd.Close()

	// aaa, _ := redisdb.Get("sdd")
	// fmt.Println(string(aaa))

}

func ExecTime(count int, funcT func()) {
	t1 := time.Now() //   測試運行時間用的
	i := 0
	for i < count {
		funcT()
		i++
	}

	elapsed := time.Since(t1)    //   測試運時間用的
	fmt.Println("time", elapsed) //印出時間
}
