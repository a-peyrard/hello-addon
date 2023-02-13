package main

import (
	"fmt"
	"github.com/a-peyrard/addon-manager/process"
)

type Process struct{}

func (p *Process) Run() (err error) {
	fmt.Printf(`
*********************
*   hello v1.0.0!   *
*********************
`)

	return
}

func NewProcess() process.Process {
	return &Process{}
}
