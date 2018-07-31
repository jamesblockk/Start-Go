package redisdb

type RDConnEntry struct {
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
