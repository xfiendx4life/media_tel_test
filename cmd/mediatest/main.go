package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"

	"github.com/xfiendx4life/media_tel_test/cmd/server"
	"github.com/xfiendx4life/media_tel_test/pkg/delivery"
	"github.com/xfiendx4life/media_tel_test/pkg/usecase"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("can't get port %s", err)
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	uCase := usecase.New()
	del := delivery.New(uCase)
	srvr := server.New(del)
	go func() {
		if err := srvr.Start(port); err != nil {
			log.Fatal(err)
		}
	}()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer close(sigChan)
	if err = srvr.Stop(ctx, sigChan); err != nil {
		log.Fatal(err)
	}

}
