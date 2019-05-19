package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	ch := make(chan string, 0)
	defer func() {
		close(ch)
	}()

	go func() {
	DONE:
		for {
			time.Sleep(1 * time.Second)
			fmt.Println(time.Now().Unix())
			i++

			select {
			case m := <-ch:
				println(m)
				break DONE
			default:
			}
		}
	}()

	time.Sleep(time.Second * 4)
	ch <- "stop"
}
