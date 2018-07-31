package redisdb

import (
	// "IMBOX/imrd/configs"

	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

type singleton struct {
	// wsList [n]*list.List
	Pool *redis.Pool
	// Co   *RDCode
}

var instance *singleton
var once sync.Once

// var cLocks [n]sync.RWMutex

func Shared() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}

func (red *singleton) InitRedis() error {
	// 47.75.171.233 -p 6379 -a Papple168169
	var cf RDConnEntry
	//cf.ProxyAddress = "192.168.6.85:6379"
	cf.ProxyAddress = "127.0.0.1:6379"
	// cf.PassWord = configs.RD().PWD
	//cf.ProxyAddress = "127.0.0.1:6379"

	//連線ＴＹＰＥ , 最大閒置連線　，　最大連線
	cf.ConnType = "tcp"
	cf.MaxIdle = 120000
	cf.MaxActive = 2400000
	/// 連線　　讀取　寫入　的逾時
	cf.ConnectTimeout = 10000
	cf.ReadTimeout = 1000
	cf.WriteTimeout = 1000
	////0
	//// 選擇ＤＢ　０,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15
	cf.DBnumber = 0
	// fmt.Println("cf", cf)
	red.Pool = newPool(cf) /////設定連線

	err := red.Pool.TestOnBorrow(red.Pool.Get(), time.Now())
	if err != nil {
		return err
	}
	// red.Co = &RDCode{
	// 	///////// 設定　　代碼
	// 	DEL:          "DEL",
	// 	Err:          "err",
	// 	Exists:       "EXIST",
	// 	Existsnofund: "Existsnofund",
	// 	Faile:        "Falise",
	// 	SendDEL:      "SendDEL",
	// 	CreateDone:   "HSET Create Done!",
	// 	UPdateDone:   "HSET UPdate Done!"}

	return nil
}

func RDConn() *redis.Pool {
	return Shared().Pool
}

func newPool(cf RDConnEntry) *redis.Pool {
	return &redis.Pool{
		MaxIdle:   cf.MaxIdle,
		MaxActive: cf.MaxActive, // max number of connections

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(cf.ConnType, cf.ProxyAddress,
				redis.DialConnectTimeout(time.Duration(cf.ConnectTimeout)*time.Millisecond),
				redis.DialReadTimeout(time.Duration(cf.ReadTimeout)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(cf.WriteTimeout)*time.Millisecond),
			)
			if err != nil {
				panic(err.Error())
			}
			if _, err := c.Do("AUTH", cf.PassWord); err != nil { ///密碼　驗證
				c.Close()
			}
			if _, err := c.Do("SELECT", cf.DBnumber); err != nil { /// 選資料庫
				c.Close()
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error { //檢查連線狀態
			if time.Since(t) > time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},

		IdleTimeout: 3 * time.Second, // 閒置連線逾時
		Wait:        true,
	}

}
