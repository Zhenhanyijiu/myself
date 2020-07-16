package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	cn := redis.NewClient(&redis.Options{Addr: "192.168.5.213:6379", Password: "123456", DB: 0})
	str, err := cn.Ping().Result()
	if err != nil {
		fmt.Printf("=======err:%v,string:%v\n", err, str)
		return
	}
	defer cn.Close()
	fmt.Printf("=======string:%v\n", str)
	///////////////////////////
	opt, err := redis.ParseURL("redis://192.168.5.213:6379/0")
	if err != nil {
		panic(err)
	}
	fmt.Printf("=======opt:%+v\n", opt)
	rdb := redis.NewClient(opt)
	str, err = rdb.Ping().Result()
	if err != nil {
		fmt.Printf("=======err:%v,string:%v\n", err, str)
		return
	}
	defer rdb.Close()
	fmt.Printf("=======string:%v\n", str)

}
