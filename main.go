package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const BlockSprite = "[]"
const EmptySprite = ".."
const PlayfieldWidth = 10
const PlayfieldHeight = 24

type Cell struct {
	Locked  bool
	Color   string
	Covered bool
}

func main() {
	grid := [PlayfieldHeight][PlayfieldWidth]Cell{
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{Covered: true}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{Covered: true}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{Covered: true}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{Covered: true}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
	}
	for {
		clearConsole()
		render(grid)
		time.Sleep(1000 * time.Millisecond)
	}
}

func clearConsole() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func render(grid [PlayfieldHeight][PlayfieldWidth]Cell) {
	for row := 0; row < PlayfieldHeight; row++ {
		for cell := 0; cell < PlayfieldWidth; cell++ {
			if grid[row][cell].Covered {
				fmt.Print(BlockSprite)
			} else {
				fmt.Print(EmptySprite)
			}
		}
		fmt.Print("\n")
	}
}
