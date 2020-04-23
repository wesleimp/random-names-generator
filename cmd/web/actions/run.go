package actions

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"github.com/wesleimp/random-names-generator/cmd/web/routes"
)

// Run starts the server
func Run(c *cli.Context) error {
	addr := c.String("addr")

	r := mux.NewRouter()
	routes.Setup(r)

	srv := http.Server{
		Addr:    addr,
		Handler: r,
	}

	return serving(&srv)
}

func serving(srv *http.Server) error {
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println(errors.Wrapf(err, "Unable to listen and serve"))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println(errors.Wrapf(err, "Unable to shutdown"))
	}
	log.Println("Server exiting")

	return nil
}
