package main

import (
	"github.com/fzzy/radix/redis"

	"fmt"
	"time"
)

const ( //Change these parameters to your redis server parms
	network = "tcp"
	addr    = "127.0.0.1:6379"
)
const (
	set = "set"
	get = "get"
)
const (
	number = 1000 // change this to the number of tests you want to run
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {

	c, err := redis.Dial(network, addr) // client connects to server
	checkErr(err)

	startTime := time.Now()
	for i := 0; i < 1000; i++ {
		c.Cmd(set, i)
		c.Cmd(get)

	}

	endTime := time.Now()
	deltatime := endTime.Sub(startTime)

	fmt.Println(deltatime)

}
