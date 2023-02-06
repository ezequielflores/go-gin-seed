package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func StartServer(config *Config, router *gin.Engine) {

	srv := &http.Server{
		Addr:         config.Server.Address,
		WriteTimeout: time.Duration(config.Server.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(config.Server.ReadTimeout) * time.Second,
		IdleTimeout:  time.Duration(config.Server.IdleTimeout) * time.Second,
		Handler:      router,
	}

	go func() {
		// conexiones de servicio
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 2)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()
	log.Println("timeout of 2 seconds.")
	log.Println("Server exiting")
}
