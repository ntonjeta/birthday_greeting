package sender

import fmt 

type ConsoleGreetingSender struct {
}

func (g *ConsoleGreetingSender) Send(name string) error {
	fmt.Printf("Subject: Happy birthday!\nHappy birthday, dear %s!", name)
	return nil
}
