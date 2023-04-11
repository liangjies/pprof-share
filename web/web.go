package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go http.ListenAndServe("0.0.0.0:6061", nil)
	for {
		go func() {
			for {
				fmt.Println("hello world")
			}
		}()
		time.Sleep(200 * time.Microsecond)
	}
}
