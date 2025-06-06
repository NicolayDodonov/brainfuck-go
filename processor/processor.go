package processor

import (
	"bytes"
	"errors"
	"fmt"
	"os"
)

// command in brainfuck
const (
	right        = '>'
	left         = '<'
	plus         = '+'
	minus        = '-'
	openLoop     = '['
	closeLoop    = ']'
	print        = '.'
	read         = ','
	memoryLength = 30000
)

// Processor is interpreted brainfuck-program.
type Processor struct {
	memory  [memoryLength]byte
	address int
}

// Program is brainfuck-program.
type Program struct {
	code    []rune
	pointer int
	len     int
}

// New make new Processor.
func New() *Processor {
	return &Processor{
		[memoryLength]byte{},
		0,
	}
}

// right increases the address cyclically.
func (p *Processor) right() {
	p.address++
	if p.address >= memoryLength {
		p.address -= memoryLength
	}
}

// left reduces the address cyclically.
func (p *Processor) left() {
	p.address--
	if p.address < 0 {
		p.address = memoryLength - p.address
	}
}

// Run interpretation brainfuck-program.
func (p *Processor) Run(prog *Program) error {
	for ; prog.pointer < prog.len; prog.pointer++ {
		switch prog.code[prog.pointer] {
		case right:
			p.right()
		case left:
			p.left()
		case plus:
			p.memory[p.address]++
		case minus:
			p.memory[p.address]--
		case openLoop:
			if p.memory[p.address] == 0 {
				loop := 1
				for loop != 0 {
					prog.pointer++
					if p.memory[p.address] == openLoop {
						loop++
					}
					if p.memory[p.address] == closeLoop {
						loop--
					}
				}
			}
		case closeLoop:
			if p.memory[p.address] != 0 {
				loop := 1
				for loop != 0 {
					prog.pointer--
					if prog.code[prog.pointer] == openLoop {
						loop--
					}
					if prog.code[prog.pointer] == closeLoop {
						loop++
					}
				}
			}
		case print:
			fmt.Print(string(p.memory[p.address]))
		case read:
			var value byte
			_, err := fmt.Scanln(&value)
			if err != nil {
				return fmt.Errorf("Invalid input: %s", err)
			} else {
				p.memory[p.address] = value
			}

		}
	}
	return nil
}

// Load get brainfuck program from user and makes it interpretable.
func Load() (*Program, error) {
	// read path from console
	var userPath string
	_, err := fmt.Scan(&userPath)
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
		code:    bytes.Runes(rawProgram),
		pointer: 0,
	}
	p.len = len(p.code)

	return p, nil
}
