package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	app, err := NewApp(
		ctx,
		GetConfig(),
	)
	if err != nil {
		exitWithError(err)
	}

	defer func() {
		log.Println("closing...")

		if err := app.Close(); err != nil {
			log.Printf("error during close: %v\n", err)
		} else {
			log.Println("closed successfully")
		}
	}()

	startApp(app)
}

func exitWithError(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
