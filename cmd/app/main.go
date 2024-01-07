package main

import (
	"fmt"
	"omg/intermal/delivery/http"
	"os"

	"omg/pkg/env"
)

func main() {
	env.LoadEnvFile()

	h := http.NewHandler()

	srv := http.CreateHTTPServer(os.Getenv("PORT"), h.InitRoutes())
	if err := srv.Run(); err != nil {
		fmt.Printf("ERROR: %s", err.Error())
	}
}
