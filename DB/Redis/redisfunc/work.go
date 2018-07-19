package redisfunc

import (
	"fmt"
	"sync"
	"time"
)

type Redis struct {
}

// pipe
func (S *Redis) pipe(com string, urls ...chan [][]interface{}) chan interface{} {
	Pipe := make(chan interface{})
	fmt.Println("go! Pipe " + com)
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		time.Sleep(1 * time.Millisecond)
		c := conn.Get() //// 從這裡 取得連線 盡量避免重複
		// if _, err := c.Do("SELECT", 1); err != nil { /// 選資料庫  　不建議使用
		// 	c.Close()
		// }
		go func(url chan [][]interface{}, conn interface{}) {
			//fmt.Println(url, "url")
			m := 0
			for test := range url {
				m = len(test)
				for _, value := range test {
					err := c.Send(com, value...)
					if err != nil {
						fmt.Println("redis set failed:", err)
					}
				}
			}
			if err := c.Flush(); err != nil { // 清空記憶體　發送
				fmt.Println("err :", err)
			}
			for i := 1; i <= m; i++ {
				re, err := c.Receive()
				if err != nil {
					//log ..s
					fmt.Println(com+"err", i, err)
				}
				Pipe <- re
			}
			c.Close()
			wg.Done()
		}(url, c)
	}

	go func() {
		wg.Wait()
		fmt.Println("Done !")
		close(Pipe)
	}()

	return Pipe
}

// scan 模糊搜尋
func (S *Redis) scanlimit(com string, value []interface{}) chan interface{} {
	fmt.Println("go! Pipe " + com)
	Pipe := make(chan interface{})
	var b string
	var wg sync.WaitGroup
	c := conn.Get() //// 從這裡 取得連線 盡量避免重複
	wg.Add(1)
	go func() {
		fmt.Println(value)
		c.Send(com, value...)
		if err := c.Flush(); err != nil { // 清空記憶體　發送
			fmt.Println("err :", err)
		}
		re, err := c.Receive()
		fmt.Println("SCAN:", fmt.Sprintf("%s", re), "SCAN  Done!")
		if err != nil {
			//log ..s
			fmt.Println(com+"err", err)
		}
		b = string((re.([]interface{})[0]).([]byte))
		Pipe <- re
		for b != "0" {
			c.Send(com, b, value[1], value[2], value[3], value[4])
			if err := c.Flush(); err != nil { // 清空記憶體　發送
				fmt.Println("err :", err)
			}
			re, err := c.Receive()
			if err != nil {
				//log ..s
				fmt.Println(com+"err", err)
			}
			b = string((re.([]interface{})[0]).([]byte))
			Pipe <- re
		}
		c.Close()
		wg.Done()
	}()
	go func() {
		wg.Wait()
		fmt.Println("Done !")
		close(Pipe)
	}()
	return Pipe
}

// 單獨執行
func (S *Redis) do(com string, urls chan []interface{}) chan interface{} {
	fmt.Println("go! DO " + com)
	DO := make(chan interface{})
	c := conn.Get()
	res, err := c.Do(com, <-urls...)
	if err != nil {
		fmt.Println(com+"err", err)
	}
	DO <- res
	return DO
}

func (S *Redis) mergeValue(com string, table string, key []string, val []interface{}) ([]interface{}, bool) {

	var urls []interface{}

	// table , key ,key
	if com == "HMGET" {
		//fmt.Println(" table , key ,key")
		urls = append(urls, table)
		for _, i := range key {
			urls = append(urls, i)
		}
		return urls, true
	}
	//  table ,key ,value , key ,value
	if com == "HMSET" {
		//	fmt.Println(" table ,key ,value , key ,value")
		for b, i := range key {
			urls = append(urls, i)
			urls = append(urls, val[b])
		}
		return urls, true
	}
	//table (key)  VALUE1  VALUE1
	if com == "SADD" || com == "SREM" {
		//fmt.Println(" key  VALUE1  VALUE1 ")
		urls = append(urls, table)
		for _, i := range val {
			urls = append(urls, i)
		}
		return urls, true

	}
	//table key value or  table key1 value1 key2 value2 ....
	if com == "HSET" {
		//fmt.Println("table key value")
		urls = append(urls, table)
		for i := 0; i <= len(key)-1; i++ {
			urls = append(urls, key[i])
			urls = append(urls, val[i])
		}
		return urls, true
	}
	//  table key
	if com == "HGET" || com == "HEXISTS" || com == "HDEL" {
		//fmt.Println("table key")
		urls = append(urls, table)
		urls = append(urls, key[0])

		return urls, true
	}
	// table
	if com == "HGETALL" || com == "SMEMBERS" || com == "DEL" {
		//	fmt.Println("table ")
		urls = append(urls, table)

		return urls, true

	}
	if com == "SCAN" {

		urls = append(urls, "0")
		urls = append(urls, "MATCH")
		urls = append(urls, "*"+table+"*") // maybe *value and value* or *value*
		urls = append(urls, "COUNT")
		urls = append(urls, "10000")
		return urls, true
	}

	return nil, false

}

//// 轉換通道
func (S *Redis) buildChan(value [][]interface{}) chan [][]interface{} {
	var lo = make(chan [][]interface{})
	go func() {
		lo <- value
		close(lo)
	}()
	return lo
}

//// 單程轉換通道
func (S *Redis) signbuildChan(value []interface{}) chan []interface{} {
	var lo = make(chan []interface{})
	go func() {
		lo <- value
		close(lo)
	}()
	return lo
}
