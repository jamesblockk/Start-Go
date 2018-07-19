package redisfunc

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var conn *redis.Pool
var Co Code
func NewPool(cf Connstruct) *redis.Pool {
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
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},

		IdleTimeout: 1 * time.Millisecond, // 閒置連線逾時
		Wait:        true,
	}

}

type Code struct {
	DEL          string
	Faile        string
	SendDEL      string
	Exists       string
	Existsnofund string
	errr         string
	CreateDone   string
	UPdateDone   string
}

func Testconn() {

	var Cf Connstruct
	//連線ＴＹＰＥ , 最大閒置連線　，　最大連線
	Cf.ProxyAddress = "127.0.0.1:6379"
	Cf.ConnType = "tcp"
	Cf.PassWord = "12345678"

	Cf.MaxIdle = 120000
	Cf.MaxActive = 2400000
	/// 連線　　讀取　寫入　的逾時
	Cf.ConnectTimeout = 600
	Cf.ReadTimeout = 600
	Cf.WriteTimeout = 600
	////
	//// 選擇ＤＢ　０,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15
	Cf.DBnumber = 0
	//	fmt.Println("cf", Cf)
	conn = NewPool(Cf) /////設定連線
	/////////// 設定　　代碼
	Co.DEL = "DEL"
	Co.errr = "err"
	Co.Exists = "EXIST"
	Co.Existsnofund = "Existsnofund"
	Co.Faile = "Falise"
	Co.SendDEL = "SendDEL"
	Co.CreateDone = "HSET Create Done!"
	Co.UPdateDone = "HSET UPdate Done!"
}

type Connstruct struct {
	ProxyAddress   string
	PassWord       string
	MaxIdle        int
	MaxActive      int
	ConnType       string
	ConnectTimeout int
	ReadTimeout    int
	WriteTimeout   int
	DBnumber       int
}

type Userid struct {
	Userid string
	Token
}

type Token struct {
	Token string
	Host  string
}

type Insider struct {
	Userid string
	Mybyte []byte
}
