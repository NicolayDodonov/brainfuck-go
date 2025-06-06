package main

import "log"

type program struct {
	code []rune
	len  int
}

func main() {
	p, err := loadProgram()
	if err != nil {
		return
	}
	proc := newProcessor()

	if err := proc.run(p); err != nil {
		log.Printf(err.Error())
	}
}

func loadProgram() (*program, error) {
	return nil, nil
}
