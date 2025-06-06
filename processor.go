package main

import "fmt"

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

type processor struct {
	line    [30000]byte
	pointer int
	stack
}

type stack struct {
	head int
	pool [10]int
}

func newProcessor() *processor {
	return &processor{
		[30000]byte{},
		0,
		stack{
			head: -1,
			pool: [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
}

func (p *processor) run(prog *program) error {
	for pointer := 0; pointer < prog.len; pointer++ {
		switch prog.code[pointer] {
		case right:
			p.pointer++
			if p.pointer >= 30000 {
				p.pointer -= 30000
			}
		case left:
			p.pointer--
			if p.pointer <= 0 {
				p.pointer += 30000
			}
		case plus:
			p.line[p.pointer]++
		case minus:
			p.line[p.pointer]--
		case openLoop:
			//todo: add
		case closeLoop:
			//todo: add
		case print:
			fmt.Print(p.line[p.pointer])
		case read:
			value, err := fmt.Scan()
			if err != nil {
				return err
			}
			p.line[p.pointer] = byte(value)
		}
	}
	return nil
}

// push add address to stack and up stack.head.
// Can get error if stack overflow.
func (s stack) push(adr int) error {
	if s.head >= -1 || s.head < len(s.pool) {
		s.head++
		s.pool[s.head] = adr
		return nil
	} else {
		//todo: add error
		return nil
	}
}

// pull get address from stack.head, return and down head.
// Can get error if stack is empty/
func (s stack) pull() (int, error) {
	if s.head > -1 || s.head < len(s.pool) {
		value := s.pool[s.head]
		s.head--
		return value, nil
	}
	//todo: add error
	return 0, nil
}
