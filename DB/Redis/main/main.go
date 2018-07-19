package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/jamesblockk/Start-Go/DB/Redis/redisfunc"
)

var STD redisfunc.RDInModel

func init() {
	hash := redisfunc.Hash{
		HDEL:    "HDEL",
		HEXISTS: "HEXISTS",
		HGET:    "HGET",
		HGETALL: "HGETALL",
		HSET:    "HSET",
	}

	Set := redisfunc.Set{
		SADD:      "SADD",
		SCARD:     "SCARD",
		SDIFF:     "SDIFF",
		SINTER:    "SINTER",
		SISMEMBER: "SISMEMBER",
		SSCAN:     "SSCAN",
		SMEMBERS:  "SMEMBERS",
	}

	Other := redisfunc.Other{
		SCAN: "SCAN",
	}
	key := redisfunc.Key{DEL: "DEL"}

	STD = redisfunc.RDInModel{Hash: hash, Set: Set, Key: key, Other: Other}
}

func main() {

	redisfunc.Testconn()
	test()
	test2()

}

// scan  test
func test() {
	var re = redisfunc.Redisfunc{}
	STD.Table = "100"

	c, _ := re.Sign(STD.Other.SCAN, STD.Table, nil, nil)

	t1 := time.Now()

	chan1 := re.Scanlimit(STD.Other.SCAN, c)

	for i := range chan1 {
		//fmt.Println(chan1)

		fmt.Println(i)
	}
	fmt.Println("ok")
	elapsed := time.Since(t1)    //   測試運時間用的
	fmt.Println("time", elapsed) //印出時間

}

// pipe  test
func test2() {
	var re = redisfunc.Redisfunc{}
	var TEST []chan [][]interface{}
	var urlss [][]interface{}
	////////////////////////
	for i := 1; i <= 100000; i++ {
		//	g := make(map[string][]interface{})
		var Field []string
		var Value []interface{}
		STD.Table = "100" + strconv.Itoa(i)
		STD.Field = "100" + strconv.Itoa(i)
		STD.Value = GetGuid(0)
		Value = append(Value, STD.Value)
		Field = append(Field, STD.Field)

		c, _ := re.Sign(STD.Hash.HDEL, STD.Table, Field, Value)
		urlss = append(urlss, c)
		//	fmt.Println(urlss)
	}
	a := re.Chan(urlss)
	TEST = append(TEST, a) //queue?
	///////////////////////////////////////////////////// 測試資料

	t1 := time.Now()
	chan1 := re.Pipe(STD.Hash.HDEL, TEST...)
	//chan1 := re.Pipe(STD.Hash.HSET, TEST...)
	for i := range chan1 {
		//fmt.Println(chan1)
		fmt.Println(i) //回傳值;　判斷何種狀態
	}
	fmt.Println("ok")
	elapsed := time.Since(t1)    // 測試運時間用的
	fmt.Println("time", elapsed) //印出時間

}

func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
func GetGuid(typeID int) string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return getMd5String(base64.URLEncoding.EncodeToString(b))
}
