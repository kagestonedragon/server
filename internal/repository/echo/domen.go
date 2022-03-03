package echo

import "fmt"

type echo struct {}

type Echo interface {
	init()
}

func NewEcho() Echo {
	return &echo{}
}

func (e *echo) init() {
	fmt.Println("init echo service")
}
