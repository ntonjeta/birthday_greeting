package main

import (
	"fmt"
)

type GreetingSender interface {
	send(name string) error
}

type ConsoleGreetingSender struct {
}

func (g *ConsoleGreetingSender) send(name string) error {
	fmt.Printf("Subject: Happy birthday!\nHappy birthday, dear %s!", name)
	return nil
}
