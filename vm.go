package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

type vm struct {
	tape tape
	pointer int
	instructions []instruction
	insPtr int
}

func newVm (instr string) vm {
	var ins []instruction
	for _, chr := range instr {
		var i = charToInstruction(chr)
		if i != -1 {
			ins = append(ins, i)
		}
	}

	return vm{
		tape: newTape(),
		pointer: 0,
		instructions: ins,
	}
}

func (v *vm) interpretAll () {
	for v.insPtr < len(v.instructions) {
		v.interpretInstruction(v.instructions[v.insPtr])
		v.insPtr++
	}
}

func (v *vm) interpretInstruction (ins instruction) {
	switch ins {
	case increment:
		v.tape.values[v.pointer]++
	case decrement:
		v.tape.values[v.pointer]--
	case left:
		v.moveLeft()
	case right:
		v.moveRight()
	case output:
		fmt.Printf("%c", v.tape.values[v.pointer])
	case input:
		v.takeInput()
	case startLoop:
		v.doStartLoop()
	case endLoop:
		v.doEndLoop()
	}
}

func (v *vm) doEndLoop () {
	if v.tape.values[v.pointer] == 0 {
		return
	}

	v.insPtr--
	var skipCount = 0
	for v.instructions[v.insPtr] != startLoop || skipCount > 0 {
		if v.instructions[v.insPtr] == endLoop {
			skipCount++
		} else if v.instructions[v.insPtr] == startLoop {
			skipCount--
		}
		v.insPtr--
	}
}

func (v *vm) doStartLoop () {
	if v.tape.values[v.pointer] != 0 {
		return
	}

	v.insPtr++
	var skipCount = 0
	for v.instructions[v.insPtr] != endLoop || skipCount > 0 {
		if v.instructions[v.insPtr] == startLoop {
			skipCount++
		} else if v.instructions[v.insPtr] == endLoop {
			skipCount--
		}
		v.insPtr++
	}
}

func (v *vm) takeInput () {
	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}
	v.tape.values[v.pointer] = fmt.Sprintf("%c", char)[0]
}

func (v *vm) moveLeft () {
	v.pointer--
	if v.pointer == -1 {
		v.pointer = 0
		v.tape.addToBeginning(0)
	}
}

func (v *vm) moveRight () {
	v.pointer++
	if v.pointer == len(v.tape.values) {
		v.tape.addToEnd(0)
	}
}