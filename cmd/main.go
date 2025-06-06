package main

import (
	"brainfuck-go/processor"
	"fmt"
)

func main() {
	//Open and close handler
	hello()
	defer exit()

	//load program
	program, err := processor.Load()
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		return
	}

	//start interpretation
	machine := processor.New()
	if err := machine.Run(program); err != nil {
		fmt.Println("ERROR: ", err.Error())
	}
}

func hello() {
	fmt.Print("Hello User!\n" +
		"It's a Brainfuck interpolator v1.0.0\n" +
		"Make by Dodonov N.A.\n\n" +
		"Please, entre path to code: ")
}

func exit() {
	fmt.Print("\nPress enter to quit...")
	var exit int
	fmt.Scanln(&exit)
}
