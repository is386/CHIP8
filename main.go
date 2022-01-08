package main

import (
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/is386/GoCHIP/chip8"
)

var (
	FILENAME   = "ibm.ch8"
	SCALE      = 10.0
	EMU_DELAY  = 1 * time.Millisecond
	TIME_DELAY = 16 * time.Millisecond
	FONT       = [80]byte{
		0xF0, 0x90, 0x90, 0x90, 0xF0,
		0x20, 0x60, 0x20, 0x20, 0x70,
		0xF0, 0x10, 0xF0, 0x80, 0xF0,
		0xF0, 0x10, 0xF0, 0x10, 0xF0,
		0x90, 0x90, 0xF0, 0x10, 0x10,
		0xF0, 0x80, 0xF0, 0x10, 0xF0,
		0xF0, 0x80, 0xF0, 0x90, 0xF0,
		0xF0, 0x10, 0x20, 0x40, 0x40,
		0xF0, 0x90, 0xF0, 0x90, 0xF0,
		0xF0, 0x90, 0xF0, 0x10, 0xF0,
		0xF0, 0x90, 0xF0, 0x90, 0x90,
		0xE0, 0x90, 0xE0, 0x90, 0xE0,
		0xF0, 0x80, 0x80, 0x80, 0xF0,
		0xE0, 0x90, 0x90, 0x90, 0xE0,
		0xF0, 0x80, 0xF0, 0x80, 0xF0,
		0xF0, 0x80, 0xF0, 0x80, 0x80}
)

func run() {
	screen := chip8.NewScreen(SCALE)
	emu := chip8.NewEmulator(screen)
	emu.LoadFont(FONT)
	emu.LoadRom(FILENAME)
	for !screen.Closed() {
		time.Sleep(EMU_DELAY)
		emu.Execute()
	}
}

func main() {
	pixelgl.Run(run)
}
