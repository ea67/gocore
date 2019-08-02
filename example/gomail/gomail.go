package main

import (
	"github.com/ea67/gocore/gomail"
	"github.com/ea67/gocore/viper"
)

func main() {

	viper.NewConfig("config", "conf")

	gomail.SendEmail("wenzhenxi@vip.qq.com", "service@sunmi.com", "service", "title", "msg")
}
