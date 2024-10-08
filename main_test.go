package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestStringToGrid_MinimalContent(t *testing.T) {
	content := `
		a   .   .
		a   .   .
		a   a   .
	`
	grid := stringToGrid(content)

	assert.Equal(t, grid[0][0].Covered, true)
	assert.Equal(t, grid[1][0].Covered, true)
	assert.Equal(t, grid[2][0].Covered, true)
	assert.Equal(t, grid[2][1].Covered, true)

	assert.Equal(t, grid[3][0].Covered, false)
	assert.Equal(t, grid[0][3].Covered, false)
	assert.Equal(t, grid[PlayfieldHeight-1][PlayfieldWidth-1].Covered, false)

	assert.Equal(t, &grid[0][0].Shape, &grid[2][1].Shape)
}

func TestStringToGrid_LockingABlockLocksTheWholeShape(t *testing.T) {
	content := `
		a   a   .
		a   .   .
		a   .   .
	`
	grid := stringToGrid(content)

	assert.Equal(t, &grid[2][0].Shape, &grid[0][1].Shape)
	assert.Equal(t, grid[2][0].Shape, grid[0][1].Shape)
	assert.Equal(t, grid[2][0].Shape.Id, grid[0][1].Shape.Id)

	grid[2][0].Shape.Locked = true
	assert.Equal(t, grid[0][1].Shape.Locked, true)
}

func TestStringToGrid_DifferentShapes(t *testing.T) {
	content := `
		a   .   .   .   b   b   .   .   .   .
		a   .   .   .   b   b   .   .   .   .
		a   a   .   .   .   .   .   .   .   .
		.   .   .   c   c   c   .   .   .   .
		.   .   .   .   c   .   .   .   .   .
	`

	grid := stringToGrid(content)

	// all cells of same shape (a) have the same shape ref
	assert.Equal(t, grid[0][0].Shape.Id, grid[1][0].Shape.Id)
	assert.Equal(t, grid[1][0].Shape.Id, grid[2][0].Shape.Id)
	assert.Equal(t, grid[2][0].Shape.Id, grid[2][1].Shape.Id)

	// different shape representations have different shape ref
	assert.NotEqual(t, grid[0][0].Shape.Id, grid[0][4].Shape.Id)
	assert.NotEqual(t, grid[0][4].Shape.Id, grid[3][4].Shape.Id)
}

func TestTick__LockL(t *testing.T) {
	content := `
		a   a   .   .   .   .   .   .   .   .
		a   .   .   .   .   .   .   .   .   .
		a   .   .   .   .   .   .   .   .   .
	`

	grid := stringToGrid(content)

	for i := 0; i <= PlayfieldHeight; i++ {
		tick(&grid)
	}
	assert.NotEqual(t, grid[PlayfieldHeight-1][0].Shape, nil)
	assert.NotEqual(t, grid[PlayfieldHeight-2][0].Shape, nil)
	assert.NotEqual(t, grid[PlayfieldHeight-3][0].Shape, nil)
	assert.NotEqual(t, grid[PlayfieldHeight-3][1].Shape, nil)

	assert.Equal(t, grid[PlayfieldHeight-1][0].Shape.Locked, true)
	assert.Equal(t, grid[PlayfieldHeight-2][0].Shape.Locked, true)
	assert.Equal(t, grid[PlayfieldHeight-3][0].Shape.Locked, true)
	assert.Equal(t, grid[PlayfieldHeight-3][1].Shape.Locked, true)

	assert.Equal(t, grid[PlayfieldHeight-1][1].Covered, false)
	assert.Equal(t, grid[PlayfieldHeight-2][2].Covered, false)
	assert.Equal(t, grid[PlayfieldHeight-1][1].Shape, nil)
	assert.Equal(t, grid[PlayfieldHeight-2][2].Shape, nil)
}
