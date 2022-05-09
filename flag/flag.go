package flag

import (
	"flag"
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	T             string // 使用的cache类型,memory
	RedisAddr     string
	RedisUserName string
	RedisPasswd   string
)

func init() {
	flag.StringVar(&T, "type", "memory", "Select the type of cache to use \n- memory \n- redis\n")
	flag.StringVar(&RedisAddr, "redisAddr", "127.0.0.1:6379", "The address of redis")
	flag.StringVar(&RedisUserName, "redisUserName", "", "The UserName of redis")
	flag.StringVar(&RedisPasswd, "redisPasswd", "", "The Passwd of redis")
	flag.Parse()
	fmt.Println("name", RedisUserName, "name")
	// 读取conf.ini配置文件
	cfg, err := ini.Load("conf.ini")
	if err == nil {
		T = cfg.Section("").Key("type").String()
		redis := cfg.Section("redis")
		RedisAddr = redis.Key("addr").String()
		RedisUserName = redis.Key("userName").String()
		if RedisUserName == "" {

		}
		fmt.Println("name", RedisUserName, "name")
		RedisPasswd = redis.Key("passwd").String()
		fmt.Println("passwd", RedisPasswd, "passwd")
	}

}
