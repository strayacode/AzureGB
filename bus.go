package main

import (
	"fmt"
	"os"
)

type Bus struct {
	cartridge Cartridge
	WRAM [0x2000]byte
	ppu PPU
	apu APU
	HRAM [0x80]byte
	timer Timer
	IE byte
	IF byte
	SB byte
	SC byte
}


func (bus *Bus) read(addr uint16) byte {
	switch {
		case addr >= 0x0000 && addr <= 0x7FFF:
			return bus.cartridge.ROM[addr]
		case addr >= 0x8000 && addr <= 0x9FFF:
			if bus.ppu.cpuVRAMAccess == true {
				return bus.ppu.VRAM[addr - 0x8000]
			}
			return 0
		case addr >= 0xA000 && addr <= 0xBFFF:
			return bus.cartridge.ERAM[addr - 0xA000]
		case addr >= 0xC000 && addr <= 0xDFFF:
			return bus.WRAM[addr - 0xC000]
		case addr >= 0xFF00 && addr <= 0xFF7F:
			return bus.readIO(addr)
		case addr >= 0xFF80 && addr <= 0xFFFE:
			return bus.HRAM[addr - 0xFF80]
		default:
			fmt.Println("DEBUG: non-readable memory location!", addr)
			os.Exit(3)
			return 0
	}
}

func (bus *Bus) read16(addr uint16) uint16 {
	switch {
		case addr >= 0x0000 && addr <= 0x7FFF:
			return uint16(bus.cartridge.ROM[addr + 1]) << 8 | uint16(bus.cartridge.ROM[addr])
		case addr >= 0x8000 && addr <= 0x9FFF:
			return uint16(bus.ppu.VRAM[addr + 1 - 0x8000]) << 8 | uint16(bus.ppu.VRAM[addr - 0x8000])
		case addr >= 0xA000 && addr <= 0xBFFF:
			return uint16(bus.cartridge.ERAM[addr + 1 - 0xA000]) << 8 | uint16(bus.cartridge.ERAM[addr - 0xA000])
		case addr >= 0xC000 && addr <= 0xDFFF:
			return uint16(bus.WRAM[addr + 1 - 0xC000]) << 8 | uint16(bus.WRAM[addr - 0xC000])
		case addr >= 0xFF80 && addr <= 0xFFFE:
			return uint16(bus.HRAM[addr + 1 - 0xFF80]) << 8 | uint16(bus.HRAM[addr - 0xFF80])
		default:
			fmt.Println("DEBUG: non-readable memory location!", addr)
			os.Exit(3)
			return 0
	}
}

func (bus *Bus) write(addr uint16, data byte) {
	switch {
		case addr >= 0x8000 && addr <= 0x9FFF:
			// if bus.ppu.cpuVRAMAccess == true {
			bus.ppu.VRAM[addr - 0x8000] = data
			// }
		case addr >= 0xA000 && addr <= 0xBFFF:
			bus.cartridge.ERAM[addr - 0xA000] = data
		case addr >= 0xC000 && addr <= 0xDFFF:
			bus.WRAM[addr - 0xC000] = data
		case addr >= 0xFF00 && addr <= 0xFF7F:
			bus.writeIO(addr, data)
		case addr >= 0xFF80 && addr <= 0xFFFE:
			bus.HRAM[addr - 0xFF80] = data
		case addr == 0xFFFF:
			bus.IE = data
		default:
			fmt.Println("DEBUG: non-writeable memory location!", addr)
			os.Exit(3)
	}
}

func (bus *Bus) readIO(addr uint16) byte {
	switch addr {
		case 0xFF01:
			return bus.SB
		case 0xFF02:
			return bus.SC
		case 0xFF40:
			return bus.ppu.LCDC
		case 0xFF42:
			return bus.ppu.SCY
		case 0xFF43:
			return bus.ppu.SCX
		case 0xFF44:
			return bus.ppu.LY
		case 0xFF45:
			return bus.ppu.LYC
		case 0xFF47:
			return bus.ppu.BGP
		default:
			fmt.Println(addr, "IO read not implemented yet!")
			os.Exit(3)
			return 0
	}
}

func (bus *Bus) writeIO(addr uint16, data byte) byte {
	switch addr {
		case 0xFF01:
			bus.SB = data
			// fmt.Println(data)
		case 0xFF02:
			bus.SC = data
		case 0xFF07:
			bus.timer.TAC = data
		case 0xFF0F:
			bus.IF = data
		case 0xFF11:
			bus.apu.NR11 = data
		case 0xFF12:
			bus.apu.NR12 = data
		case 0xFF13:
			bus.apu.NR13 = data
		case 0xFF14:
			bus.apu.NR14 = data
		case 0xFF24:
			bus.apu.NR50 = data
		case 0xFF25:
			bus.apu.NR51 = data
		case 0xFF26:
			bus.apu.NR52 = data
		case 0xFF40:
			bus.ppu.LCDC = data
		case 0xFF41:
			bus.ppu.LCDCSTAT = data
		case 0xFF42:
			bus.ppu.SCY = data
		case 0xFF43:
			bus.ppu.SCX = data
		case 0xFF44:
			bus.ppu.LY = data
		case 0xFF45:
			bus.ppu.LYC = data
		case 0xFF47:
			bus.ppu.BGP = data
		case 0xFF50:
			bus.cartridge.unmapBootROM()
		
		default:
			fmt.Println("IO reg not handled!", addr)
			os.Exit(3)
			return 0
	}
	return 0
}