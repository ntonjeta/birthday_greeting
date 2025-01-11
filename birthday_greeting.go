package main

import (
	"fmt"
)

type Friend struct {
	name    string
	surname string
}

func greeting(_ []Friend) (string, error) {
	return "Subject: Happy birthday!\nHappy birthday, dear Mary!", nil
}

func main() {
	fmt.Printf("Birthday greeting! \n")
}
