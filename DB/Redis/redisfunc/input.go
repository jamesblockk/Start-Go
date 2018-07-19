package redisfunc

type RDInModel struct {
	DB    string
	Host  string
	Table string
	Field string
	Value interface{}
	RDState
	Key   Key
	Hash  Hash
	Set   Set
	List  List
	Other Other
}

////
type RDState struct {
	Err   error
	Stats string
}

type Hash struct {
	HDEL    string //HDEL key field1 [field2] 删除一个或多个哈希表字段
	HEXISTS string // HEXISTS key field  查看哈希表 key 中，指定的字段是否存在
	HGET    string // 	HGET key field   获取存储在哈希表中指定字段的值。
	HGETALL string //	HGETALL key  获取在哈希表中指定 key 的所有字段和值
	HSET    string // 	HLEN key  获取哈希表中字段的数量
}
type List struct {
}

type Set struct {
	SADD      string //SADD key member1 [member2]  向集合添加一个或多个成员
	SCARD     string //	SCARD key  获取集合的成员数
	SDIFF     string //	SDIFF key1 [key2]  返回给定所有集合的差集
	SINTER    string //SINTER key1 [key2]  返回给定所有集合的交集
	SISMEMBER string // SISMEMBER key member  判断 member 元素是否是集合 key 的成员
	SSCAN     string // SSCAN key cursor [MATCH pattern] [COUNT count]  迭代集合中的元素
	SMEMBERS  string
}

type Key struct {
	DEL string //SADD key member1 [member2]  向集合添加一个或多个成员
}
type Other struct {
	SCAN string //SCAN 
}
