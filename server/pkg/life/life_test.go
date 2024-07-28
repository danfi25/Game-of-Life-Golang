package life_test

import (
	"game/server/pkg/life"
	"testing"
)

func TestNewWorld(t *testing.T) {
	// grid settings
	height := 10
	width := 4
	// calling the function to test
	world, _ := life.NewWorld(height, width)
	// checking grid height
	if world.Height != height {
		t.Errorf("expected height: %d, actual height: %d", height, world.Height)
	}
	// checking grid width
	if world.Width != width {
		t.Errorf("expected width: %d, actual width: %d", width, world.Width)
	}
	if len(world.Cells) != height {
		t.Errorf("expected height: %d, actual number of rows: %d", height, len(world.Cells))
	}
	for i, row := range world.Cells {
		if len(row) != width {
			t.Errorf("expected width: %d, actual row's %d len: %d", width, i, world.Width)
		}
	}
}
