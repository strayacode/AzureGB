package main

import (
	"fmt"
	"os"
	"strconv"
)

type Opcode struct {
	Cycles int
	Exec func(*CPU)
}

func op_null(cpu *CPU) {
	fmt.Println("unimplemented opcode!", "0x" + strconv.FormatUint(uint64(cpu.Opcode), 16))
	cpu.debugCPU()
	// cpu.debugVRAM()
	os.Exit(3)
}

func cbop_null(cpu *CPU) {
	fmt.Println("unimplemented cb opcode!", "0x" + strconv.FormatUint(uint64(cpu.Opcode), 16))
	cpu.debugCPU()
	os.Exit(3)
}

// NOP 
func op_0x00(cpu *CPU) {

}

// LD BC, u16 
func op_0x01(cpu *CPU) {
	cpu.B = cpu.bus.read(cpu.PC + 1)
	cpu.C = cpu.bus.read(cpu.PC)
	cpu.PC += 2
}

// LD (BC), A 
func op_0x02(cpu *CPU) {
	cpu.bus.write((uint16(cpu.B) << 8 | uint16(cpu.C)), cpu.A)
}

// INC BC 
func op_0x03(cpu *CPU) {
	cpu.C++
	if cpu.C == 0 {
		cpu.B++
	}
}

