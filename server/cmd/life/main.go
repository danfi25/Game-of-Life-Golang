package main

import (
	"context"
	"game/server/internal/application"
	"os"
)

func main() {
	ctx := context.Background()
	// Exit with given err_code
	os.Exit(mainWithExitCode(ctx))
}

func mainWithExitCode(ctx context.Context) int {
	cfg := application.Config{
		Width:  100,
		Height: 100,
	}
	app := application.New(cfg)

	return app.Run(ctx)
}
