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
const TickTime = 250 * time.Millisecond

type Cell struct {
	Locked  bool
	Covered bool
}

type Grid [PlayfieldHeight][PlayfieldWidth]Cell

func main() {
	grid := Grid{
		{Cell{Covered: true}, Cell{Covered: true}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{Covered: true}, Cell{Covered: true}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
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
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
	}
	for {
		clearConsole()
		tick(&grid)
		render(grid)
		time.Sleep(TickTime)
	}
}

func clearConsole() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func tick(grid *Grid) {
	for row := PlayfieldHeight - 1; row >= 0; row-- {
		for col := 0; col < PlayfieldWidth; col++ {
			cell := &grid[row][col]

			if cell.Covered {
				if row == PlayfieldHeight-1 || (row+1 < PlayfieldHeight && grid[row+1][col].Covered && grid[row+1][col].Locked) {
					cell.Locked = true
				} else if row+1 < PlayfieldHeight && !grid[row+1][col].Covered {
					grid[row+1][col] = *cell
					grid[row][col] = Cell{}
				}
			} else if row == PlayfieldHeight {

			}
		}
	}
}

func render(grid Grid) {
	for row := 0; row < PlayfieldHeight; row++ {
		for col := 0; col < PlayfieldWidth; col++ {
			cell := grid[row][col]
			if cell.Covered {
				// if cell.Locked {
				// 	fmt.Print("LL")
				// } else {
				// 	fmt.Print("UU")
				// }
				fmt.Print(BlockSprite)
			} else {
				fmt.Print(EmptySprite)
			}
		}
		fmt.Print("\n")
	}

	fmt.Printf("Score: %d\n", 10)
}
