package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/senago/sputnik/internal/gui"
	"github.com/senago/sputnik/internal/ioc"
)

var (
	pgDSN string
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	flag.StringVar(&pgDSN, "pg_dsn", "", "postgres dsn string")
	flag.Parse()

	container, err := ioc.New(
		ctx,
	)
	if err != nil {
		exitWithError(err)
	}

	if pgDSN != "" {
		if err := container.ConnectDB(ctx, pgDSN); err != nil {
			exitWithError(fmt.Errorf(
				"failed to connect to [%v]: %v",
				pgDSN,
				err,
			))
		}
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
	_, _ = fmt.Println(err)
	os.Exit(1)
}
