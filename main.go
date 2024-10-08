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
const TickTime = 100 * time.Millisecond

type Cell struct {
	Covered bool
	Shape   *Shape
}

type Shape struct {
	Id     int
	Locked bool
}

type Grid [PlayfieldHeight][PlayfieldWidth]Cell

func main() {
	s0 := Shape{
		Id:     0,
		Locked: false,
	}

	s1 := Shape{
		Id:     1,
		Locked: false,
	}

	s2 := Shape{
		Id:     2,
		Locked: false,
	}

	grid := Grid{
		{Cell{Covered: true, Shape: &s0}, Cell{Covered: true, Shape: &s0}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{Covered: true, Shape: &s0}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{Covered: true, Shape: &s0}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{Covered: true, Shape: &s2}, Cell{}, Cell{Covered: true, Shape: &s2}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{Covered: true, Shape: &s2}, Cell{Covered: true, Shape: &s2}, Cell{Covered: true, Shape: &s2}, Cell{Covered: true, Shape: &s2}, Cell{Covered: true, Shape: &s2}, Cell{Covered: true, Shape: &s2}, Cell{Covered: true, Shape: &s2}, Cell{Covered: true, Shape: &s2}, Cell{Covered: true, Shape: &s2}, Cell{Covered: true, Shape: &s2}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}, Cell{}},
		{Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}},
		{Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}},
		{Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}, Cell{Covered: true, Shape: &s1}},
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
		clearBottom(&grid)
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

			if !cell.Covered {
				continue
			}

			cellBellowIsCoveredAndLocked := row+1 < PlayfieldHeight &&
				grid[row+1][col].Covered &&
				grid[row+1][col].Shape.Locked

			if row == PlayfieldHeight-1 || cellBellowIsCoveredAndLocked {
				cell.Shape.Locked = true
			} else {
				cellBellowIsNotCovered := row+1 < PlayfieldHeight &&
					!grid[row+1][col].Covered

				if !cell.Shape.Locked && cellBellowIsNotCovered {
					grid[row+1][col] = *cell
					grid[row][col] = Cell{}
				}
			}
		}
	}
}

func clearBottom(grid *Grid) {
	for row := PlayfieldHeight - 1; row >= 0; row-- {
		fullLine := true

		for col := 0; col < PlayfieldWidth; col++ {
			cell := &grid[row][col]

			if !cell.Covered {
				fullLine = false
				continue
			}

			fullLine = fullLine && cell.Covered && cell.Shape.Locked
		}

		if !fullLine {
			continue
		}

		for col := 0; col < PlayfieldWidth; col++ {
			grid[row][col] = Cell{}
			if row-1 >= 0 {
				cellAbove := grid[row-1][col]
				if cellAbove.Covered && cellAbove.Shape.Locked {
					cellAbove.Shape.Locked = false
				}
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
