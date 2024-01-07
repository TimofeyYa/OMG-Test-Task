package main

import (
	"context"
	"fmt"
	"omg/intermal/delivery/http"
	"omg/intermal/repository"
	"omg/intermal/service"
	"omg/pkg/env"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	env.LoadEnvFile()

	quit := make(chan os.Signal, 1)
	ctx := context.Background()

	r := repository.NewRepository()
	s := service.NewService(r)
	h := http.NewHandler(s)

	srv := http.CreateHTTPServer(os.Getenv("PORT"), h.InitRoutes())
	go func() {
		if err := srv.Run(); err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			quit <- syscall.SIGPIPE
		}
	}()

	// Graceful Shutdown
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
}
