package main

import (
	"context"
	"fmt"
	"os"

	"git.woa.com/mojoma/sharpedge/cmd/app"
)

var (
	ctx = context.Background()
)

func main() {
	cmd := app.NewServerCommand(ctx)
	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
