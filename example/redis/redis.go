package main

import (
	"fmt"
	"github.com/ea67/gocore/redis"
	"github.com/ea67/gocore/viper"
)

func main() {
	viper.NewConfig("config", "conf")

	redis.GetRedisOptions("email_check")
	redis.GetRedisDB("email_check").Set("test", "sunmi")
	fmt.Println(redis.GetRedisDB("encryption").Get("test").String())
}
