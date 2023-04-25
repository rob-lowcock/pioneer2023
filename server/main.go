package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
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
	defer connection.Close()

	http.Handle("/", &handlers.LoginHandler{})
	http.Handle("/health", &handlers.HealthHandler{})

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
