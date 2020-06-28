package main

import (
	// "fmt"
	// "strconv"
)

type CPU struct {
	A byte
	B byte
	C byte
	D byte
	E byte
	F byte // used for storing z, n, h, c flags | ordered like znhc0000
	H byte
	L byte
	PC uint16
	SP uint16
	Cycles int
	Opcode byte
	bus Bus
	biosSkip bool
	IME byte
}

func (cpu *CPU) tick() {
	if cpu.Cycles == 0 {
		// cpu.debugCPU()
		// cpu.setPCBreakpoint(0xC00B)
		// fmt.Println(cpu.bus.read(cpu.PC))
		// fmt.Println(strconv.FormatUint(uint64(cpu.PC), 16))
		cpu.Opcode = cpu.bus.read(cpu.PC)
		cpu.PC++
		cpu.Cycles = opcodes[cpu.Opcode].Cycles
		opcodes[cpu.Opcode].Exec(cpu)
		
		
	}
	cpu.Cycles--
}



func (cpu *CPU) ZFlag(value byte) {
	if value == 1 {
		cpu.F |= (1 << 7)
	} else {
		cpu.F &= 0x7F
	}
}

func (cpu *CPU) NFlag(value byte) {
	if value == 1 {
		cpu.F |= (1 << 6)
	} else {
		cpu.F &= 0xBF
	}
}
func (cpu *CPU) HFlag(value byte) {
	if value == 1 {
		cpu.F |= (1 << 5)
	} else {
		cpu.F &= 0xDF
	}
}
func (cpu *CPU) CFlag(value byte) {
	if value == 1 {
		cpu.F |= (1 << 4)
	} else {
		cpu.F &= 0xEF
	}
}

func (cpu *CPU) getZFlag() byte {
	return (cpu.F & 0x80) >> 7
}

func (cpu *CPU) getCFlag() byte {
	return (cpu.F & 0x10) >> 4
}

