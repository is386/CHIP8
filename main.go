package main

import (
	"fmt"

	"github.com/faiface/pixel/pixelgl"
	"github.com/is386/GoCHIP/chip8"
)

var (
	FILENAME = "ibm.ch8"
	SCALE    = 10.0
	FONT     = [80]byte{
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
	emu := chip8.NewEmulator(*screen)
	emu.LoadFont(FONT)
	for !screen.Closed() {
	}
}

func main() {
	fmt.Println(FILENAME)
	pixelgl.Run(run)
}
