package server

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/xfiendx4life/media_tel_test/pkg/delivery"
)

type EchoServer struct {
	e echo.Echo
}

func New(del *delivery.Deliver) *EchoServer {
	es := &EchoServer{
		e: *echo.New(),
	}
	es.e.Logger.SetLevel(log.INFO)
	es.e.POST("/add", del.Add)
	es.e.GET("/graph", del.Graph)
	return es
}

func (es *EchoServer) Start(port int) error {
	es.e.HideBanner = true
	if err := es.e.Start(fmt.Sprintf(":%d", port)); err != nil && err != http.ErrServerClosed {
		es.e.Logger.Errorf("Error while serving %s", err)
		return fmt.Errorf("error while serving %w", err)
	}
	return nil
}

func (es *EchoServer) Stop(ctx context.Context, sigChan chan os.Signal) error {
	<-sigChan
	es.e.Logger.Info("Ready to stop serving")
	if err := es.e.Shutdown(ctx); err != nil {
		es.e.Logger.Errorf("Error while shutting down %s", err)
		return fmt.Errorf("error while shutting down %w", err)
	}
	return nil
}
