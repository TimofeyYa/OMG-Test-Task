package main

import (
	"context"
	"log"
	"omg/intermal/delivery/http"
	"omg/intermal/repository"
	"omg/intermal/service"
	"omg/pkg/clients"
	"omg/pkg/env"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	env.LoadEnvFile()
	ctx := context.Background()

	dbConf := clients.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DataBase: os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	pgConn, err := clients.InitPostgresClient(dbConf, 5)
	if err != nil {
		log.Fatalln(err.Error())
	}

	r := repository.NewRepository(os.Getenv("YCLIENTS_URI"), pgConn)
	s := service.NewService(r)
	h := http.NewHandler(s)

	srv := http.CreateHTTPServer(os.Getenv("PORT"), h.InitRoutes())
	go func() {
		if err := srv.Run(); err != nil {
			log.Fatalf("ERROR: %s\n", err.Error())
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("ERROR: %s\n", err.Error())
	}
}
