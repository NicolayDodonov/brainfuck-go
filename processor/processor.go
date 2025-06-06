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
	addressPointer int
}

type Program struct {
	code           []rune
	programPointer int
	len            int
}

// make new Processor
func New() *Processor {
	return &Processor{
		[30000]byte{},
		0,
	}
}

func (p *Processor) up() {
	p.addressPointer++
	if p.addressPointer >= 30000 {
		p.addressPointer -= 30000
	}
}

func (p *Processor) down() {
	p.addressPointer--
	if p.addressPointer < 0 {
		p.addressPointer = 30000 - p.addressPointer
	}
}

// run - processor function to interpretation brainfuck-program.
func (p *Processor) Run(prog *Program) error {
	for ; prog.programPointer < prog.len; prog.programPointer++ {
		switch prog.code[prog.programPointer] {
		case right:
			p.up()
		case left:
			p.down()
		case plus:
			p.line[p.addressPointer]++
		case minus:
			p.line[p.addressPointer]--
		case openLoop:
			if p.line[p.addressPointer] == 0 {
				loop := 1
				for loop != 0 {
					prog.programPointer++
					if p.line[p.addressPointer] == openLoop {
						loop++
					}
					if p.line[p.addressPointer] == closeLoop {
						loop--
					}
				}
			}
		case closeLoop:
			if p.line[p.addressPointer] != 0 {
				loop := 1
				for loop != 0 {
					prog.programPointer--
					if prog.code[prog.programPointer] == openLoop {
						loop--
					}
					if prog.code[prog.programPointer] == closeLoop {
						loop++
					}
				}
			}
		case print:
			fmt.Print(string(p.line[p.addressPointer]))
		case read:
			var value byte
			_, err := fmt.Scanln(&value)
			if err != nil {
				return err
			}
			p.line[p.addressPointer] = value
		}
	}
	return nil
}

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
		code:           bytes.Runes(rawProgram),
		programPointer: 0,
	}
	p.len = len(p.code)

	//clear program

	return p, nil
}
