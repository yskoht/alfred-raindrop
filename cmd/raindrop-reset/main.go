package main

import (
	"fmt"

	raindropReset "github.com/yskoht/alfred-raindrop/internal/raindrop-reset"
)

func run() {
	err := raindropReset.RaindropReset()
	if err != nil {
		fmt.Print(err)
		return
	}
}

func main() {
	run()
}
