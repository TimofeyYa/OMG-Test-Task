package main

import (
	"context"
	"log"
	"omg/intermal/delivery/http"
	"omg/intermal/models"
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

	tokens := models.TokenPair{
		User:   os.Getenv("USER_TOKEN"),
		Bearer: os.Getenv("AUTH_TOKEN"),
	}

	r := repository.NewRepository(os.Getenv("YCLIENTS_URI"), pgConn, &tokens)
	s := service.NewService(r)
	h := http.NewHandler(s)

	srv := http.CreateHTTPServer(os.Getenv("PORT"), h.InitRoutes())
	go func() {
		if err := srv.Run(); err != nil {
			log.Fatalf("ERROR: %s\n", err.Error())
		}
	}()

	// Job Start
	go s.ResolveTasks(ctx)

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("ERROR: %s\n", err.Error())
	}
}
