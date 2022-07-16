package main

import (
	"fmt"
	"github.com/deanishe/awgo"
	"github.com/yskoht/alfred-raindrop/internal/raindrop-sync"
	"os"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func run() {
	testToken := os.Getenv("testToken")
	token := fmt.Sprintf("Bearer %s", testToken)

	err := raindropSync.RaindropSync(token)
	if err != nil {
		return
	}
}

func main() {
	wf.Run(run)
}
