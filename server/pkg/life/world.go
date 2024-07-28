package life

import (
	"errors"
	"math/rand"
	"time"
)

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func NewWorld(height, width int) (*World, error) {
	if height <= 0 || width <= 0 {
		return nil, errors.New("new world error")
	}
	// creating new World type with (height) slices
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width) // creating new slice in every line
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}, nil
}

func (w *World) Next(x, y int) bool {
	n := w.Neighbors(x, y)       // getting number of alive neighbors
	alive := w.Cells[y][x]       // current cell state
	if n < 4 && n > 1 && alive { // if there's two or three neighbors and the cell is alive
		return true // then the next state is still alive
	}
	if n == 3 && !alive { // if the cell is dead but it has three neighbors
		return true // cell is born
	}

	return false // other cases - cell is dead
}

func (w *World) Neighbors(x, y int) int {
	height, width := w.Height, w.Width
	neighbors := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			nx, ny := x+i, y+j

			if nx >= 0 && nx < height && ny >= 0 && ny < width {
				if w.Cells[nx][ny] {
					neighbors++
				}
			}
		}
	}

	return neighbors
}

func NextState(oldWorld, newWorld *World) {
	for i := 0; i < oldWorld.Height; i++ {
		for j := 0; j < oldWorld.Width; j++ {
			// for every cell getting its new state
			newWorld.Cells[i][j] = oldWorld.Next(j, i)
		}
	}
}

// RandInit fills grid with given percent of alive cells
func (w *World) RandInit(percentage int) {
	// number of alive cells
	numAlive := percentage * w.Height * w.Width / 100
	// filling first cells with alive ones
	w.fillAlive(numAlive)
	// creating random numbers
	r := rand.New(rand.NewSource(time.Now().Unix()))

	// randomly changing places
	for i := 0; i < w.Height*w.Width; i++ {
		randRowLeft := r.Intn(w.Width)
		randColLeft := r.Intn(w.Height)
		randRowRight := r.Intn(w.Width)
		randColRight := r.Intn(w.Height)

		w.Cells[randRowLeft][randColLeft] = w.Cells[randRowRight][randColRight]
	}
}

func (w *World) fillAlive(num int) {
	aliveCount := 0
	for j, row := range w.Cells {
		for k := range row {
			w.Cells[j][k] = true
			aliveCount++
			if aliveCount == num {
				return
			}
		}
	}
}
