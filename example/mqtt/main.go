package main

import (
	"fmt"
	"github.com/ea67/gocore/mqtt"
)

func main() {
	fmt.Println(mqtt.QOS_0)
	fmt.Println(mqtt.QOS_1)
	fmt.Println(mqtt.QOS_2)
}
