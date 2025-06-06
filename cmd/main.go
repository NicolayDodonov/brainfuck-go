package main

import (
	"brainfuck-go/processor"
	"fmt"
	"log"
)

func main() {

	hello()

	p, err := processor.Load()
	if err != nil {
		return
	}
	proc := processor.New()

	if err := proc.Run(p); err != nil {
		log.Printf(err.Error())
	}

	fmt.Scanln()
}

func hello() {
	fmt.Println("Hello User!")
}
