package greeting

type GreetingSender interface {
	Send(name string) error
}
