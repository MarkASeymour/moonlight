package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var send = primeSend
var trim = strings.TrimSpace
var low = strings.ToLower

func uci(fromGUI chan string) {
	send("info string Inside UCI")
	fromEngine, toEngine := engine()
	quit := false
	cmd := ""
	words := []string{}
	bm := ""
	for quit == false {
		select {
		case cmd = <-fromGUI:
			words = strings.Split(cmd, " ")
		case bm = <-fromEngine:
			handleBestMove(bm)
			continue

		}
		words[0] = trim(low(words[0]))
		switch words[0] {
		case "uci":
			handleUci()
		case "isready":
			handleIsReady()
		case "setoption":
			handleSetOption(words)
		case "stop":
			handleStop(toEngine)
		case "quit", "q":
			quit = true
			continue

		}
	}
}

func handleSetOption(option []string) {
	send("info string set option ", strings.Join(option, " "))
	send("info string not yet implemented")
}

func handleIsReady() {
	send("readyok")
}

func handleUci() {
	send("id name Moonlight")
	send("id author Markinho")
	send("option name Hash type spin default 32 min 1 max 1024")
	send("option name Threads type spin default 1 min 1 max 16")
	send("uciok")

}
func handleBestMove(bm string) {
	send(bm)
}
func handleStop(toEng chan string) {
	toEng <- "stop"
}

func input() chan string {
	line := make(chan string)
	go func() {
		var reader *bufio.Reader
		reader = bufio.NewReader(os.Stdin)
		for {
			text, err := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			if err != io.EOF && len(text) > 0 {
				line <- text
			}
		}
	}()
	return line
}
func primeSend(text ...string) {
	toGUI := ""
	for _, t := range text {
		toGUI += t
	}
	fmt.Println(toGUI)
}
