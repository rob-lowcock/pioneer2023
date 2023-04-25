package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/rob-lowcock/pioneer2023/auth"
	"github.com/rob-lowcock/pioneer2023/db"
	"github.com/rob-lowcock/pioneer2023/handlers"
)

func main() {
	_ = godotenv.Load("env/.env")

	database := db.Db{}

	connection, err := database.Connect()
	if err != nil {
		log.Fatal("Database connection error", err)
	}
	defer connection.Close(context.Background())

	auth := auth.Auth{
		Db: connection,
	}

	authServer, err := auth.BuildServer()
	if err != nil {
		log.Fatal("Auth setup error", err)
	}

	http.Handle("/", &handlers.LoginHandler{})
	http.Handle("/health", &handlers.HealthHandler{})
	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := authServer.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})
	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		authServer.HandleTokenRequest(w, r)
	})

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := http.ListenAndServe("127.0.0.1:8123", nil); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	log.Print("Server started on port 8123")

	<-done
	log.Print("Server stopped")

}
