package main

import (
	"fmt"
	"github.com/Joker666/AsyncGoDemo/async"
	"time"
)

func DoneAsync() int {
	fmt.Println("Warming up ...")
	time.Sleep(3 * time.Second)
	fmt.Println("Done ...")
	return 1
}

func main() {
	fmt.Println("Let's start ...")
	future := async.Exec(func() interface{} {
		return DoneAsync()
	})
	fmt.Println("Done is running ...")
	val := future.Await()
	fmt.Println(val)
}
