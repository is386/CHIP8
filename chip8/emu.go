package chip8

import (
	"fmt"
	"io/ioutil"
)

type Emulator struct {
	memory     [4096]uint16
	pc         uint16
	index      uint16
	registers  [16]uint8
	stack      Stack
	soundTimer uint8
	delayTimer uint8
	screen     Screen
}

func NewEmulator(screen Screen) *Emulator {
	emu := Emulator{pc: 0x200}
	return &emu
}

func (emu *Emulator) LoadFont(font [80]byte) {
	for i := 0; i < len(font); i++ {
		emu.memory[i] = uint16(font[i])
	}
}

func (emu *Emulator) LoadRom(filename string) {
	rom, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(rom); i++ {
		emu.memory[emu.pc+uint16(i)] = uint16(rom[i])
	}
}

func (emu *Emulator) fetch() uint16 {
	byte1 := emu.memory[emu.pc]
	byte2 := emu.memory[emu.pc+1]
	emu.pc += 2
	return (byte1 << 8) | byte2
}

func (emu *Emulator) Execute() {
	instr := emu.fetch()
	fmt.Printf("%x\n", instr)
}
