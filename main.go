package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const BlockSprite = "[]"
const EmptySprite = ".."
const PlayfieldWidth = 10
const PlayfieldHeight = 24
const DefaultTickTime = 150 * time.Millisecond

type Cell struct {
	Covered bool
	Shape   *Shape
}

type Shape struct {
	Id     int
	Locked bool
	Sprite string
}

type Grid [PlayfieldHeight][PlayfieldWidth]Cell

func main() {
	content, _ := os.ReadFile("inputgrid.txt")
	grid := stringToGrid(string(content))

	tickTime := DefaultTickTime

	godotenv.Load()
	tickTimeValue, err := strconv.Atoi(os.Getenv("TICK_TIME"))

	if err == nil {
		tickTime = time.Duration(tickTimeValue) * time.Millisecond
	}

	for {
		clearConsole()
		clearBottom(&grid)
		tick(&grid)
		render(grid)
		time.Sleep(tickTime)
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
			}
		}

		for col := 0; col < PlayfieldWidth; col++ {
			cell := &grid[row][col]

			if cell.Covered && !cell.Shape.Locked {
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
				if cell.Shape.Locked {
					fmt.Print("L" + cell.Shape.Sprite)
				} else {

					fmt.Print("U" + cell.Shape.Sprite)
				}
				// fmt.Print(BlockSprite)
			} else {
				fmt.Print(EmptySprite)
			}
		}
		fmt.Print("\n")
	}

	fmt.Printf("\nScore: %d\n", 10)
}

func stringToGrid(content string) Grid {
	var grid Grid

	var shapes []Shape
	lines := strings.Split(content, "\n")

	row := 0
	for i := 0; i < len(lines) && row < PlayfieldHeight; i++ {
		line := strings.Fields(lines[i])

		if len(line) == 0 {
			continue
		}

		var gridRow [PlayfieldWidth]Cell

		for col := 0; col < len(line); col++ {
			char := string(line[col])

			if char == "." {
				gridRow[col] = Cell{Covered: false}
				continue
			}

			id := int(rune(char[0]))

			foundIdx := -1
			for k := 0; k < len(shapes); k++ {
				if shapes[k].Id == id {
					foundIdx = k
					break
				}
			}

			if foundIdx == -1 {
				newShape := Shape{Id: id, Locked: false}
				newShape.Sprite = string(char)
				shapes = append(shapes, newShape)
				foundIdx = len(shapes) - 1
			}

			gridRow[col] = Cell{Covered: true, Shape: &shapes[foundIdx]}
		}

		grid[row] = gridRow
		row++
	}

	return grid
}
