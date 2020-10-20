package hello

import "fmt"

type Greeter struct {
	Template string
}

type GreeterInput struct {
	Name string
}

type GreeterOutput struct {
	Message string
}

type GreetError struct {
}

func (g *Greeter) Greet(conf *GreeterInput) (*GreeterOutput, error) {
	return &GreeterOutput{
		Message: fmt.Sprintf(g.Template, conf.Name),
	}, nil
}