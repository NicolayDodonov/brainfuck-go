package processor

import (
	"bytes"
	"errors"
	"fmt"
	"os"
)

// command in brainfuck
const (
	right     = '>'
	left      = '<'
	plus      = '+'
	minus     = '-'
	openLoop  = '['
	closeLoop = ']'
	print     = '.'
	read      = ','
)

// processor is interpreted brainfuck-program.
type Processor struct {
	line           [30000]byte
	programPointer int
}

type Program struct {
	code []rune
	len  int
}

// make new Processor
func New() *Processor {
	return &Processor{
		[30000]byte{},
		0,
	}
}

func (p *Processor) up() {
	p.programPointer++
	if p.programPointer >= 30000 {
		p.programPointer -= 30000
	}
}

func (p *Processor) down() {
	p.programPointer--
	if p.programPointer < 0 {
		p.programPointer = 30000 - p.programPointer
	}
}

// run - processor function to interpretation brainfuck-program.
func (p *Processor) Run(prog *Program) error {
	for pointer := 0; pointer < prog.len; pointer++ {
		switch prog.code[pointer] {
		case right:
			p.up()
		case left:
			p.down()
		case plus:
			p.line[p.programPointer]++
		case minus:
			p.line[p.programPointer]--
		case openLoop:
			if p.line[p.programPointer] == 0 {
				loop := 1
				for loop != 0 {
					p.up()
					if p.line[p.programPointer] == openLoop {
						loop++
					}
					if p.line[p.programPointer] == closeLoop {
						loop--
					}
				}
			}
		case closeLoop:
			if p.line[p.programPointer] != 0 {
				loop := 1
				for loop != 0 {
					p.down()
					if p.line[p.programPointer] == openLoop {
						loop--
					}
					if p.line[p.programPointer] == closeLoop {
						loop++
					}
				}
			}
		case print:
			fmt.Print(p.line[p.programPointer])
		case read:
			var value byte
			_, err := fmt.Scanln(value)
			if err != nil {
				return err
			}
			p.line[p.programPointer] = value
		}
	}
	return nil
}

func Load() (*Program, error) {
	// read path from console
	var userPath string
	_, err := fmt.Scanln(userPath)

	if err != nil {
		return nil, err
	}

	// verification query
	if len(userPath) == 0 {
		return nil, errors.New("path is empty")
	}
	_, err = os.Stat(userPath)
	if err != nil {
		return nil, err
	}

	//read file
	rawProgram, err := os.ReadFile(userPath)
	if err != nil {
		return nil, err
	}
	p := &Program{
		code: bytes.Runes(rawProgram),
	}
	p.len = len(p.code)

	//clear program

	return p, nil
}
