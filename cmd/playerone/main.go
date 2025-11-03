package main

import (
	"log"

	httpx "github.com/kyerans/playerone/internal/presenters/http"
	"github.com/kyerans/playerone/internal/presenters/http/handlers"
	"github.com/kyerans/playerone/internal/services"
)

func main() {
	svc := services.New()
	hdl := handlers.New(svc)
	server := httpx.NewServer(hdl)

	if err := server.ListenAndServe(":8080"); err != nil {
		log.Fatal(err)
	}
}
