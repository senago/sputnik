package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/senago/sputnik/internal/ioc"
	"github.com/senago/sputnik/tests/gen"
)

const addr = ":8888"

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

	httpHandler := &HTTPHandler{
		container: container,
	}

	http.HandleFunc("/create_orbit", respond500OnErr(httpHandler.CreateOrbit))
	http.HandleFunc("/get_orbits", respond500OnErr(httpHandler.GetOrbits))

	http.HandleFunc("/create_satellite", respond500OnErr(httpHandler.CreateSatellite))
	http.HandleFunc("/get_satellites", respond500OnErr(httpHandler.GetSatellites))

	server := &http.Server{Addr: addr, Handler: nil}
	go func() {
		log.Printf("running on [%s]\n", addr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("finished serving: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown failed: %v", err)
	}
}

type HTTPHandler struct {
	container *ioc.Container
}

func (h *HTTPHandler) CreateOrbit(w http.ResponseWriter, r *http.Request) error {
	insertObit := h.container.PortInsertOrbit()

	return insertObit(context.Background(), gen.RandOrbit())
}

func (h *HTTPHandler) GetOrbits(w http.ResponseWriter, r *http.Request) error {
	getOrbits := h.container.PortGetOrbits()

	_, err := getOrbits(context.Background())
	return err
}

func (h *HTTPHandler) CreateSatellite(w http.ResponseWriter, r *http.Request) error {
	insertSatellite := h.container.PortInsertSatellite()

	return insertSatellite(context.Background(), gen.RandSatellite())
}

func (h *HTTPHandler) GetSatellites(w http.ResponseWriter, r *http.Request) error {
	getSatellites := h.container.PortGetSatellites()

	_, err := getSatellites(context.Background())
	return err
}

func respond500OnErr(
	wrappedFunc func(http.ResponseWriter, *http.Request) error,
) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := wrappedFunc(w, r)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func exitWithError(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