// INC B 
func op_0x04(cpu *CPU) {
	if ((cpu.B & 0xF) + (1 & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.B++;
	if cpu.B == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
}

// DEC B 
func op_0x05(cpu *CPU) {
	if ((cpu.B & 0xF) - (1 & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.B--;
	if cpu.B == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
}

// LD B, u8 
func op_0x06(cpu *CPU) {
	cpu.B = cpu.bus.read(cpu.PC)
	cpu.PC++
}

// RLCA 
func op_0x07(cpu *CPU) {
	tempcarry := cpu.A >> 7
	cpu.CFlag(tempcarry)
	cpu.A = (cpu.A << 1 | tempcarry)
	cpu.ZFlag(0)
	cpu.NFlag(0)
	cpu.HFlag(0)

}

// LD (u16), SP 
func op_0x08(cpu *CPU) {
	cpu.bus.write(cpu.bus.read16(cpu.PC), byte(cpu.SP & 0xFF))
	cpu.bus.write(cpu.bus.read16(cpu.PC) + 1, byte(cpu.SP >> 8))
	cpu.PC += 2
}

// ADD HL, BC
func op_0x09(cpu *CPU) {
	hl := uint16(cpu.H) << 8 | uint16(cpu.L)
	bc := uint16(cpu.B) << 8 | uint16(cpu.C)
	if uint32(hl) + uint32(bc) > 0xFFFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	if ((hl & 0x0FFF) + ((bc) & 0x0FFF)) & 0x1000 == 0x1000 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	hl += bc
	
	
	cpu.NFlag(0)
	cpu.H = byte((hl >> 8) & 0xFF)
	cpu.L = byte(hl & 0xFF)
}

// LD A, (BC) 
func op_0x0A(cpu *CPU) {
	cpu.A = cpu.bus.read(uint16(cpu.B) << 8 | uint16(cpu.C))
}

// DEC BC 
func op_0x0B(cpu *CPU) {
	cpu.C--
	if cpu.C == 0xFF {
		cpu.B--
	}
}

// INC C DONE
func op_0x0C(cpu *CPU) {
	if ((cpu.C & 0xF) + (1 & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.C++;
	if cpu.C == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
}

// DEC C 
func op_0x0D(cpu *CPU) {
	if ((cpu.C & 0xF) - (1 & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.C--;
	if cpu.C == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
}

// LD C, u8 
func op_0x0E(cpu *CPU) {
	cpu.C = cpu.bus.read(cpu.PC)
	cpu.PC++
}

// RRCA 
func op_0x0F(cpu *CPU) {
	cpu.CFlag(cpu.A & 0x1)
	cpu.A = (cpu.getCFlag() << 7 | cpu.A >> 1)
	cpu.ZFlag(0)
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// STOP 
func op_0x10(cpu *CPU) {
	if cpu.bus.read(cpu.PC) == 0x00 {
		cpu.PC++
	}
}

// LD DE, u16 
func op_0x11(cpu *CPU) {
	cpu.D = cpu.bus.read(cpu.PC + 1)
	cpu.E = cpu.bus.read(cpu.PC)
	cpu.PC += 2
}

// LD (DE), A 
func op_0x12(cpu *CPU) {
	cpu.bus.write((uint16(cpu.D) << 8 | uint16(cpu.E)), cpu.A)
}

// INC DE 
func op_0x13(cpu *CPU) {
	cpu.E++
	if cpu.E == 0 {
		cpu.D++
	}
}

// INC D 
func op_0x14(cpu *CPU) {
	if ((cpu.D & 0xF) + (1 & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.D++;
	if cpu.D == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
}

// DEC D 
func op_0x15(cpu *CPU) {
	if ((cpu.D & 0xF) - (1 & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.D--;
	if cpu.D == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
}

// LD D, u8 
func op_0x16(cpu *CPU) {
	cpu.D = cpu.bus.read(cpu.PC)
	cpu.PC++
}

// RLA 
func op_0x17(cpu *CPU) {
	tempcarry := cpu.getCFlag()
	newcarry := cpu.A >> 7
	cpu.A = (cpu.A << 1 | tempcarry)
	cpu.CFlag(newcarry)
	cpu.ZFlag(0)
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// JR i8 
func op_0x18(cpu *CPU) {
	offset := int8(cpu.bus.read(cpu.PC) + 1)
	cpu.PC += uint16(offset)
}

// ADD HL, DE
func op_0x19(cpu *CPU) {
	hl := uint16(cpu.H) << 8 | uint16(cpu.L)
	de := uint16(cpu.D) << 8 | uint16(cpu.E)
	if uint32(hl) + uint32(de) > 0xFFFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	if ((hl & 0x0FFF) + ((de) & 0x0FFF)) & 0x1000 == 0x1000 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	hl += de
	
	
	cpu.NFlag(0)
	cpu.H = byte((hl >> 8) & 0xFF)
	cpu.L = byte(hl & 0xFF)
}

// LD A, (DE) DONE
func op_0x1A(cpu *CPU) {
	cpu.A = cpu.bus.read(uint16(cpu.D) << 8 | uint16(cpu.E))
}

// DEC DE 
func op_0x1B(cpu *CPU) {
	cpu.E--
	if cpu.E == 0xFF {
		cpu.D--
	}
}

// INC E 
func op_0x1C(cpu *CPU) {
	if ((cpu.E & 0xF) + (1 & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.E++;
	if cpu.E == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
}

// DEC E 
func op_0x1D(cpu *CPU) {
	if ((cpu.E & 0xF) - (1 & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.E--;
	if cpu.E == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
}

// LD E, u8 
func op_0x1E(cpu *CPU) {
	cpu.E = cpu.bus.read(cpu.PC)
	cpu.PC++
}

// RRA 
func op_0x1F(cpu *CPU) {
	tempcarry := cpu.getCFlag()
	newcarry := (cpu.A & 0x01)
	cpu.CFlag(newcarry)
	cpu.A = (tempcarry << 7 | cpu.A >> 1)
	cpu.ZFlag(0)
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// JR NZ, i8 DONE
func op_0x20(cpu *CPU) {
	if cpu.getZFlag() == 0 {
		offset := int8(cpu.bus.read(cpu.PC) + 1)
		cpu.PC += uint16(offset)
		cpu.Cycles += 4
	} else {
		cpu.PC++
	}

}

// LD HL, u16 DONE
func op_0x21(cpu *CPU) {
	cpu.H = cpu.bus.read(cpu.PC + 1)
	cpu.L = cpu.bus.read(cpu.PC)
	cpu.PC += 2
}

// LD (HL+), A 
func op_0x22(cpu *CPU) {
	cpu.bus.write((uint16(cpu.H) << 8 | uint16(cpu.L)), cpu.A)
	cpu.L++
	if cpu.L == 0 {
		cpu.H++
	}
}

// INC HL 
func op_0x23(cpu *CPU) {
	cpu.L++
	if cpu.L == 0 {
		cpu.H++
	}
}

// INC H 
func op_0x24(cpu *CPU) {
	if ((cpu.H & 0xF) + (1 & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.H++;
	if cpu.H == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
}

// DEC H 
func op_0x25(cpu *CPU) {
	if ((cpu.H & 0xF) - (1 & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.H--;
	if cpu.H == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
}

// LD H, u8
func op_0x26(cpu *CPU) {
	cpu.H = cpu.bus.read(cpu.PC)
	cpu.PC++
}

// DAA 
func op_0x27(cpu *CPU) {
	// addition
	if cpu.getNFlag() == 0 {
		if cpu.A > 0x99 || cpu.getCFlag() == 1 {
			cpu.A += 0x60
			cpu.CFlag(1)
		}
		if cpu.getHFlag() == 1 || (cpu.A & 0xF) > 0x09 {
			cpu.A += 0x06
		}
	} else {
		// subtraction
		if cpu.getCFlag() == 1 {
			cpu.A -= 0x60
		}
		if cpu.getHFlag() == 1 {
			cpu.A -= 0x06
		}
	}

	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.HFlag(0)
}

// JR Z, i8 
func op_0x28(cpu *CPU) {
	if cpu.getZFlag() == 1 {
		offset := int8(cpu.bus.read(cpu.PC) + 1)
		cpu.PC += uint16(offset)
		cpu.Cycles += 4
	} else {
		cpu.PC++
	}

}

// ADD HL, HL
func op_0x29(cpu *CPU) {
	hl := uint16(cpu.H) << 8 | uint16(cpu.L)
	if uint32(hl) + uint32(hl) > 0xFFFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	if ((hl & 0x0FFF) + ((hl) & 0x0FFF)) & 0x1000 == 0x1000 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	hl += hl
	
	
	cpu.NFlag(0)
	cpu.H = byte((hl >> 8) & 0xFF)
	cpu.L = byte(hl & 0xFF)
}

// LD A, (HL+) 
func op_0x2A(cpu *CPU) {
	cpu.A = cpu.bus.read((uint16(cpu.H) << 8 | uint16(cpu.L)))
	cpu.L++
	if cpu.L == 0 {
		cpu.H++
	}
}

// DEC HL
func op_0x2B(cpu *CPU) {
	cpu.L--
	if cpu.L == 0xFF {
		cpu.H--
	}
}

// INC L 
func op_0x2C(cpu *CPU) {
	if ((cpu.L & 0xF) + (1 & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.L++;
	if cpu.L == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
}

// DEC L
func op_0x2D(cpu *CPU) {
	if ((cpu.L & 0xF) - (1 & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.L--;
	if cpu.L == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
}

// LD L, u8
func op_0x2E(cpu *CPU) {
	cpu.L = cpu.bus.read(cpu.PC)
	cpu.PC++

}

// CPL
func op_0x2F(cpu *CPU) {
	cpu.A = ^cpu.A
	cpu.NFlag(1)
	cpu.HFlag(1)

}

// JR NC, i8 
func op_0x30(cpu *CPU) {
	if cpu.getCFlag() == 0 {
		offset := int8(cpu.bus.read(cpu.PC) + 1)
		cpu.PC += uint16(offset)
		cpu.Cycles += 4
	} else {
		cpu.PC++
	}

}

// LD SP, u16 
func op_0x31(cpu *CPU) {
	cpu.SP = cpu.bus.read16(cpu.PC)
	cpu.PC += 2
}

// LD (HL-), A 
func op_0x32(cpu *CPU) {
	cpu.bus.write((uint16(cpu.H) << 8 | uint16(cpu.L)), cpu.A)
	cpu.L--
	if cpu.L == 0xFF {
		cpu.H--
	}
}

// INC SP
func op_0x33(cpu *CPU) {
	cpu.SP++
}

// INC (HL) 
func op_0x34(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	if ((result & 0xF) + (1 & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	
	result++
	if result == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)


}

// DEC (HL) 
func op_0x35(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	if ((result & 0xF) - (1 & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	
	result--
	if result == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)


}

// LD (HL), u8
func op_0x36(cpu *CPU) {
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), cpu.bus.read(cpu.PC))
	cpu.PC++
}

// SCF
func op_0x37(cpu *CPU) {
	cpu.CFlag(1)
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// JR C, i8
func op_0x38(cpu *CPU) {
	if cpu.getCFlag() == 1 {
		offset := int8(cpu.bus.read(cpu.PC) + 1)
		cpu.PC += uint16(offset)
		cpu.Cycles += 4
	} else {
		cpu.PC++
	}
}

// ADD HL, SP
func op_0x39(cpu *CPU) {
	hl := uint16(cpu.H) << 8 | uint16(cpu.L)
	if uint32(hl) + uint32(cpu.SP) > 0xFFFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	if ((hl & 0x0FFF) + ((cpu.SP) & 0x0FFF)) & 0x1000 == 0x1000 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	hl += cpu.SP
	
	
	cpu.NFlag(0)
	cpu.H = byte((hl >> 8) & 0xFF)
	cpu.L = byte(hl & 0xFF)
}

// LD A, (HL-) 
func op_0x3A(cpu *CPU) {
	cpu.A = cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	cpu.L--
	if cpu.L == 0xFF {
		cpu.H--
	}
}

// DEC SP
func op_0x3B(cpu *CPU) {
	cpu.SP--
}

// INC A 
func op_0x3C(cpu *CPU) {
	if ((cpu.A & 0xF) + (1 & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.A++;
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
}

// DEC A 
func op_0x3D(cpu *CPU) {
	if ((cpu.A & 0xF) - (1 & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.A--;
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)

}

// LD A, u8 
func op_0x3E(cpu *CPU) {
	cpu.A = cpu.bus.read(cpu.PC)
	cpu.PC++
}

// CCF
func op_0x3F(cpu *CPU) {
	carry := cpu.getCFlag()
	if carry == 1 {
		cpu.CFlag(0)
	} else {
		cpu.CFlag(1)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// LD B, B 
func op_0x40(cpu *CPU) {
	cpu.B = cpu.B
}

// LD B, C 
func op_0x41(cpu *CPU) {
	cpu.B = cpu.C
}

// LD B, D 
func op_0x42(cpu *CPU) {
	cpu.B = cpu.D
}

// LD B, E 
func op_0x43(cpu *CPU) {
	cpu.B = cpu.E
}

// LD B, H 
func op_0x44(cpu *CPU) {
	cpu.B = cpu.H
}

// LD B, L 
func op_0x45(cpu *CPU) {
	cpu.B = cpu.L
}

// LD B, (HL) 
func op_0x46(cpu *CPU) {
	cpu.B = cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
}

// LD B, A 
func op_0x47(cpu *CPU) {
	cpu.B = cpu.A
}

// LD C, B 
func op_0x48(cpu *CPU) {
	cpu.C = cpu.B
}

// LD C, C 
func op_0x49(cpu *CPU) {
	cpu.C = cpu.C
}

// LD C, D 
func op_0x4A(cpu *CPU) {
	cpu.C = cpu.D
}

// LD C, E 
func op_0x4B(cpu *CPU) {
	cpu.C = cpu.E
}

// LD C, H 
func op_0x4C(cpu *CPU) {
	cpu.C = cpu.H
}

// LD C, L 
func op_0x4D(cpu *CPU) {
	cpu.C = cpu.L
}

// LD C, (HL) 
func op_0x4E(cpu *CPU) {
	cpu.C = cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
}

// LD C, A 
func op_0x4F(cpu *CPU) {
	cpu.C = cpu.A
}

// LD D, B 
func op_0x50(cpu *CPU) {
	cpu.D = cpu.B
}

// LD D, C 
func op_0x51(cpu *CPU) {
	cpu.D = cpu.C
}

// LD D, D 
func op_0x52(cpu *CPU) {
	cpu.D = cpu.D
}

// LD D, E 
func op_0x53(cpu *CPU) {
	cpu.D = cpu.E
}

// LD D, H 
func op_0x54(cpu *CPU) {
	cpu.D = cpu.H
}

// LD D, L 
func op_0x55(cpu *CPU) {
	cpu.D = cpu.L
}

// LD D, (HL) 
func op_0x56(cpu *CPU) {
	cpu.D = cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
}

// LD D, A 
func op_0x57(cpu *CPU) {
	cpu.D = cpu.A
}

// LD E, B 
func op_0x58(cpu *CPU) {
	cpu.E = cpu.B
}

// LD E, C 
func op_0x59(cpu *CPU) {
	cpu.E = cpu.C
}

// LD E, D 
func op_0x5A(cpu *CPU) {
	cpu.E = cpu.D
}

// LD E, E 
func op_0x5B(cpu *CPU) {
	cpu.E = cpu.E
}

// LD E, H 
func op_0x5C(cpu *CPU) {
	cpu.E = cpu.H
}

// LD E, L 
func op_0x5D(cpu *CPU) {
	cpu.E = cpu.L
}

// LD E, (HL) 
func op_0x5E(cpu *CPU) {
	cpu.E = cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
}

// LD E, A 
func op_0x5F(cpu *CPU) {
	cpu.E = cpu.A
}

// LD H, B 
func op_0x60(cpu *CPU) {
	cpu.H = cpu.B
}
// LD H, C 
func op_0x61(cpu *CPU) {
	cpu.H = cpu.C
}
// LD H, D 
func op_0x62(cpu *CPU) {
	cpu.H = cpu.D
}
// LD H, E 
func op_0x63(cpu *CPU) {
	cpu.H = cpu.E
}
// LD H, H 
func op_0x64(cpu *CPU) {
	cpu.H = cpu.H
}
// LD H, L
func op_0x65(cpu *CPU) {
	cpu.H = cpu.L
}
// LD H, (HL) 
func op_0x66(cpu *CPU) {
	cpu.H = cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
}

// LD H, A 
func op_0x67(cpu *CPU) {
	cpu.H = cpu.A
}
// LD L, B 
func op_0x68(cpu *CPU) {
	cpu.L = cpu.B
}
// LD L, C 
func op_0x69(cpu *CPU) {
	cpu.L = cpu.C
}	

// LD L, D 
func op_0x6A(cpu *CPU) {
	cpu.L = cpu.D
}
// LD L, E 
func op_0x6B(cpu *CPU) {
	cpu.L = cpu.E
}
// LD L, H 
func op_0x6C(cpu *CPU) {
	cpu.L = cpu.H
}

// LD L, L 
func op_0x6D(cpu *CPU) {
	cpu.L = cpu.L
}

// LD L, (HL) 
func op_0x6E(cpu *CPU) {
	cpu.L = cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
}

// LD L, A 
func op_0x6F(cpu *CPU) {
	cpu.L = cpu.A
}

// LD (HL), B 
func op_0x70(cpu *CPU) {
	cpu.bus.write((uint16(cpu.H) << 8 | uint16(cpu.L)), cpu.B)
}

// LD (HL), C 
func op_0x71(cpu *CPU) {
	cpu.bus.write((uint16(cpu.H) << 8 | uint16(cpu.L)), cpu.C)
}

// LD (HL), D 
func op_0x72(cpu *CPU) {
	cpu.bus.write((uint16(cpu.H) << 8 | uint16(cpu.L)), cpu.D)
}

// LD (HL), E 
func op_0x73(cpu *CPU) {
	cpu.bus.write((uint16(cpu.H) << 8 | uint16(cpu.L)), cpu.E)
}

// LD (HL), H 
func op_0x74(cpu *CPU) {
	cpu.bus.write((uint16(cpu.H) << 8 | uint16(cpu.L)), cpu.H)
}

// LD (HL), L 
func op_0x75(cpu *CPU) {
	cpu.bus.write((uint16(cpu.H) << 8 | uint16(cpu.L)), cpu.L)
}

// HALT
func op_0x76(cpu *CPU) {
	cpu.halt = true
}

// LD (HL), A 
func op_0x77(cpu *CPU) {
	cpu.bus.write((uint16(cpu.H) << 8 | uint16(cpu.L)), cpu.A)
}

// LD A, B 
func op_0x78(cpu *CPU) {
	cpu.A = cpu.B
}

// LD A, C 
func op_0x79(cpu *CPU) {
	cpu.A = cpu.C
}

// LD A, D 
func op_0x7A(cpu *CPU) {
	cpu.A = cpu.D
}

// LD A, E 
func op_0x7B(cpu *CPU) {
	cpu.A = cpu.E
}

// LD A, H 
func op_0x7C(cpu *CPU) {
	cpu.A = cpu.H
}

// LD A, L 
func op_0x7D(cpu *CPU) {
	cpu.A = cpu.L
}

// LD A, (HL) 
func op_0x7E(cpu *CPU) {
	cpu.A = cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
}

// LD A, A 
func op_0x7F(cpu *CPU) {
	cpu.A = cpu.A
}

// ADD A, B
func op_0x80(cpu *CPU) {
	result := uint16(cpu.A) + uint16(cpu.B)
	if ((cpu.A & 0xF) + ((cpu.B) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.A += cpu.B
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	
	if result > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// ADD A, C
func op_0x81(cpu *CPU) {
	result := uint16(cpu.A) + uint16(cpu.C)
	if ((cpu.A & 0xF) + ((cpu.C) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.A += cpu.C
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	if result > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// ADD A, D
func op_0x82(cpu *CPU) {
	result := uint16(cpu.A) + uint16(cpu.D)
	if ((cpu.A & 0xF) + ((cpu.D) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.A += cpu.D
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	
	if result > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// ADD A, E
func op_0x83(cpu *CPU) {
	result := uint16(cpu.A) + uint16(cpu.E)
	if ((cpu.A & 0xF) + ((cpu.E) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.A += cpu.E
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	
	if result > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// ADD A, H
func op_0x84(cpu *CPU) {
	result := uint16(cpu.A) + uint16(cpu.H)
	if ((cpu.A & 0xF) + ((cpu.H) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.A += cpu.H
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	
	if result > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// ADD A, L
func op_0x85(cpu *CPU) {
	result := uint16(cpu.A) + uint16(cpu.L)
	if ((cpu.A & 0xF) + ((cpu.L) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.A += cpu.L
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	
	if result > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// ADD A, (HL)
func op_0x86(cpu *CPU) {
	result := uint16(cpu.A) + uint16(cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L)))
	if ((cpu.A & 0xF) + ((cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.A += cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	
	if result > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// ADD A, A
func op_0x87(cpu *CPU) {
	result := uint16(cpu.A) + uint16(cpu.A)
	if ((cpu.A & 0xF) + ((cpu.A) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.A += cpu.A
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	
	if result > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// ADC A, B 
func op_0x88(cpu *CPU) {
	bitOverflow := uint16(cpu.A) + uint16(cpu.getCFlag()) + uint16(cpu.B)
	if ((cpu.A & 0xF) + (cpu.getCFlag() & 0xF) + (cpu.B & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.A += (cpu.getCFlag() + cpu.B)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)

	if bitOverflow > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// ADC A, C
func op_0x89(cpu *CPU) {
	bitOverflow := uint16(cpu.A) + uint16(cpu.getCFlag()) + uint16(cpu.C)
	if ((cpu.A & 0xF) + (cpu.getCFlag() & 0xF) + (cpu.C & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.A += (cpu.getCFlag() + cpu.C)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)

	if bitOverflow > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// ADC A, D 
func op_0x8A(cpu *CPU) {
	bitOverflow := uint16(cpu.A) + uint16(cpu.getCFlag()) + uint16(cpu.D)
	if ((cpu.A & 0xF) + (cpu.getCFlag() & 0xF) + (cpu.D & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.A += (cpu.getCFlag() + cpu.D)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)

	if bitOverflow > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// ADC A, E
func op_0x8B(cpu *CPU) {
	bitOverflow := uint16(cpu.A) + uint16(cpu.getCFlag()) + uint16(cpu.E)
	if ((cpu.A & 0xF) + (cpu.getCFlag() & 0xF) + (cpu.E & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.A += (cpu.getCFlag() + cpu.E)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)

	if bitOverflow > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// ADC A, H
func op_0x8C(cpu *CPU) {
	bitOverflow := uint16(cpu.A) + uint16(cpu.getCFlag()) + uint16(cpu.H)
	if ((cpu.A & 0xF) + (cpu.getCFlag() & 0xF) + (cpu.H & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.A += (cpu.getCFlag() + cpu.H)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)

	if bitOverflow > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// ADC A, L 
func op_0x8D(cpu *CPU) {
	bitOverflow := uint16(cpu.A) + uint16(cpu.getCFlag()) + uint16(cpu.L)
	if ((cpu.A & 0xF) + (cpu.getCFlag() & 0xF) + (cpu.L & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.A += (cpu.getCFlag() + cpu.L)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)

	if bitOverflow > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// ADC A, (HL)
func op_0x8E(cpu *CPU) {
	bitOverflow := uint16(cpu.A) + uint16(cpu.getCFlag()) + uint16(cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L)))
	if ((cpu.A & 0xF) + (cpu.getCFlag() & 0xF) + (cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L)) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.A += (cpu.getCFlag() + cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L)))
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)

	if bitOverflow > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// ADC A, A
func op_0x8F(cpu *CPU) {
	bitOverflow := uint16(cpu.A) + uint16(cpu.getCFlag()) + uint16(cpu.A)
	if ((cpu.A & 0xF) + (cpu.getCFlag() & 0xF) + (cpu.A & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.A += (cpu.getCFlag() + cpu.A)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)

	if bitOverflow > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// SUB A, B 
func op_0x90(cpu *CPU) {
	if ((cpu.A & 0xF) - ((cpu.B) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if cpu.B > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.A -= cpu.B
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	

}

// SUB A, C
func op_0x91(cpu *CPU) {
	if ((cpu.A & 0xF) - ((cpu.C) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if cpu.C > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.A -= cpu.C
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	
	

}

// SUB A, D
func op_0x92(cpu *CPU) {
	if ((cpu.A & 0xF) - ((cpu.D) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if cpu.D > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.A -= cpu.D
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	

}

// SUB A, E
func op_0x93(cpu *CPU) {
	if ((cpu.A & 0xF) - ((cpu.E) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if cpu.E > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.A -= cpu.E
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	

}

// SUB A, H
func op_0x94(cpu *CPU) {
	if ((cpu.A & 0xF) - ((cpu.H) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if cpu.H > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.A -= cpu.H
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	

}

// SUB A, L
func op_0x95(cpu *CPU) {
	if ((cpu.A & 0xF) - ((cpu.L) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if cpu.L > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.A -= cpu.L
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	

}

// SUB A, (HL)
func op_0x96(cpu *CPU) {
	if ((cpu.A & 0xF) - ((cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L)) > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.A -= cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	

}

// SUB A, A
func op_0x97(cpu *CPU) {
	if ((cpu.A & 0xF) - ((cpu.A) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if cpu.A > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.A -= cpu.A
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	

}

// SBC A, B
func op_0x98(cpu *CPU) {
	result := cpu.A - (cpu.B + cpu.getCFlag())
	if ((cpu.A & 0xF) - (cpu.getCFlag() & 0xF) - (cpu.B & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.NFlag(1)
	if uint16(cpu.getCFlag()) + uint16(cpu.B) > uint16(cpu.A) {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.A = result
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
}



// SBC A, C
func op_0x99(cpu *CPU) {
	result := cpu.A - (cpu.C + cpu.getCFlag())
	if ((cpu.A & 0xF) - (cpu.getCFlag() & 0xF) - (cpu.C & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.NFlag(1)
	if uint16(cpu.getCFlag()) + uint16(cpu.C) > uint16(cpu.A) {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.A = result
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
}

// SBC A, D
func op_0x9A(cpu *CPU) {
	result := cpu.A - (cpu.D + cpu.getCFlag())
	if ((cpu.A & 0xF) - (cpu.getCFlag() & 0xF) - (cpu.D & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.NFlag(1)
	if uint16(cpu.getCFlag()) + uint16(cpu.D) > uint16(cpu.A) {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.A = result
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
}

// SBC A, E
func op_0x9B(cpu *CPU) {
	result := cpu.A - (cpu.E + cpu.getCFlag())
	if ((cpu.A & 0xF) - (cpu.getCFlag() & 0xF) - (cpu.E & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.NFlag(1)
	if uint16(cpu.getCFlag()) + uint16(cpu.E) > uint16(cpu.A) {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.A = result
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
}

// SBC A, H
func op_0x9C(cpu *CPU) {
	result := cpu.A - (cpu.H + cpu.getCFlag())
	if ((cpu.A & 0xF) - (cpu.getCFlag() & 0xF) - (cpu.H & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.NFlag(1)
	if uint16(cpu.getCFlag()) + uint16(cpu.H) > uint16(cpu.A) {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.A = result
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
}

// SBC A, L
func op_0x9D(cpu *CPU) {
	result := cpu.A - (cpu.L + cpu.getCFlag())
	if ((cpu.A & 0xF) - (cpu.getCFlag() & 0xF) - (cpu.L & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.NFlag(1)
	if uint16(cpu.getCFlag()) + uint16(cpu.L) > uint16(cpu.A) {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.A = result
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
}

// SBC A, (HL)
func op_0x9E(cpu *CPU) {
	result := cpu.A - (cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L)) + cpu.getCFlag())
	if ((cpu.A & 0xF) - (cpu.getCFlag() & 0xF) - (cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L)) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.NFlag(1)
	if uint16(cpu.getCFlag()) + uint16(cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))) > uint16(cpu.A) {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.A = result
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
}

// SBC A, A
func op_0x9F(cpu *CPU) {
	result := cpu.A - (cpu.A + cpu.getCFlag())
	if ((cpu.A & 0xF) - (cpu.getCFlag() & 0xF) - (cpu.A & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.NFlag(1)
	if uint16(cpu.getCFlag()) + uint16(cpu.A) > uint16(cpu.A) {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.A = result
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
}

// AND A, B
func op_0xA0(cpu *CPU) {
	cpu.A &= cpu.B
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
	cpu.CFlag(0)
}

// AND A, C
func op_0xA1(cpu *CPU) {
	cpu.A &= cpu.C
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
	cpu.CFlag(0)
}

// AND A, D 
func op_0xA2(cpu *CPU) {
	cpu.A &= cpu.D
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
	cpu.CFlag(0)
}

// AND A, E
func op_0xA3(cpu *CPU) {
	cpu.A &= cpu.E
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
	cpu.CFlag(0)
}

// AND A, H
func op_0xA4(cpu *CPU) {
	cpu.A &= cpu.H
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
	cpu.CFlag(0)
}

// AND A, L
func op_0xA5(cpu *CPU) {
	cpu.A &= cpu.L
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
	cpu.CFlag(0)
}

// AND A, (HL)
func op_0xA6(cpu *CPU) {
	cpu.A &= cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
	cpu.CFlag(0)
}

// AND A, A 
func op_0xA7(cpu *CPU) {
	cpu.A &= cpu.A
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
	cpu.CFlag(0)
}

// XOR A, B
func op_0xA8(cpu *CPU) {
	cpu.A ^= cpu.B
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)

}

// XOR A, C 
func op_0xA9(cpu *CPU) {
	cpu.A ^= cpu.C
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)

}

// XOR A, D
func op_0xAA(cpu *CPU) {
	cpu.A ^= cpu.D
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)

}

// XOR A, E
func op_0xAB(cpu *CPU) {
	cpu.A ^= cpu.E
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)

}

// XOR A, H
func op_0xAC(cpu *CPU) {
	cpu.A ^= cpu.H
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)

}

// XOR A, L 
func op_0xAD(cpu *CPU) {
	cpu.A ^= cpu.L
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)

}

// XOR A, (HL) 
func op_0xAE(cpu *CPU) {
	cpu.A ^= cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)

}

// XOR A, A DONE
func op_0xAF(cpu *CPU) {
	cpu.A ^= cpu.A
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)

}

// OR A, B
func op_0xB0(cpu *CPU) {
	cpu.A |= cpu.B
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
}

// OR A, C
func op_0xB1(cpu *CPU) {
	cpu.A |= cpu.C
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
}

// OR A, D
func op_0xB2(cpu *CPU) {
	cpu.A |= cpu.D
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
}

// OR A, E
func op_0xB3(cpu *CPU) {
	cpu.A |= cpu.E
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
}

// OR A, H
func op_0xB4(cpu *CPU) {
	cpu.A |= cpu.H
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
}

// OR A, L
func op_0xB5(cpu *CPU) {
	cpu.A |= cpu.L
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
}

// OR A, (HL)
func op_0xB6(cpu *CPU) {
	cpu.A |= cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
}

// OR A, A
func op_0xB7(cpu *CPU) {
	cpu.A |= cpu.A
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
}

// CP A, B
func op_0xB8(cpu *CPU) {
	result := cpu.A - cpu.B
	if result == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	if ((cpu.A & 0xF) - ((cpu.B) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if cpu.B > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// CP A, C
func op_0xB9(cpu *CPU) {
	result := cpu.A - cpu.C
	if result == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	if ((cpu.A & 0xF) - ((cpu.C) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if cpu.C > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// CP A, D 
func op_0xBA(cpu *CPU) {
	result := cpu.A - cpu.D
	if result == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	if ((cpu.A & 0xF) - ((cpu.D) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if cpu.D > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// CP A, E 
func op_0xBB(cpu *CPU) {
	result := cpu.A - cpu.E
	if result == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	if ((cpu.A & 0xF) - ((cpu.E) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if cpu.E > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// CP A, H
func op_0xBC(cpu *CPU) {
	result := cpu.A - cpu.H
	if result == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	if ((cpu.A & 0xF) - ((cpu.H) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if cpu.H > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// CP A, L
func op_0xBD(cpu *CPU) {
	result := cpu.A - cpu.L
	if result == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	if ((cpu.A & 0xF) - ((cpu.L) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if cpu.L > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}


// CP A, (HL) 
func op_0xBE(cpu *CPU) {
	result := cpu.A - cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	if result == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	if ((cpu.A & 0xF) - ((cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L)) > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// CP A, A
func op_0xBF(cpu *CPU) {
	result := cpu.A - cpu.A
	if result == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	if ((cpu.A & 0xF) - ((cpu.A) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if cpu.A > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// RET NZ 
func op_0xC0(cpu *CPU) {
	// if Z Flag is reset condition
	if cpu.getZFlag() == 0 {
		cpu.PC = cpu.bus.read16(cpu.SP)
		cpu.SP += 2
		cpu.Cycles += 12
	} 
}

// POP BC 
func op_0xC1(cpu *CPU) {
	cpu.C = cpu.bus.read(cpu.SP)
	cpu.SP++
	cpu.B = cpu.bus.read(cpu.SP)
	cpu.SP++
}

// JP NZ, u16
func op_0xC2(cpu *CPU) {
	if cpu.getZFlag() == 0 {
		cpu.PC = cpu.bus.read16(cpu.PC)
		cpu.Cycles += 4
	} else {
		cpu.PC += 2
	}
}

// JP u16
func op_0xC3(cpu *CPU) {
	cpu.PC = cpu.bus.read16(cpu.PC)
}

// CALL NZ, u16 
func op_0xC4(cpu *CPU) {
	// if z flag is not set
	if cpu.getZFlag() == 0 {
		hi := byte((cpu.PC + 2) >> 8)
		lo := byte(((cpu.PC + 2) & 0xFF))
		cpu.SP--
		cpu.bus.write(cpu.SP, hi)
		cpu.SP--
		cpu.bus.write(cpu.SP, lo)
		cpu.PC = cpu.bus.read16(cpu.PC)
		cpu.Cycles += 12
	} else {
		cpu.PC += 2
	}
}

// PUSH BC 
func op_0xC5(cpu *CPU) {
	cpu.SP--
	cpu.bus.write(cpu.SP, cpu.B)
	cpu.SP--
	cpu.bus.write(cpu.SP, cpu.C)
}

// ADD A, u8
func op_0xC6(cpu *CPU) {
	result := uint16(cpu.A) + uint16(cpu.bus.read(cpu.PC))
	if ((cpu.A & 0xF) + ((cpu.bus.read(cpu.PC)) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.A += cpu.bus.read(cpu.PC)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	
	if result > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.PC++
}

// RST 0x00 
func op_0xC7(cpu *CPU) {
	hi := byte((cpu.PC) >> 8)
	lo := byte(((cpu.PC) & 0xFF))
	cpu.SP--
	cpu.bus.write(cpu.SP, hi)
	cpu.SP--
	cpu.bus.write(cpu.SP, lo)
	cpu.PC = 0x0000
}

// RET Z 
func op_0xC8(cpu *CPU) {
	if cpu.getZFlag() == 1 {
		cpu.PC = cpu.bus.read16(cpu.SP)
		cpu.SP += 2
		cpu.Cycles += 12
	} 
}

// RET 
func op_0xC9(cpu *CPU) {
	cpu.PC = cpu.bus.read16(cpu.SP)
	cpu.SP += 2
}

// JP Z, u16
func op_0xCA(cpu *CPU) {
	if cpu.getZFlag() == 1 {
		cpu.PC = cpu.bus.read16(cpu.PC)
		cpu.Cycles += 4
	} else {
		cpu.PC += 2
	}
}

// PREFIX CB 
func op_0xCB(cpu *CPU) {
	cpu.Opcode = cpu.bus.read(cpu.PC)
	cbopcodes[cpu.Opcode].Exec(cpu)
	cpu.PC++
}

// CALL Z, u16 
func op_0xCC(cpu *CPU) {
	
	if cpu.getZFlag() == 1 {
		hi := byte((cpu.PC + 2) >> 8)
		lo := byte(((cpu.PC + 2) & 0xFF))
		cpu.SP--
		cpu.bus.write(cpu.SP, hi)
		cpu.SP--
		cpu.bus.write(cpu.SP, lo)
		cpu.PC = cpu.bus.read16(cpu.PC)
		cpu.Cycles += 12
	} else {
		cpu.PC += 2
	}
}

// CALL u16 DONE
func op_0xCD(cpu *CPU) {
	hi := byte((cpu.PC + 2) >> 8)
	lo := byte(((cpu.PC + 2) & 0xFF))
	cpu.SP--
	cpu.bus.write(cpu.SP, hi)
	cpu.SP--
	cpu.bus.write(cpu.SP, lo)
	cpu.PC = cpu.bus.read16(cpu.PC)
}

// ADC A, u8 
func op_0xCE(cpu *CPU) {
	bitOverflow := uint16(cpu.A) + uint16(cpu.getCFlag()) + uint16(cpu.bus.read(cpu.PC))
	if ((cpu.A & 0xF) + (cpu.getCFlag() & 0xF) + (cpu.bus.read(cpu.PC) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.A += (cpu.getCFlag() + cpu.bus.read(cpu.PC))
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)

	if bitOverflow > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.PC++
}

// RST 0x08 
func op_0xCF(cpu *CPU) {
	hi := byte((cpu.PC) >> 8)
	lo := byte(((cpu.PC) & 0xFF))
	cpu.SP--
	cpu.bus.write(cpu.SP, hi)
	cpu.SP--
	cpu.bus.write(cpu.SP, lo)
	cpu.PC = 0x0008
}

// RET NC 
func op_0xD0(cpu *CPU) {
	// if C Flag is reset condition
	if cpu.getCFlag() == 0 {
		cpu.PC = cpu.bus.read16(cpu.SP)
		cpu.SP += 2
		cpu.Cycles += 12
	} 
}

// POP DE 
func op_0xD1(cpu *CPU) {
	cpu.E = cpu.bus.read(cpu.SP)
	cpu.SP++
	cpu.D = cpu.bus.read(cpu.SP)
	cpu.SP++
}

// JP NC, u16
func op_0xD2(cpu *CPU) {
	if cpu.getCFlag() == 0 {
		cpu.PC = cpu.bus.read16(cpu.PC)
		cpu.Cycles += 4
	} else {
		cpu.PC += 2
	}
}

// CALL NC, u16 
func op_0xD4(cpu *CPU) {
	// if c flag is not set
	if cpu.getCFlag() == 0 {
		hi := byte((cpu.PC + 2) >> 8)
		lo := byte(((cpu.PC + 2) & 0xFF))
		cpu.SP--
		cpu.bus.write(cpu.SP, hi)
		cpu.SP--
		cpu.bus.write(cpu.SP, lo)
		cpu.PC = cpu.bus.read16(cpu.PC)
		cpu.Cycles += 12
	} else {
		cpu.PC += 2
	}
}


// PUSH DE 
func op_0xD5(cpu *CPU) {
	cpu.SP--
	cpu.bus.write(cpu.SP, cpu.D)
	cpu.SP--
	cpu.bus.write(cpu.SP, cpu.E)
}

// SUB A, u8 
func op_0xD6(cpu *CPU) {
	if ((cpu.A & 0xF) - ((cpu.bus.read(cpu.PC)) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if cpu.bus.read(cpu.PC) > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.A -= cpu.bus.read(cpu.PC)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	
	
	cpu.PC++
}

// RST 0x10
func op_0xD7(cpu *CPU) {
	hi := byte((cpu.PC) >> 8)
	lo := byte(((cpu.PC) & 0xFF))
	cpu.SP--
	cpu.bus.write(cpu.SP, hi)
	cpu.SP--
	cpu.bus.write(cpu.SP, lo)
	cpu.PC = 0x0010
}

// RET C 
func op_0xD8(cpu *CPU) {
	// if C Flag is set condition
	if cpu.getCFlag() == 1 {
		cpu.PC = cpu.bus.read16(cpu.SP)
		cpu.SP += 2
		cpu.Cycles += 12
	} 
}

// RETI 
func op_0xD9(cpu *CPU) {
	cpu.PC = cpu.bus.read16(cpu.SP)
	cpu.SP += 2
	cpu.bus.interrupt.IME = 1
}

// JP C, u16
func op_0xDA(cpu *CPU) {
	if cpu.getCFlag() == 1 {
		cpu.PC = cpu.bus.read16(cpu.PC)
		cpu.Cycles += 4
	} else {
		cpu.PC += 2
	}
}

// CALL C, u16 
func op_0xDC(cpu *CPU) {
	// if c flag is set
	if cpu.getCFlag() == 1 {
		hi := byte((cpu.PC + 2) >> 8)
		lo := byte(((cpu.PC + 2) & 0xFF))
		cpu.SP--
		cpu.bus.write(cpu.SP, hi)
		cpu.SP--
		cpu.bus.write(cpu.SP, lo)
		cpu.PC = cpu.bus.read16(cpu.PC)
		cpu.Cycles += 12
	} else {
		cpu.PC += 2
	}
}

// SBC A, u8 
func op_0xDE(cpu *CPU) {
	result := cpu.A - (cpu.bus.read(cpu.PC) + cpu.getCFlag())
	if ((cpu.A & 0xF) - (cpu.getCFlag() & 0xF) - (cpu.bus.read(cpu.PC) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.NFlag(1)
	if uint16(cpu.getCFlag()) + uint16(cpu.bus.read(cpu.PC)) > uint16(cpu.A) {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.A = result
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.PC++
	
}

// RST 0x18
func op_0xDF(cpu *CPU) {
	hi := byte((cpu.PC) >> 8)
	lo := byte(((cpu.PC) & 0xFF))
	cpu.SP--
	cpu.bus.write(cpu.SP, hi)
	cpu.SP--
	cpu.bus.write(cpu.SP, lo)
	cpu.PC = 0x0018
}

// LD (FF00 + u8), A 
func op_0xE0(cpu *CPU) {
	cpu.bus.write((0xFF00 + uint16(cpu.bus.read(cpu.PC))), cpu.A)
	cpu.PC++
}

// POP HL 
func op_0xE1(cpu *CPU) {
	cpu.L = cpu.bus.read(cpu.SP)
	cpu.SP++
	cpu.H = cpu.bus.read(cpu.SP)
	cpu.SP++
}

// LD (FF00 + C), A 
func op_0xE2(cpu *CPU) {
	cpu.bus.write((0xFF00 + uint16(cpu.C)), cpu.A)
}

// PUSH HL 
func op_0xE5(cpu *CPU) {
	cpu.SP--
	cpu.bus.write(cpu.SP, cpu.H)
	cpu.SP--
	cpu.bus.write(cpu.SP, cpu.L)
}

// AND A, u8 
func op_0xE6(cpu *CPU) {
	cpu.A &= cpu.bus.read(cpu.PC)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
	cpu.CFlag(0)
	cpu.PC++
}

// RST 0x20
func op_0xE7(cpu *CPU) {
	hi := byte((cpu.PC) >> 8)
	lo := byte(((cpu.PC) & 0xFF))
	cpu.SP--
	cpu.bus.write(cpu.SP, hi)
	cpu.SP--
	cpu.bus.write(cpu.SP, lo)
	cpu.PC = 0x0020
}

// ADD SP, i8
func op_0xE8(cpu *CPU) {
	i8 := int8(cpu.bus.read(cpu.PC))
	if (((cpu.SP) & 0xF) + (uint16(i8) & 0xF)) >= 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if (uint16(i8) & 0xFF) + (cpu.SP & 0xFF) > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.SP += uint16(i8)
	cpu.ZFlag(0)
	cpu.NFlag(0)
	cpu.PC++

}

// JP HL DONE
func op_0xE9(cpu *CPU) {
	cpu.PC = uint16(cpu.H) << 8 | uint16(cpu.L)
}

// LD (u16), A 
func op_0xEA(cpu *CPU) {
	cpu.bus.write((uint16(cpu.bus.read(cpu.PC + 1))) << 8 | uint16((cpu.bus.read(cpu.PC))), cpu.A)
	cpu.PC += 2
}

// XOR A, u8
func op_0xEE(cpu *CPU) {
	cpu.A ^= cpu.bus.read(cpu.PC)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
	cpu.PC++
}

// RST 0x28 
func op_0xEF(cpu *CPU) {
	hi := byte((cpu.PC) >> 8)
	lo := byte(((cpu.PC) & 0xFF))
	cpu.SP--
	cpu.bus.write(cpu.SP, hi)
	cpu.SP--
	cpu.bus.write(cpu.SP, lo)
	cpu.PC = 0x0028
}

// LD A, (FF00 + u8) 
func op_0xF0(cpu *CPU) {
	cpu.A = cpu.bus.read(0xFF00 + uint16(cpu.bus.read(cpu.PC)))
	cpu.PC++

}

// POP AF 
func op_0xF1(cpu *CPU) {
	cpu.F = cpu.bus.read(cpu.SP) & 0xF0
	cpu.SP++
	cpu.A = cpu.bus.read(cpu.SP)
	cpu.SP++
}

// LD A, (FF00 + C)
func op_0xF2(cpu *CPU) {
	cpu.A = cpu.bus.read(0xFF00 + uint16(cpu.C))
}

// DI 
func op_0xF3(cpu *CPU) {
	cpu.bus.interrupt.IME = 0
	cpu.bus.interrupt.IMEDelay = false
}

// PUSH AF 
func op_0xF5(cpu *CPU) {
	cpu.SP--
	cpu.bus.write(cpu.SP, cpu.A)
	cpu.SP--
	cpu.bus.write(cpu.SP, cpu.F)
}

// OR A, u8
func op_0xF6(cpu *CPU) {
	cpu.A |= cpu.bus.read(cpu.PC)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
	cpu.PC++
}

// RST 0x30
func op_0xF7(cpu *CPU) {
	hi := byte((cpu.PC) >> 8)
	lo := byte(((cpu.PC) & 0xFF))
	cpu.SP--
	cpu.bus.write(cpu.SP, hi)
	cpu.SP--
	cpu.bus.write(cpu.SP, lo)
	cpu.PC = 0x0030
}

// // ADD SP, i8
// func op_0xE8(cpu *CPU) {
// 	i8 := int8(cpu.bus.read(cpu.PC))
// 	if (((cpu.SP) & 0xF) + (uint16(i8) & 0xF)) >= 0x10 {
// 		cpu.HFlag(1)
// 	} else {
// 		cpu.HFlag(0)
// 	}
// 	if (uint16(i8) & 0xFF) + (cpu.SP & 0xFF) > 0xFF {
// 		cpu.CFlag(1)
// 	} else {
// 		cpu.CFlag(0)
// 	}
// 	cpu.SP += uint16(i8)
// 	cpu.ZFlag(0)
// 	cpu.NFlag(0)
// 	cpu.PC++

// }

// LD HL, SP + i8
func op_0xF8(cpu *CPU) {
	i8 := int8(cpu.bus.read(cpu.PC))
	result := uint16(cpu.SP) + uint16(i8)
	if ((cpu.SP & 0xF) + ((uint16(i8)) & 0xF)) >= 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if (uint16(i8) & 0xFF) + (cpu.SP & 0xFF) > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	// cpu.SP = uint16(result)
	cpu.H = byte(result >> 8)
	cpu.L = byte(result & 0xFF)
	cpu.ZFlag(0)
	cpu.NFlag(0)
	cpu.PC++
}

// LD SP, HL
func op_0xF9(cpu *CPU) {
	cpu.SP = uint16(cpu.H) << 8 | uint16(cpu.L)
}

// LD A, (u16) 
func op_0xFA(cpu *CPU) {
	addr := cpu.bus.read16(cpu.PC)
	cpu.A = cpu.bus.read(addr)
	cpu.PC += 2
}

// EI 
func op_0xFB(cpu *CPU) {
	cpu.bus.interrupt.IMEDelay = true
}

// CP A, u8 
func op_0xFE(cpu *CPU) {
	result := cpu.A - cpu.bus.read(cpu.PC)
	if result == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	if ((cpu.A & 0xF) - (cpu.bus.read(cpu.PC) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if cpu.bus.read(cpu.PC) > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.PC++
}

// RST 0x38
func op_0xFF(cpu *CPU) {
	hi := byte((cpu.PC) >> 8)
	lo := byte(((cpu.PC) & 0xFF))
	cpu.SP--
	cpu.bus.write(cpu.SP, hi)
	cpu.SP--
	cpu.bus.write(cpu.SP, lo)
	cpu.PC = 0x0038
}

// RLC B 
func cbop_0x00(cpu *CPU) {
	cpu.CFlag((cpu.B & 0x80) >> 7)
	cpu.B = (cpu.B << 1 | cpu.getCFlag())
	if cpu.B == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)

}

// RLC C
func cbop_0x01(cpu *CPU) {
	cpu.CFlag((cpu.C & 0x80) >> 7)
	cpu.C = (cpu.C << 1 | cpu.getCFlag())
	if cpu.C == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)

}

// RLC D
func cbop_0x02(cpu *CPU) {
	cpu.CFlag((cpu.D & 0x80) >> 7)
	cpu.D = (cpu.D << 1 | cpu.getCFlag())
	if cpu.D == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)

}

// RLC E 
func cbop_0x03(cpu *CPU) {
	cpu.CFlag((cpu.E & 0x80) >> 7)
	cpu.E = (cpu.E << 1 | cpu.getCFlag())
	if cpu.E == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)

}

// RLC H
func cbop_0x04(cpu *CPU) {
	cpu.CFlag((cpu.H & 0x80) >> 7)
	cpu.H = (cpu.H << 1 | cpu.getCFlag())
	if cpu.H == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)

}

// RLC L
func cbop_0x05(cpu *CPU) {
	cpu.CFlag((cpu.L & 0x80) >> 7)
	cpu.L = (cpu.L << 1 | cpu.getCFlag())
	if cpu.L == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)

}

// RLC (HL)
func cbop_0x06(cpu *CPU) {
	hl := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	cpu.CFlag((hl & 0x80) >> 7)
	hl = (hl << 1 | cpu.getCFlag())
	if hl == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), hl)
}

// RLC A
func cbop_0x07(cpu *CPU) {
	cpu.CFlag((cpu.A & 0x80) >> 7)
	cpu.A = (cpu.A << 1 | cpu.getCFlag())
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// RRC B 
func cbop_0x08(cpu *CPU) {
	cpu.CFlag(cpu.B & 0x1)
	cpu.B = (cpu.getCFlag() << 7 | cpu.B >> 1)
	if cpu.B == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// RRC C 
func cbop_0x09(cpu *CPU) {
	cpu.CFlag(cpu.C & 0x1)
	cpu.C = (cpu.getCFlag() << 7 | cpu.C >> 1)
	if cpu.C == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// RRC D
func cbop_0x0A(cpu *CPU) {
	cpu.CFlag(cpu.D & 0x1)
	cpu.D = (cpu.getCFlag() << 7 | cpu.D >> 1)
	if cpu.D == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// RRC E
func cbop_0x0B(cpu *CPU) {
	cpu.CFlag(cpu.E & 0x1)
	cpu.E = (cpu.getCFlag() << 7 | cpu.E >> 1)
	if cpu.E == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// RRC H 
func cbop_0x0C(cpu *CPU) {
	cpu.CFlag(cpu.H & 0x1)
	cpu.H = (cpu.getCFlag() << 7 | cpu.H >> 1)
	if cpu.H == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// RRC L
func cbop_0x0D(cpu *CPU) {
	cpu.CFlag(cpu.L & 0x1)
	cpu.L = (cpu.getCFlag() << 7 | cpu.L >> 1)
	if cpu.L == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}
// RRC (HL)
func cbop_0x0E(cpu *CPU) {
	hl := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	cpu.CFlag(hl & 0x1)
	hl = (cpu.getCFlag() << 7 | hl >> 1)
	if hl == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), hl)
}


// RRC A
func cbop_0x0F(cpu *CPU) {
	cpu.CFlag(cpu.A & 0x1)
	cpu.A = (cpu.getCFlag() << 7 | cpu.A >> 1)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// RL B 
func cbop_0x10(cpu *CPU) {
	tempcarry := cpu.getCFlag()
	newcarry := cpu.B >> 7
	cpu.B = (cpu.B << 1 | tempcarry)
	cpu.CFlag(newcarry)
	if cpu.B == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// RL C 
func cbop_0x11(cpu *CPU) {
	tempcarry := cpu.getCFlag()
	newcarry := cpu.C >> 7
	cpu.C = (cpu.C << 1 | tempcarry)
	cpu.CFlag(newcarry)
	if cpu.C == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// RL D
func cbop_0x12(cpu *CPU) {
	tempcarry := cpu.getCFlag()
	newcarry := cpu.D >> 7
	cpu.D = (cpu.D << 1 | tempcarry)
	cpu.CFlag(newcarry)
	if cpu.D == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// RL E
func cbop_0x13(cpu *CPU) {
	tempcarry := cpu.getCFlag()
	newcarry := cpu.E >> 7
	cpu.E = (cpu.E << 1 | tempcarry)
	cpu.CFlag(newcarry)
	if cpu.E == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// RL H
func cbop_0x14(cpu *CPU) {
	tempcarry := cpu.getCFlag()
	newcarry := cpu.H >> 7
	cpu.H = (cpu.H << 1 | tempcarry)
	cpu.CFlag(newcarry)
	if cpu.H == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// RL L
func cbop_0x15(cpu *CPU) {
	tempcarry := cpu.getCFlag()
	newcarry := cpu.L >> 7
	cpu.L = (cpu.L << 1 | tempcarry)
	cpu.CFlag(newcarry)
	if cpu.L == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// RL (HL)
func cbop_0x16(cpu *CPU) {
	hl := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	tempcarry := cpu.getCFlag()
	newcarry := hl >> 7
	hl = (hl << 1 | tempcarry)
	cpu.CFlag(newcarry)
	if hl == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), hl)
}
// RL A
func cbop_0x17(cpu *CPU) {
	tempcarry := cpu.getCFlag()
	newcarry := cpu.A >> 7
	cpu.A = (cpu.A << 1 | tempcarry)
	cpu.CFlag(newcarry)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// RR B
func cbop_0x18(cpu *CPU) {
	tempcarry := cpu.getCFlag()
	cpu.CFlag(cpu.B & 0x1)
	cpu.B = (tempcarry << 7 | cpu.B >> 1)
	if cpu.B == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// RR C 
func cbop_0x19(cpu *CPU) {
	tempcarry := cpu.getCFlag()
	cpu.CFlag(cpu.C & 0x1)
	cpu.C = (tempcarry << 7 | cpu.C >> 1)
	if cpu.C == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// RR D 
func cbop_0x1A(cpu *CPU) {
	tempcarry := cpu.getCFlag()
	cpu.CFlag(cpu.D & 0x1)
	cpu.D = (tempcarry << 7 | cpu.D >> 1)
	if cpu.D == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// RR E 
func cbop_0x1B(cpu *CPU) {
	tempcarry := cpu.getCFlag()
	cpu.CFlag(cpu.E & 0x1)
	cpu.E = (tempcarry << 7 | cpu.E >> 1)
	if cpu.E == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)


}

// RR H 
func cbop_0x1C(cpu *CPU) {
	tempcarry := cpu.getCFlag()
	cpu.CFlag(cpu.H & 0x1)
	cpu.H = (tempcarry << 7 | cpu.H >> 1)
	if cpu.H == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// RR L
func cbop_0x1D(cpu *CPU) {
	tempcarry := cpu.getCFlag()
	cpu.CFlag(cpu.L & 0x1)
	cpu.L = (tempcarry << 7 | cpu.L >> 1)
	if cpu.L == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// RR (HL) 
func cbop_0x1E(cpu *CPU) {
	tempcarry := cpu.getCFlag()
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	cpu.CFlag(result & 0x1)
	result = (tempcarry << 7 | result >> 1)
	if result == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// RR A 
func cbop_0x1F(cpu *CPU) {
	tempcarry := cpu.getCFlag()
	cpu.CFlag(cpu.A & 0x1)
	cpu.A = (tempcarry << 7 | cpu.A >> 1)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// SLA B
func cbop_0x20(cpu *CPU) {
	cpu.CFlag((cpu.B & (1 << 7) >> 7))
	cpu.B <<= 1
	if cpu.B == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// SLA C
func cbop_0x21(cpu *CPU) {
	cpu.CFlag((cpu.C & (1 << 7) >> 7))
	cpu.C <<= 1
	if cpu.C == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// SLA D
func cbop_0x22(cpu *CPU) {
	cpu.CFlag((cpu.D & (1 << 7) >> 7))
	cpu.D <<= 1
	if cpu.D == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// SLA E
func cbop_0x23(cpu *CPU) {
	cpu.CFlag((cpu.E & (1 << 7) >> 7))
	cpu.E <<= 1
	if cpu.E == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// SLA H
func cbop_0x24(cpu *CPU) {
	cpu.CFlag((cpu.H & (1 << 7) >> 7))
	cpu.H <<= 1
	if cpu.H == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// SLA L
func cbop_0x25(cpu *CPU) {
	cpu.CFlag((cpu.L & (1 << 7) >> 7))
	cpu.L <<= 1
	if cpu.L == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// SLA (HL)
func cbop_0x26(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	cpu.CFlag((result & (1 << 7) >> 7))
	result <<= 1
	if result == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// SLA A
func cbop_0x27(cpu *CPU) {
	cpu.CFlag((cpu.A & (1 << 7) >> 7))
	cpu.A <<= 1
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// SRA B
func cbop_0x28(cpu *CPU) {
	bit7 := cpu.B >> 7
	cpu.CFlag(cpu.B & 0x1)
	cpu.B = (bit7 << 7 | cpu.B >> 1)
	if cpu.B == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// SRA C
func cbop_0x29(cpu *CPU) {
	bit7 := cpu.C >> 7
	cpu.CFlag(cpu.C & 0x1)
	cpu.C = (bit7 << 7 | cpu.C >> 1)
	if cpu.C == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// SRA D
func cbop_0x2A(cpu *CPU) {
	bit7 := cpu.D >> 7
	cpu.CFlag(cpu.D & 0x1)
	cpu.D = (bit7 << 7 | cpu.D >> 1)
	if cpu.D == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// SRA E
func cbop_0x2B(cpu *CPU) {
	bit7 := cpu.E >> 7
	cpu.CFlag(cpu.E & 0x1)
	cpu.E = (bit7 << 7 | cpu.E >> 1)
	if cpu.E == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// SRA H
func cbop_0x2C(cpu *CPU) {
	bit7 := cpu.H >> 7
	cpu.CFlag(cpu.H & 0x1)
	cpu.H = (bit7 << 7 | cpu.H >> 1)
	if cpu.H == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// SRA L
func cbop_0x2D(cpu *CPU) {
	bit7 := cpu.L >> 7
	cpu.CFlag(cpu.L & 0x1)
	cpu.L = (bit7 << 7 | cpu.L >> 1)
	if cpu.L == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// SRA (HL)
func cbop_0x2E(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	bit7 := result >> 7
	cpu.CFlag(result & 0x1)
	result = (bit7 << 7 | result >> 1)
	if result == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// SRA A
func cbop_0x2F(cpu *CPU) {
	bit7 := cpu.A >> 7
	cpu.CFlag(cpu.A & 0x1)
	cpu.A = (bit7 << 7 | cpu.A >> 1)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// SWAP B
func cbop_0x30(cpu *CPU) {
	hi := cpu.B >> 4
	lo := cpu.B & 0xFF
	cpu.B = (lo << 4 | hi)
	if cpu.B == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
}

// SWAP C
func cbop_0x31(cpu *CPU) {
	hi := cpu.C >> 4
	lo := cpu.C & 0xFF
	cpu.C = (lo << 4 | hi)
	if cpu.C == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
}

// SWAP D
func cbop_0x32(cpu *CPU) {
	hi := cpu.D >> 4
	lo := cpu.D & 0xFF
	cpu.D = (lo << 4 | hi)
	if cpu.D == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
}

// SWAP E
func cbop_0x33(cpu *CPU) {
	hi := cpu.E >> 4
	lo := cpu.E & 0xFF
	cpu.E = (lo << 4 | hi)
	if cpu.E == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
}

// SWAP H
func cbop_0x34(cpu *CPU) {
	hi := cpu.H >> 4
	lo := cpu.H & 0xFF
	cpu.H = (lo << 4 | hi)
	if cpu.H == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
}

// SWAP L
func cbop_0x35(cpu *CPU) {
	hi := cpu.L >> 4
	lo := cpu.L & 0xFF
	cpu.L = (lo << 4 | hi)
	if cpu.L == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
}

// SWAP (HL)
func cbop_0x36(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	hi := result >> 4
	lo := result & 0xFF
	result = (lo << 4 | hi)
	if result == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// SWAP A
func cbop_0x37(cpu *CPU) {
	hi := cpu.A >> 4
	lo := cpu.A & 0xFF
	cpu.A = (lo << 4 | hi)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
}

// SRL B 
func cbop_0x38(cpu *CPU) {
	cpu.CFlag(cpu.B & 0x1)
	cpu.B >>= 1
	if cpu.B == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// SRL C
func cbop_0x39(cpu *CPU) {
	cpu.CFlag(cpu.C & 0x1)
	cpu.C >>= 1
	if cpu.C == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// SRL D
func cbop_0x3A(cpu *CPU) {
	cpu.CFlag(cpu.D & 0x1)
	cpu.D >>= 1
	if cpu.D == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// SRL E
func cbop_0x3B(cpu *CPU) {
	cpu.CFlag(cpu.E & 0x1)
	cpu.E >>= 1
	if cpu.E == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// SRL H
func cbop_0x3C(cpu *CPU) {
	cpu.CFlag(cpu.H & 0x1)
	cpu.H >>= 1
	if cpu.H == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// SRL L
func cbop_0x3D(cpu *CPU) {
	cpu.CFlag(cpu.L & 0x1)
	cpu.L >>= 1
	if cpu.L == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// SRL (HL)
func cbop_0x3E(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	cpu.CFlag(result & 0x1)
	result >>= 1
	if result == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// SRL A 
func cbop_0x3F(cpu *CPU) {
	cpu.CFlag(cpu.A & 0x1)
	cpu.A >>= 1
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// BIT 0, B
func cbop_0x40(cpu *CPU) {
	bit := cpu.B & 1
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 0, C
func cbop_0x41(cpu *CPU) {
	bit := cpu.C & 1
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 0, D
func cbop_0x42(cpu *CPU) {
	bit := cpu.D & 1
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 0, E
func cbop_0x43(cpu *CPU) {
	bit := cpu.E & 1
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 0, H
func cbop_0x44(cpu *CPU) {
	bit := cpu.H & 1
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 0, L
func cbop_0x45(cpu *CPU) {
	bit := cpu.L & 1
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 0, (HL)
func cbop_0x46(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	bit := result & 1
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// BIT 0, A
func cbop_0x47(cpu *CPU) {
	bit := cpu.A & 1
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 1, B
func cbop_0x48(cpu *CPU) {
	bit := (cpu.B & (1 << 1)) >> 1
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 1, C
func cbop_0x49(cpu *CPU) {
	bit := (cpu.C & (1 << 1)) >> 1
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 1, D
func cbop_0x4A(cpu *CPU) {
	bit := (cpu.D & (1 << 1)) >> 1
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 1, E
func cbop_0x4B(cpu *CPU) {
	bit := (cpu.E & (1 << 1)) >> 1
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 1, H
func cbop_0x4C(cpu *CPU) {
	bit := (cpu.H & (1 << 1)) >> 1
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 1, L
func cbop_0x4D(cpu *CPU) {
	bit := (cpu.L & (1 << 1)) >> 1
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 1, (HL)
func cbop_0x4E(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	bit := result & (1 << 1)
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// BIT 1, A
func cbop_0x4F(cpu *CPU) {
	bit := (cpu.A & (1 << 1)) >> 1
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 2, B
func cbop_0x50(cpu *CPU) {
	bit := (cpu.B & (1 << 2)) >> 2
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 2, C
func cbop_0x51(cpu *CPU) {
	bit := (cpu.C & (1 << 2)) >> 2
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 2, D
func cbop_0x52(cpu *CPU) {
	bit := (cpu.D & (1 << 2)) >> 2
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 2, E
func cbop_0x53(cpu *CPU) {
	bit := (cpu.E & (1 << 2)) >> 2
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 2, H
func cbop_0x54(cpu *CPU) {
	bit := (cpu.H & (1 << 2)) >> 2
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 2, L
func cbop_0x55(cpu *CPU) {
	bit := (cpu.L & (1 << 2)) >> 2
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 2, (HL)
func cbop_0x56(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	bit := result & (1 << 2)
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// BIT 2, A
func cbop_0x57(cpu *CPU) {
	bit := (cpu.A & (1 << 2)) >> 2
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 3, B
func cbop_0x58(cpu *CPU) {
	bit := (cpu.B & (1 << 3)) >> 3
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 3, C
func cbop_0x59(cpu *CPU) {
	bit := (cpu.C & (1 << 3)) >> 3
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 3, D
func cbop_0x5A(cpu *CPU) {
	bit := (cpu.D & (1 << 3)) >> 3
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 3, E
func cbop_0x5B(cpu *CPU) {
	bit := (cpu.E & (1 << 3)) >> 3
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 3, H
func cbop_0x5C(cpu *CPU) {
	bit := (cpu.H & (1 << 3)) >> 3
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 3, L
func cbop_0x5D(cpu *CPU) {
	bit := (cpu.L & (1 << 3)) >> 3
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 3, (HL)
func cbop_0x5E(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	bit := result & (1 << 3)
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// BIT 3, A
func cbop_0x5F(cpu *CPU) {
	bit := (cpu.A & (1 << 3)) >> 3
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 4, B
func cbop_0x60(cpu *CPU) {
	bit := (cpu.B & (1 << 4)) >> 4
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 4, C
func cbop_0x61(cpu *CPU) {
	bit := (cpu.C & (1 << 4)) >> 4
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 4, D
func cbop_0x62(cpu *CPU) {
	bit := (cpu.D & (1 << 4)) >> 4
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 4, E
func cbop_0x63(cpu *CPU) {
	bit := (cpu.E & (1 << 4)) >> 4
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 4, H
func cbop_0x64(cpu *CPU) {
	bit := (cpu.H & (1 << 4)) >> 4
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 4, L
func cbop_0x65(cpu *CPU) {
	bit := (cpu.L & (1 << 4)) >> 4
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 4, (HL)
func cbop_0x66(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	bit := result & (1 << 4)
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// BIT 4, A
func cbop_0x67(cpu *CPU) {
	bit := (cpu.A & (1 << 4)) >> 4
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 5, B
func cbop_0x68(cpu *CPU) {
	bit := (cpu.B & (1 << 5)) >> 5
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 5, C
func cbop_0x69(cpu *CPU) {
	bit := (cpu.C & (1 << 5)) >> 5
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 5, D
func cbop_0x6A(cpu *CPU) {
	bit := (cpu.D & (1 << 5)) >> 5
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 5, E
func cbop_0x6B(cpu *CPU) {
	bit := (cpu.E & (1 << 5)) >> 5
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 5, H
func cbop_0x6C(cpu *CPU) {
	bit := (cpu.H & (1 << 5)) >> 5
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 5, L
func cbop_0x6D(cpu *CPU) {
	bit := (cpu.L & (1 << 5)) >> 5
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 5, (HL)
func cbop_0x6E(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	bit := result & (1 << 5)
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// BIT 5, A
func cbop_0x6F(cpu *CPU) {
	bit := (cpu.A & (1 << 5)) >> 5
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 6, B
func cbop_0x70(cpu *CPU) {
	bit := (cpu.B & (1 << 6)) >> 6
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 6, C
func cbop_0x71(cpu *CPU) {
	bit := (cpu.C & (1 << 6)) >> 6
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 6, D
func cbop_0x72(cpu *CPU) {
	bit := (cpu.D & (1 << 6)) >> 6
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 6, E
func cbop_0x73(cpu *CPU) {
	bit := (cpu.E & (1 << 6)) >> 6
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 6, H
func cbop_0x74(cpu *CPU) {
	bit := (cpu.H & (1 << 6)) >> 6
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 6, L
func cbop_0x75(cpu *CPU) {
	bit := (cpu.L & (1 << 6)) >> 6
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 6, (HL)
func cbop_0x76(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	bit := result & (1 << 6)
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// BIT 6, A
func cbop_0x77(cpu *CPU) {
	bit := (cpu.A & (1 << 6)) >> 6
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 7, B
func cbop_0x78(cpu *CPU) {
	bit := cpu.B >> 7
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 7, C
func cbop_0x79(cpu *CPU) {
	bit := cpu.C >> 7
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 7, D
func cbop_0x7A(cpu *CPU) {
	bit := cpu.D >> 7
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 7, E
func cbop_0x7B(cpu *CPU) {
	bit := cpu.E >> 7
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 7, H 
func cbop_0x7C(cpu *CPU) {
	bit := cpu.H >> 7
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 7, L
func cbop_0x7D(cpu *CPU) {
	bit := cpu.L >> 7
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 7, (HL)
func cbop_0x7E(cpu *CPU) {
	bit := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L)) >> 7
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// BIT 7, A
func cbop_0x7F(cpu *CPU) {
	bit := cpu.A >> 7
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

// RES 0, B
func cbop_0x80(cpu *CPU) {
	cpu.B &= 0xFE
}

// RES 0, C
func cbop_0x81(cpu *CPU) {
	cpu.C &= 0xFE
}

// RES 0, D
func cbop_0x82(cpu *CPU) {
	cpu.D &= 0xFE
}

// RES 0, E
func cbop_0x83(cpu *CPU) {
	cpu.E &= 0xFE
}

// RES 0, H
func cbop_0x84(cpu *CPU) {
	cpu.H &= 0xFE
}

// RES 0, L
func cbop_0x85(cpu *CPU) {
	cpu.L &= 0xFE
}

// RES 0, (HL)
func cbop_0x86(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	result &= 0xFE
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// RES 0, A
func cbop_0x87(cpu *CPU) {
	cpu.A &= 0xFE
}

// RES 1, B
func cbop_0x88(cpu *CPU) {
	cpu.B &= 0xFD
}

// RES 1, C
func cbop_0x89(cpu *CPU) {
	cpu.C &= 0xFD
}

// RES 1, D
func cbop_0x8A(cpu *CPU) {
	cpu.D &= 0xFD
}

// RES 1, E
func cbop_0x8B(cpu *CPU) {
	cpu.E &= 0xFD
}

// RES 1, H
func cbop_0x8C(cpu *CPU) {
	cpu.H &= 0xFD
}

// RES 1, L
func cbop_0x8D(cpu *CPU) {
	cpu.L &= 0xFD
}

// RES 1, (HL)
func cbop_0x8E(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	result &= 0xFD
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// RES 1, A
func cbop_0x8F(cpu *CPU) {
	cpu.A &= 0xFD
}

// RES 2, B
func cbop_0x90(cpu *CPU) {
	cpu.B &= 0xFB
}

// RES 2, C
func cbop_0x91(cpu *CPU) {
	cpu.C &= 0xFB
}

// RES 2, D
func cbop_0x92(cpu *CPU) {
	cpu.D &= 0xFB
}

// RES 2, E
func cbop_0x93(cpu *CPU) {
	cpu.E &= 0xFB
}

// RES 2, H
func cbop_0x94(cpu *CPU) {
	cpu.H &= 0xFB
}

// RES 2, L
func cbop_0x95(cpu *CPU) {
	cpu.L &= 0xFB
}

// RES 2, (HL)
func cbop_0x96(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	result &= 0xFB
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// RES 2, A
func cbop_0x97(cpu *CPU) {
	cpu.A &= 0xFB
}

// RES 3, B
func cbop_0x98(cpu *CPU) {
	cpu.B &= 0xF7
}

// RES 3, C
func cbop_0x99(cpu *CPU) {
	cpu.C &= 0xF7
}

// RES 3, D
func cbop_0x9A(cpu *CPU) {
	cpu.D &= 0xF7
}

// RES 3, E
func cbop_0x9B(cpu *CPU) {
	cpu.E &= 0xF7
}

// RES 3, H
func cbop_0x9C(cpu *CPU) {
	cpu.H &= 0xF7
}

// RES 3, L
func cbop_0x9D(cpu *CPU) {
	cpu.L &= 0xF7
}

// RES 3, (HL)
func cbop_0x9E(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	result &= 0xF7
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// RES 3, A
func cbop_0x9F(cpu *CPU) {
	cpu.A &= 0xF7
}

// RES 4, B
func cbop_0xA0(cpu *CPU) {
	cpu.B &= 0xEF
}

// RES 4, C
func cbop_0xA1(cpu *CPU) {
	cpu.C &= 0xEF
}

// RES 4, D
func cbop_0xA2(cpu *CPU) {
	cpu.D &= 0xEF
}

// RES 4, E
func cbop_0xA3(cpu *CPU) {
	cpu.E &= 0xEF
}

// RES 4, H
func cbop_0xA4(cpu *CPU) {
	cpu.H &= 0xEF
}

// RES 4, L
func cbop_0xA5(cpu *CPU) {
	cpu.L &= 0xEF
}

// RES 4, (HL)
func cbop_0xA6(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	result &= 0xEF
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// RES 4, A
func cbop_0xA7(cpu *CPU) {
	cpu.A &= 0xEF
}

// RES 5, B
func cbop_0xA8(cpu *CPU) {
	cpu.B &= 0xDF
}

// RES 5, C
func cbop_0xA9(cpu *CPU) {
	cpu.C &= 0xDF
}

// RES 5, D
func cbop_0xAA(cpu *CPU) {
	cpu.D &= 0xDF
}

// RES 5, E
func cbop_0xAB(cpu *CPU) {
	cpu.E &= 0xDF
}

// RES 5, H
func cbop_0xAC(cpu *CPU) {
	cpu.H &= 0xDF
}

// RES 5, L
func cbop_0xAD(cpu *CPU) {
	cpu.L &= 0xDF
}

// RES 5, (HL)
func cbop_0xAE(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	result &= 0xDF
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// RES 5, A
func cbop_0xAF(cpu *CPU) {
	cpu.A &= 0xDF
}

// RES 6, B
func cbop_0xB0(cpu *CPU) {
	cpu.B &= 0xBF
}

// RES 6, C
func cbop_0xB1(cpu *CPU) {
	cpu.C &= 0xBF
}

// RES 6, D
func cbop_0xB2(cpu *CPU) {
	cpu.D &= 0xBF
}

// RES 6, E
func cbop_0xB3(cpu *CPU) {
	cpu.E &= 0xBF
}

// RES 6, H
func cbop_0xB4(cpu *CPU) {
	cpu.H &= 0xBF
}

// RES 6, L
func cbop_0xB5(cpu *CPU) {
	cpu.L &= 0xBF
}

// RES 6, (HL)
func cbop_0xB6(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	result &= 0xBF
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// RES 6, A
func cbop_0xB7(cpu *CPU) {
	cpu.A &= 0xBF
}

// RES 7, B
func cbop_0xB8(cpu *CPU) {
	cpu.B &= 0x7F
}

// RES 7, C
func cbop_0xB9(cpu *CPU) {
	cpu.C &= 0x7F
}

// RES 7, D
func cbop_0xBA(cpu *CPU) {
	cpu.D &= 0x7F
}

// RES 7, E
func cbop_0xBB(cpu *CPU) {
	cpu.E &= 0x7F
}

// RES 7, H
func cbop_0xBC(cpu *CPU) {
	cpu.H &= 0x7F
}

// RES 7, L
func cbop_0xBD(cpu *CPU) {
	cpu.L &= 0x7F
}

// RES 7, (HL)
func cbop_0xBE(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	result &= 0x7F
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// RES 7, A
func cbop_0xBF(cpu *CPU) {
	cpu.A &= 0x7F
}

// SET 0, B
func cbop_0xC0(cpu *CPU) {
	cpu.B |= 1
}

// SET 0, C
func cbop_0xC1(cpu *CPU) {
	cpu.C |= 1
}

// SET 0, D
func cbop_0xC2(cpu *CPU) {
	cpu.D |= 1
}

// SET 0, E
func cbop_0xC3(cpu *CPU) {
	cpu.E |= 1
}

// SET 0, H
func cbop_0xC4(cpu *CPU) {
	cpu.H |= 1
}

// SET 0, L
func cbop_0xC5(cpu *CPU) {
	cpu.L |= 1
}

// SET 0, (HL)
func cbop_0xC6(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	result |= 1
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// SET 0, A
func cbop_0xC7(cpu *CPU) {
	cpu.A |= 1
}

// SET 1, B
func cbop_0xC8(cpu *CPU) {
	cpu.B |= (1 << 1)
}

// SET 1, C
func cbop_0xC9(cpu *CPU) {
	cpu.C |= (1 << 1)
}

// SET 1, D
func cbop_0xCA(cpu *CPU) {
	cpu.D |= (1 << 1)
}

// SET 1, E
func cbop_0xCB(cpu *CPU) {
	cpu.E |= (1 << 1)
}

// SET 1, H
func cbop_0xCC(cpu *CPU) {
	cpu.H |= (1 << 1)
}

// SET 1, L
func cbop_0xCD(cpu *CPU) {
	cpu.L |= (1 << 1)
}

// SET 1, (HL)
func cbop_0xCE(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	result |= (1 << 1)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// SET 1, A
func cbop_0xCF(cpu *CPU) {
	cpu.A |= (1 << 1)
}

// SET 2, B
func cbop_0xD0(cpu *CPU) {
	cpu.B |= (1 << 2)
}

// SET 2, C
func cbop_0xD1(cpu *CPU) {
	cpu.C |= (1 << 2)
}

// SET 2, D
func cbop_0xD2(cpu *CPU) {
	cpu.D |= (1 << 2)
}

// SET 2, E
func cbop_0xD3(cpu *CPU) {
	cpu.E |= (1 << 2)
}

// SET 2, H
func cbop_0xD4(cpu *CPU) {
	cpu.H |= (1 << 2)
}

// SET 2, L
func cbop_0xD5(cpu *CPU) {
	cpu.L |= (1 << 2)
}

// SET 2, (HL)
func cbop_0xD6(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	result |= (1 << 2)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// SET 2, A
func cbop_0xD7(cpu *CPU) {
	cpu.A |= (1 << 2)
}

// SET 3, B
func cbop_0xD8(cpu *CPU) {
	cpu.B |= (1 << 3)
}

// SET 3, C
func cbop_0xD9(cpu *CPU) {
	cpu.C |= (1 << 3)
}

// SET 3, D
func cbop_0xDA(cpu *CPU) {
	cpu.D |= (1 << 3)
}

// SET 3, E
func cbop_0xDB(cpu *CPU) {
	cpu.E |= (1 << 3)
}

// SET 3, H
func cbop_0xDC(cpu *CPU) {
	cpu.H |= (1 << 3)
}

// SET 3, L
func cbop_0xDD(cpu *CPU) {
	cpu.L |= (1 << 3)
}

// SET 3, (HL)
func cbop_0xDE(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	result |= (1 << 3)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// SET 3, A
func cbop_0xDF(cpu *CPU) {
	cpu.A |= (1 << 3)
}

// SET 4, B
func cbop_0xE0(cpu *CPU) {
	cpu.B |= (1 << 4)
}

// SET 4, C
func cbop_0xE1(cpu *CPU) {
	cpu.C |= (1 << 4)
}

// SET 4, D
func cbop_0xE2(cpu *CPU) {
	cpu.D |= (1 << 4)
}

// SET 4, E
func cbop_0xE3(cpu *CPU) {
	cpu.E |= (1 << 4)
}

// SET 4, H
func cbop_0xE4(cpu *CPU) {
	cpu.H |= (1 << 4)
}

// SET 4, L
func cbop_0xE5(cpu *CPU) {
	cpu.L |= (1 << 4)
}

// SET 4, (HL)
func cbop_0xE6(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	result |= (1 << 4)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// SET 4, A
func cbop_0xE7(cpu *CPU) {
	cpu.A |= (1 << 4)
}

// SET 5, B
func cbop_0xE8(cpu *CPU) {
	cpu.B |= (1 << 5)
}

// SET 5, C
func cbop_0xE9(cpu *CPU) {
	cpu.C |= (1 << 5)
}

// SET 5, D
func cbop_0xEA(cpu *CPU) {
	cpu.D |= (1 << 5)
}

// SET 5, E
func cbop_0xEB(cpu *CPU) {
	cpu.E |= (1 << 5)
}

// SET 5, H
func cbop_0xEC(cpu *CPU) {
	cpu.H |= (1 << 5)
}

// SET 5, L
func cbop_0xED(cpu *CPU) {
	cpu.L |= (1 << 5)
}

// SET 5, (HL)
func cbop_0xEE(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	result |= (1 << 5)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// SET 5, A
func cbop_0xEF(cpu *CPU) {
	cpu.A |= (1 << 5)
}

// SET 6, B
func cbop_0xF0(cpu *CPU) {
	cpu.B |= (1 << 6)
}

// SET 6, C
func cbop_0xF1(cpu *CPU) {
	cpu.C |= (1 << 6)
}

// SET 6, D
func cbop_0xF2(cpu *CPU) {
	cpu.D |= (1 << 6)
}

// SET 6, E
func cbop_0xF3(cpu *CPU) {
	cpu.E |= (1 << 6)
}

// SET 6, H
func cbop_0xF4(cpu *CPU) {
	cpu.H |= (1 << 6)
}

// SET 6, L
func cbop_0xF5(cpu *CPU) {
	cpu.L |= (1 << 6)
}

// SET 6, (HL)
func cbop_0xF6(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	result |= (1 << 6)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// SET 6, A
func cbop_0xF7(cpu *CPU) {
	cpu.A |= (1 << 6)
}

// SET 7, B
func cbop_0xF8(cpu *CPU) {
	cpu.B |= (1 << 7)
}

// SET 7, C
func cbop_0xF9(cpu *CPU) {
	cpu.C |= (1 << 7)
}

// SET 7, D
func cbop_0xFA(cpu *CPU) {
	cpu.D |= (1 << 7)
}

// SET 7, E
func cbop_0xFB(cpu *CPU) {
	cpu.E |= (1 << 7)
}

// SET 7, H
func cbop_0xFC(cpu *CPU) {
	cpu.H |= (1 << 7)
}

// SET 7, L
func cbop_0xFD(cpu *CPU) {
	cpu.L |= (1 << 7)
}

// SET 7, (HL)
func cbop_0xFE(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	result |= (1 << 7)
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
}

// SET 7, A
func cbop_0xFF(cpu *CPU) {
	cpu.A |= (1 << 7)
}
var opcodes [256]Opcode = [256]Opcode {
	Opcode{4, op_0x00}, Opcode{12, op_0x01}, Opcode{8, op_0x02}, Opcode{8, op_0x03}, Opcode{4, op_0x04}, Opcode{4, op_0x05}, Opcode{8, op_0x06}, Opcode{4, op_0x07}, Opcode{20, op_0x08}, Opcode{8, op_0x09}, Opcode{8, op_0x0A}, Opcode{8, op_0x0B}, Opcode{4, op_0x0C}, Opcode{4, op_0x0D}, Opcode{8, op_0x0E}, Opcode{4, op_0x0F},
	Opcode{4, op_0x10}, Opcode{12, op_0x11}, Opcode{8, op_0x12}, Opcode{8, op_0x13}, Opcode{4, op_0x14}, Opcode{4, op_0x15}, Opcode{8, op_0x16}, Opcode{4, op_0x17}, Opcode{12, op_0x18}, Opcode{8, op_0x19}, Opcode{8, op_0x1A}, Opcode{8, op_0x1B}, Opcode{4, op_0x1C}, Opcode{4, op_0x1D}, Opcode{8, op_0x1E}, Opcode{4, op_0x1F},
	Opcode{8, op_0x20}, Opcode{12, op_0x21}, Opcode{8, op_0x22}, Opcode{8, op_0x23}, Opcode{4, op_0x24}, Opcode{4, op_0x25}, Opcode{8, op_0x26}, Opcode{4, op_0x27}, Opcode{8, op_0x28}, Opcode{8, op_0x29}, Opcode{8, op_0x2A}, Opcode{8, op_0x2B}, Opcode{4, op_0x2C}, Opcode{4, op_0x2D}, Opcode{8, op_0x2E}, Opcode{4, op_0x2F},
	Opcode{8, op_0x30}, Opcode{12, op_0x31}, Opcode{8, op_0x32}, Opcode{8, op_0x33}, Opcode{12, op_0x34}, Opcode{12, op_0x35}, Opcode{12, op_0x36}, Opcode{4, op_0x37}, Opcode{8, op_0x38}, Opcode{8, op_0x39}, Opcode{8, op_0x3A}, Opcode{8, op_0x3B}, Opcode{4, op_0x3C}, Opcode{4, op_0x3D}, Opcode{8, op_0x3E}, Opcode{4, op_0x3F},
	Opcode{4, op_0x40}, Opcode{4, op_0x41}, Opcode{4, op_0x42}, Opcode{4, op_0x43}, Opcode{4, op_0x44}, Opcode{4, op_0x45}, Opcode{8, op_0x46}, Opcode{4, op_0x47}, Opcode{4, op_0x48}, Opcode{4, op_0x49}, Opcode{4, op_0x4A}, Opcode{4, op_0x4B}, Opcode{4, op_0x4C}, Opcode{4, op_0x4D}, Opcode{8, op_0x4E}, Opcode{4, op_0x4F},
	Opcode{4, op_0x50}, Opcode{4, op_0x51}, Opcode{4, op_0x52}, Opcode{4, op_0x53}, Opcode{4, op_0x54}, Opcode{4, op_0x55}, Opcode{8, op_0x56}, Opcode{4, op_0x57}, Opcode{4, op_0x58}, Opcode{4, op_0x59}, Opcode{4, op_0x5A}, Opcode{4, op_0x5B}, Opcode{4, op_0x5C}, Opcode{4, op_0x5D}, Opcode{8, op_0x5E}, Opcode{4, op_0x5F},
	Opcode{4, op_0x60}, Opcode{4, op_0x61}, Opcode{4, op_0x62}, Opcode{4, op_0x63}, Opcode{4, op_0x64}, Opcode{4, op_0x65}, Opcode{8, op_0x66}, Opcode{4, op_0x67}, Opcode{4, op_0x68}, Opcode{4, op_0x69}, Opcode{4, op_0x6A}, Opcode{4, op_0x6B}, Opcode{4, op_0x6C}, Opcode{4, op_0x6D}, Opcode{8, op_0x6E}, Opcode{4, op_0x6F},
	Opcode{8, op_0x70}, Opcode{8, op_0x71}, Opcode{8, op_0x72}, Opcode{8, op_0x73}, Opcode{8, op_0x74}, Opcode{8, op_0x75}, Opcode{4, op_0x76}, Opcode{8, op_0x77}, Opcode{4, op_0x78}, Opcode{4, op_0x79}, Opcode{4, op_0x7A}, Opcode{4, op_0x7B}, Opcode{4, op_0x7C}, Opcode{4, op_0x7D}, Opcode{8, op_0x7E}, Opcode{4, op_0x7F},
	Opcode{4, op_0x80}, Opcode{4, op_0x81}, Opcode{4, op_0x82}, Opcode{4, op_0x83}, Opcode{4, op_0x84}, Opcode{4, op_0x85}, Opcode{8, op_0x86}, Opcode{4, op_0x87}, Opcode{4, op_0x88}, Opcode{4, op_0x89}, Opcode{4, op_0x8A}, Opcode{4, op_0x8B}, Opcode{4, op_0x8C}, Opcode{4, op_0x8D}, Opcode{8, op_0x8E}, Opcode{4, op_0x8F},
	Opcode{4, op_0x90}, Opcode{4, op_0x91}, Opcode{4, op_0x92}, Opcode{4, op_0x93}, Opcode{4, op_0x94}, Opcode{4, op_0x95}, Opcode{8, op_0x96}, Opcode{4, op_0x97}, Opcode{4, op_0x98}, Opcode{4, op_0x99}, Opcode{4, op_0x9A}, Opcode{4, op_0x9B}, Opcode{4, op_0x9C}, Opcode{4, op_0x9D}, Opcode{8, op_0x9E}, Opcode{4, op_0x9F},
	Opcode{4, op_0xA0}, Opcode{4, op_0xA1}, Opcode{4, op_0xA2}, Opcode{4, op_0xA3}, Opcode{4, op_0xA4}, Opcode{4, op_0xA5}, Opcode{8, op_0xA6}, Opcode{4, op_0xA7}, Opcode{4, op_0xA8}, Opcode{4, op_0xA9}, Opcode{4, op_0xAA}, Opcode{4, op_0xAB}, Opcode{4, op_0xAC}, Opcode{4, op_0xAD}, Opcode{8, op_0xAE}, Opcode{4, op_0xAF},
	Opcode{4, op_0xB0}, Opcode{4, op_0xB1}, Opcode{4, op_0xB2}, Opcode{4, op_0xB3}, Opcode{4, op_0xB4}, Opcode{4, op_0xB5}, Opcode{8, op_0xB6}, Opcode{4, op_0xB7}, Opcode{4, op_0xB8}, Opcode{4, op_0xB9}, Opcode{4, op_0xBA}, Opcode{4, op_0xBB}, Opcode{4, op_0xBC}, Opcode{4, op_0xBD}, Opcode{8, op_0xBE}, Opcode{4, op_0xBF},
	Opcode{8, op_0xC0}, Opcode{12, op_0xC1}, Opcode{12, op_0xC2}, Opcode{16, op_0xC3}, Opcode{12, op_0xC4}, Opcode{16, op_0xC5}, Opcode{8, op_0xC6}, Opcode{16, op_0xC7}, Opcode{8, op_0xC8}, Opcode{16, op_0xC9}, Opcode{12, op_0xCA}, Opcode{4, op_0xCB}, Opcode{12, op_0xCC}, Opcode{24, op_0xCD}, Opcode{8, op_0xCE}, Opcode{16, op_0xCF},
	Opcode{8, op_0xD0}, Opcode{12, op_0xD1}, Opcode{12, op_0xD2}, Opcode{8, op_null}, Opcode{12, op_0xD4}, Opcode{16, op_0xD5}, Opcode{8, op_0xD6}, Opcode{16, op_0xD7}, Opcode{8, op_0xD8}, Opcode{16, op_0xD9}, Opcode{12, op_0xDA}, Opcode{8, op_null}, Opcode{12, op_0xDC}, Opcode{4, op_null}, Opcode{8, op_0xDE}, Opcode{16, op_0xDF},
	Opcode{12, op_0xE0}, Opcode{12, op_0xE1}, Opcode{8, op_0xE2}, Opcode{8, op_null}, Opcode{4, op_null}, Opcode{16, op_0xE5}, Opcode{8, op_0xE6}, Opcode{16, op_0xE7}, Opcode{16, op_0xE8}, Opcode{4, op_0xE9}, Opcode{16, op_0xEA}, Opcode{8, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_0xEE}, Opcode{16, op_0xEF},
	Opcode{12, op_0xF0}, Opcode{12, op_0xF1}, Opcode{8, op_0xF2}, Opcode{4, op_0xF3}, Opcode{4, op_null}, Opcode{16, op_0xF5}, Opcode{8, op_0xF6}, Opcode{16, op_0xF7}, Opcode{12, op_0xF8}, Opcode{8, op_0xF9}, Opcode{16, op_0xFA}, Opcode{4, op_0xFB}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_0xFE}, Opcode{16, op_0xFF},

}

var cbopcodes [256]Opcode = [256]Opcode {
	Opcode{8, cbop_0x00}, Opcode{8, cbop_0x01}, Opcode{8, cbop_0x02}, Opcode{8, cbop_0x03}, Opcode{8, cbop_0x04}, Opcode{8, cbop_0x05}, Opcode{16, cbop_0x06}, Opcode{8, cbop_0x07}, Opcode{8, cbop_0x08}, Opcode{8, cbop_0x09}, Opcode{8, cbop_0x0A}, Opcode{8, cbop_0x0B}, Opcode{8, cbop_0x0C}, Opcode{8, cbop_0x0D}, Opcode{16, cbop_0x0E}, Opcode{8, cbop_0x0F},
	Opcode{8, cbop_0x10}, Opcode{8, cbop_0x11}, Opcode{8, cbop_0x12}, Opcode{8, cbop_0x13}, Opcode{8, cbop_0x14}, Opcode{8, cbop_0x15}, Opcode{16, cbop_0x16}, Opcode{8, cbop_0x17}, Opcode{8, cbop_0x18}, Opcode{8, cbop_0x19}, Opcode{8, cbop_0x1A}, Opcode{8, cbop_0x1B}, Opcode{8, cbop_0x1C}, Opcode{8, cbop_0x1D}, Opcode{16, cbop_0x1E}, Opcode{8, cbop_0x1F},
	Opcode{8, cbop_0x20}, Opcode{8, cbop_0x21}, Opcode{8, cbop_0x22}, Opcode{8, cbop_0x23}, Opcode{8, cbop_0x24}, Opcode{8, cbop_0x25}, Opcode{16, cbop_0x26}, Opcode{8, cbop_0x27}, Opcode{8, cbop_0x28}, Opcode{8, cbop_0x29}, Opcode{8, cbop_0x2A}, Opcode{8, cbop_0x2B}, Opcode{8, cbop_0x2C}, Opcode{8, cbop_0x2D}, Opcode{16, cbop_0x2E}, Opcode{8, cbop_0x2F},
	Opcode{8, cbop_0x30}, Opcode{8, cbop_0x31}, Opcode{8, cbop_0x32}, Opcode{8, cbop_0x33}, Opcode{8, cbop_0x34}, Opcode{8, cbop_0x35}, Opcode{16, cbop_0x36}, Opcode{8, cbop_0x37}, Opcode{8, cbop_0x38}, Opcode{8, cbop_0x39}, Opcode{8, cbop_0x3A}, Opcode{8, cbop_0x3B}, Opcode{8, cbop_0x3C}, Opcode{8, cbop_0x3D}, Opcode{16, cbop_0x3E}, Opcode{8, cbop_0x3F},
	Opcode{8, cbop_0x40}, Opcode{8, cbop_0x41}, Opcode{8, cbop_0x42}, Opcode{8, cbop_0x43}, Opcode{8, cbop_0x44}, Opcode{8, cbop_0x45}, Opcode{12, cbop_0x46}, Opcode{8, cbop_0x47}, Opcode{8, cbop_0x48}, Opcode{8, cbop_0x49}, Opcode{8, cbop_0x4A}, Opcode{8, cbop_0x4B}, Opcode{8, cbop_0x4C}, Opcode{8, cbop_0x4D}, Opcode{12, cbop_0x4E}, Opcode{8, cbop_0x4F},
	Opcode{8, cbop_0x50}, Opcode{8, cbop_0x51}, Opcode{8, cbop_0x52}, Opcode{8, cbop_0x53}, Opcode{8, cbop_0x54}, Opcode{8, cbop_0x55}, Opcode{12, cbop_0x56}, Opcode{8, cbop_0x57}, Opcode{8, cbop_0x58}, Opcode{8, cbop_0x59}, Opcode{8, cbop_0x5A}, Opcode{8, cbop_0x5B}, Opcode{8, cbop_0x5C}, Opcode{8, cbop_0x5D}, Opcode{12, cbop_0x5E}, Opcode{8, cbop_0x5F},
	Opcode{8, cbop_0x60}, Opcode{8, cbop_0x61}, Opcode{8, cbop_0x62}, Opcode{8, cbop_0x63}, Opcode{8, cbop_0x64}, Opcode{8, cbop_0x65}, Opcode{12, cbop_0x66}, Opcode{8, cbop_0x67}, Opcode{8, cbop_0x68}, Opcode{8, cbop_0x69}, Opcode{8, cbop_0x6A}, Opcode{8, cbop_0x6B}, Opcode{8, cbop_0x6C}, Opcode{8, cbop_0x6D}, Opcode{12, cbop_0x6E}, Opcode{8, cbop_0x6F},
	Opcode{8, cbop_0x70}, Opcode{8, cbop_0x71}, Opcode{8, cbop_0x72}, Opcode{8, cbop_0x73}, Opcode{8, cbop_0x74}, Opcode{8, cbop_0x75}, Opcode{12, cbop_0x76}, Opcode{8, cbop_0x77}, Opcode{8, cbop_0x78}, Opcode{8, cbop_0x79}, Opcode{8, cbop_0x7A}, Opcode{8, cbop_0x7B}, Opcode{8, cbop_0x7C}, Opcode{8, cbop_0x7D}, Opcode{12, cbop_0x7E}, Opcode{8, cbop_0x7F},
	Opcode{8, cbop_0x80}, Opcode{8, cbop_0x81}, Opcode{8, cbop_0x82}, Opcode{8, cbop_0x83}, Opcode{8, cbop_0x84}, Opcode{8, cbop_0x85}, Opcode{16, cbop_0x86}, Opcode{8, cbop_0x87}, Opcode{8, cbop_0x88}, Opcode{8, cbop_0x89}, Opcode{8, cbop_0x8A}, Opcode{8, cbop_0x8B}, Opcode{8, cbop_0x8C}, Opcode{8, cbop_0x8D}, Opcode{16, cbop_0x8E}, Opcode{8, cbop_0x8F},
	Opcode{8, cbop_0x90}, Opcode{8, cbop_0x91}, Opcode{8, cbop_0x92}, Opcode{8, cbop_0x93}, Opcode{8, cbop_0x94}, Opcode{8, cbop_0x95}, Opcode{16, cbop_0x96}, Opcode{8, cbop_0x97}, Opcode{8, cbop_0x98}, Opcode{8, cbop_0x99}, Opcode{8, cbop_0x9A}, Opcode{8, cbop_0x9B}, Opcode{8, cbop_0x9C}, Opcode{8, cbop_0x9D}, Opcode{16, cbop_0x9E}, Opcode{8, cbop_0x9F},
	Opcode{8, cbop_0xA0}, Opcode{8, cbop_0xA1}, Opcode{8, cbop_0xA2}, Opcode{8, cbop_0xA3}, Opcode{8, cbop_0xA4}, Opcode{8, cbop_0xA5}, Opcode{16, cbop_0xA6}, Opcode{8, cbop_0xA7}, Opcode{8, cbop_0xA8}, Opcode{8, cbop_0xA9}, Opcode{8, cbop_0xAA}, Opcode{8, cbop_0xAB}, Opcode{8, cbop_0xAC}, Opcode{8, cbop_0xAD}, Opcode{16, cbop_0xAE}, Opcode{8, cbop_0xAF},
	Opcode{8, cbop_0xB0}, Opcode{8, cbop_0xB1}, Opcode{8, cbop_0xB2}, Opcode{8, cbop_0xB3}, Opcode{8, cbop_0xB4}, Opcode{8, cbop_0xB5}, Opcode{16, cbop_0xB6}, Opcode{8, cbop_0xB7}, Opcode{8, cbop_0xB8}, Opcode{8, cbop_0xB9}, Opcode{8, cbop_0xBA}, Opcode{8, cbop_0xBB}, Opcode{8, cbop_0xBC}, Opcode{8, cbop_0xBD}, Opcode{16, cbop_0xBE}, Opcode{8, cbop_0xBF},
	Opcode{8, cbop_0xC0}, Opcode{8, cbop_0xC1}, Opcode{8, cbop_0xC2}, Opcode{8, cbop_0xC3}, Opcode{8, cbop_0xC4}, Opcode{8, cbop_0xC5}, Opcode{16, cbop_0xC6}, Opcode{8, cbop_0xC7}, Opcode{8, cbop_0xC8}, Opcode{8, cbop_0xC9}, Opcode{8, cbop_0xCA}, Opcode{8, cbop_0xCB}, Opcode{8, cbop_0xCC}, Opcode{8, cbop_0xCD}, Opcode{16, cbop_0xCE}, Opcode{8, cbop_0xCF},
	Opcode{8, cbop_0xD0}, Opcode{8, cbop_0xD1}, Opcode{8, cbop_0xD2}, Opcode{8, cbop_0xD3}, Opcode{8, cbop_0xD4}, Opcode{8, cbop_0xD5}, Opcode{16, cbop_0xD6}, Opcode{8, cbop_0xD7}, Opcode{8, cbop_0xD8}, Opcode{8, cbop_0xD9}, Opcode{8, cbop_0xDA}, Opcode{8, cbop_0xDB}, Opcode{8, cbop_0xDC}, Opcode{8, cbop_0xDD}, Opcode{16, cbop_0xDE}, Opcode{8, cbop_0xDF},
	Opcode{8, cbop_0xE0}, Opcode{8, cbop_0xE1}, Opcode{8, cbop_0xE2}, Opcode{8, cbop_0xE3}, Opcode{8, cbop_0xE4}, Opcode{8, cbop_0xE5}, Opcode{16, cbop_0xE6}, Opcode{8, cbop_0xE7}, Opcode{8, cbop_0xE8}, Opcode{8, cbop_0xE9}, Opcode{8, cbop_0xEA}, Opcode{8, cbop_0xEB}, Opcode{8, cbop_0xEC}, Opcode{8, cbop_0xED}, Opcode{16, cbop_0xEE}, Opcode{8, cbop_0xEF},
	Opcode{8, cbop_0xF0}, Opcode{8, cbop_0xF1}, Opcode{8, cbop_0xF2}, Opcode{8, cbop_0xF3}, Opcode{8, cbop_0xF4}, Opcode{8, cbop_0xF5}, Opcode{16, cbop_0xF6}, Opcode{8, cbop_0xF7}, Opcode{8, cbop_0xF8}, Opcode{8, cbop_0xF9}, Opcode{8, cbop_0xFA}, Opcode{8, cbop_0xFB}, Opcode{8, cbop_0xFC}, Opcode{8, cbop_0xFD}, Opcode{16, cbop_0xFE}, Opcode{8, cbop_0xFF},

}