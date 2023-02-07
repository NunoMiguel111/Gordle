package main

import (
	"bufio"
	"gordle/gordle"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	g := gordle.New(r)
	g.Play()
}
