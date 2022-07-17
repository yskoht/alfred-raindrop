package main

import (
	"errors"
	"fmt"
	"os"

	raindropSync "github.com/yskoht/alfred-raindrop/internal/raindrop-sync"
)

func getToken() (string, error) {
	testToken := os.Getenv("testToken")
	if testToken == "" {
		return "", errors.New("testToken not found")
	}

	token := fmt.Sprintf("Bearer %s", testToken)
	return token, nil
}

func run() {
	token, err := getToken()
	if err != nil {
		fmt.Print(err)
		return
	}

	err = raindropSync.RaindropSync(token)
	if err != nil {
		fmt.Print(err)
		return
	}
}

func main() {
	run()
}
