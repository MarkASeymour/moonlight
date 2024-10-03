package main

import "fmt"

func engine() (fromEngine, toEngine chan string) {
	fmt.Println("info string Engine Initialize")
	fromEngine = make(chan string)
	toEngine = make(chan string)
	go func() {
		for cmd := range toEngine {
			//primeSend("info string engine rcvd ", cmd)
			switch cmd {
			case "stop":
			case "quit":
			}
		}
	}()
	return fromEngine, toEngine
}
