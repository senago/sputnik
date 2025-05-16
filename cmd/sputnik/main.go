package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/senago/sputnik/internal/gui"
	"github.com/senago/sputnik/internal/ioc"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	container, err := ioc.New(
		ctx,
		GetConfig(),
	)
	if err != nil {
		exitWithError(err)
	}

	window := gui.New(container)

	window.SetOnClosed(func() {
		log.Println("closing...")

		if err := container.Close(); err != nil {
			log.Printf("error during close: %v\n", err)
		} else {
			log.Println("closed successfully")
		}
	})

	window.ShowAndRun()
}

func exitWithError(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
