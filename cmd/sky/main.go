package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"sky/internal/cli"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	if err := cli.New().ExecuteContext(ctx); err != nil {
		fmt.Fprintln(os.Stderr, "sky:", err)
		os.Exit(1)
	}
}
