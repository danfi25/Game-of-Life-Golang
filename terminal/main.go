package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func NewWorld(height, width int) *World {
	// new world init
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}

	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}
}

func (w *World) Seed() {
	// randomly filling grid with alive cells
	for _, row := range w.Cells {
		for i := range row {
			if rand.Intn(10) == 1 {
				row[i] = true
			}
		}
	}
}

func (w *World) Next(x, y int) bool {
	// analyzing cell neighbors and sending the next state (alive/dead)
	n := w.Neighbours(x, y)
	alive := w.Cells[y][x]

	if n < 4 && n > 1 && alive {
		return true
	}

	if n == 3 && !alive {
		return true
	}

	return false
}

func NextState(oldWorld, newWorld *World) {
	// Updating grid cells
	for i := 0; i < oldWorld.Height; i++ {
		for j := 0; j < oldWorld.Width; j++ {
			newWorld.Cells[i][j] = oldWorld.Next(j, i)
		}
	}
}

func (w *World) SaveState(filename string) error {
	// saving the game state in dedicated file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := range w.Cells {
		arr, end := []string{}, "\n"

		for j := range w.Cells[i] {
			if w.Cells[i][j] == true {
				arr = append(arr, "1")
			} else {
				arr = append(arr, "0")
			}
		}

		row := strings.Join(arr, "")

		if i == len(w.Cells)-1 {
			end = ""
		}

		fmt.Fprint(writer, row+end)
	}

	return nil
}

func (w *World) LoadState(filename string) error {
	// loading game state from file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	new_cells := [][]bool{}

	file_scanner := bufio.NewScanner(file)
	for file_scanner.Scan() {
		col := []bool{}
		for _, let := range file_scanner.Text() {
			if string(let) == "1" {
				col = append(col, true)
			} else {
				col = append(col, false)
			}
		}

		new_cells = append(new_cells, col)

		index := len(new_cells) - 1
		if index > 0 && len(new_cells[index]) != len(new_cells[index-1]) {
			return fmt.Errorf("Different count")
		}
	}

	w.Cells = new_cells
	w.Height, w.Width = len(new_cells), len(new_cells[0])

	return nil
}

func (w *World) String() string {
	// grid showcase in terminal
	var result string

	brownSquare := "\xF0\x9F\x9F\xAB"
	greenSquare := "\xF0\x9F\x9F\xA9"

	for i := range w.Cells {
		for _, col := range w.Cells[i] {
			if col {
				result += greenSquare
			} else {
				result += brownSquare
			}
		}
		result += "\n"
	}

	return result
}

func (w *World) Neighbours(x, y int) int {
	// number of cell's alive neighbors
	n_count := 0
	positions := [][]int{
		{-1, -1}, {-1, 0},
		{-1, 1}, {0, -1},
		{0, 1}, {1, -1},
		{1, 0}, {1, 1},
	}

	for _, pos := range positions {
		new_x := (x + pos[0] + w.Width) % w.Width
		new_y := (y + pos[1] + w.Height) % w.Height

		if w.Cells[new_y][new_x] {
			n_count++
		}
	}

	return n_count
}

func main() {
	height, width := 10, 10
	current_world := NewWorld(height, width) // world init
	next_world := NewWorld(height, width)    // world for writing init

	current_world.Seed()
	for {
		fmt.Println(current_world)

		NextState(current_world, next_world)
		current_world = next_world

		time.Sleep(100 * time.Millisecond)
	}
}
