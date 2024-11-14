package main

import (
	"context"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/telegram/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(),
		os.Interrupt, os.Kill, syscall.SIGQUIT, syscall.SIGTERM)
	defer cancel()
	a := app.New()
	a.Run(ctx)
}
