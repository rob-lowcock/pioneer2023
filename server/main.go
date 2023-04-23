package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rob-lowcock/pioneer2023/handlers"
)

func main() {
	http.Handle("/", &handlers.LoginHandler{})

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
