package main

import "fmt"

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
type processor struct {
	line           [30000]byte
	programPointer int
}

// make new processor
func newProcessor() *processor {
	return &processor{
		[30000]byte{},
		0,
	}
}

func (p *processor) up() {
	p.programPointer++
	if p.programPointer >= 30000 {
		p.programPointer -= 30000
	}
}

func (p *processor) down() {
	p.programPointer--
	if p.programPointer < 0 {
		p.programPointer = 30000 - p.programPointer
	}
}

// run - processor function to interpretation brainfuck-program.
func (p *processor) run(prog *program) error {
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
			value, err := fmt.Scan()
			if err != nil {
				return err
			}
			p.line[p.programPointer] = byte(value)
		}
	}
	return nil
}
