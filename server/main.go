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
	"github.com/rob-lowcock/pioneer2023/helpers"
)

func main() {
	_ = godotenv.Load("env/.env")

	database := db.Db{}

	connection, err := database.Connect()
	if err != nil {
		log.Fatal("Database connection error", err)
	}
	defer connection.Close(context.Background())

	dbUser := db.User{
		Db: connection,
	}

	dbRetrocard := db.Retrocard{
		Db: connection,
	}

	auth := auth.Auth{
		Db:     connection,
		DbUser: dbUser,
	}

	// Handlers
	loginHandler := handlers.LoginHandler{
		Auth: auth,
	}
	healthHandler := handlers.HealthHandler{}
	getRetrocardHandler := handlers.GetRetrocardHandler{
		RetrocardDb: dbRetrocard,
	}

	middleware := helpers.Middleware{}

	http.Handle("/api/health", middleware.ContentType(&healthHandler))
	http.Handle(
		"/api/login",
		middleware.Adapt(
			&loginHandler,
			middleware.ContentType,
			middleware.Cors(http.MethodPost),
		),
	)
	http.Handle(
		"/api/retrocards",
		middleware.Adapt(
			&getRetrocardHandler,
			middleware.ContentType,
			middleware.Cors(http.MethodGet),
		),
	)

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
