package main

import (
	"errors"
	"fmt"
	"time"
)

type Friend struct {
	name    string
	surname string
}

func greeting(_ []Friend, _ time.Time) (string, error) {
	return "", errors.New("todo")
}

func main() {
	fmt.Printf("Birthday greeting! \n")
}
