package flag

import "flag"

var (
	T             string // 使用的cache类型,memory
	RedisAddr     string
	RedisUserName string
	RedisPasswd   string
)

func init() {
	flag.StringVar(&T, "type", "memory", "Select the type of cache to use \n memory \n redis")
	flag.StringVar(&RedisAddr, "redisAddr", "127.0.0.1:6379", "The address of redis")
	flag.StringVar(&RedisUserName, "redisUserName", "", "The UserName of redis")
	flag.StringVar(&RedisPasswd, "redisPasswd", "", "The Passwd of redis")
	flag.Parse()
}
