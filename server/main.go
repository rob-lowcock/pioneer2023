package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/joho/godotenv"
	"github.com/rob-lowcock/pioneer2023/auth"
	"github.com/rob-lowcock/pioneer2023/db"
	"github.com/rob-lowcock/pioneer2023/handlers"
	"github.com/rob-lowcock/pioneer2023/handlers/retrocard"
	"github.com/rob-lowcock/pioneer2023/helpers"
	pg "github.com/vgarvardt/go-oauth2-pg/v4"
	"github.com/vgarvardt/go-pg-adapter/pgx4adapter"
)

func main() {
	_ = godotenv.Load("env/.env")

	database := db.Db{}

	connection, err := database.Connect()
	if err != nil {
		log.Fatal("Database connection error", err)
	}
	defer connection.Close()

	dbUser := db.User{
		Db: connection,
	}

	auth := auth.Auth{
		Db:     connection,
		DbUser: dbUser,
	}

	// Set up token management
	manager := manage.NewDefaultManager()
	adapter := pgx4adapter.NewPool(connection)
	tokenStore, _ := pg.NewTokenStore(adapter, pg.WithTokenStoreGCInterval(time.Minute))
	defer tokenStore.Close()

	clientStore, _ := pg.NewClientStore(adapter)

	manager.MapTokenStorage(tokenStore)
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	// TODO: Set up Auth Code with PKCE. For the time being we use straightforward password authorization instead
	srv.SetUserAuthorizationHandler(func(w http.ResponseWriter, r *http.Request) (string, error) {
		return "", errors.ErrAccessDenied
	})

	srv.SetPasswordAuthorizationHandler(func(ctx context.Context, clientID, username, password string) (string, error) {
		user, err := auth.ValidateCredentials(username, password)
		if err != nil {
			return "", errors.ErrAccessDenied
		}

		return user.ID, nil
	})

	dbRetrocard := db.Retrocard{
		Db: connection,
	}

	// Handlers
	healthHandler := handlers.HealthHandler{}
	getRetrocardHandler := retrocard.GetRetrocardHandler{
		RetrocardDb: dbRetrocard,
	}
	createRetrocardHandler := retrocard.CreateRetrocardHandler{
		RetrocardDb: dbRetrocard,
	}
	updateRetrocardHandler := retrocard.UpdateRetrocardHandler{
		RetrocardDb: dbRetrocard,
	}
	loginHandler := handlers.LoginHandler{
		AuthServer: srv,
	}

	middleware := helpers.Middleware{
		Manager: manager,
	}

	// Routes
	r := chi.NewRouter()
	r.Use(middleware.Cors(http.MethodGet, http.MethodPost, http.MethodPut))
	r.Handle("/api/health", middleware.ContentType(&healthHandler))
	r.Route("/api/retrocards", func(r chi.Router) {
		r.Use(middleware.Protected, middleware.ContentType)
		r.Get("/", getRetrocardHandler.ServeHTTP)
		r.Post("/", createRetrocardHandler.ServeHTTP)
		r.Put("/{id}", updateRetrocardHandler.ServeHTTP)
	})

	r.HandleFunc("/api/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})
	r.Handle("/api/token", middleware.Adapt(&loginHandler))

	// Graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := http.ListenAndServe("127.0.0.1:8123", r); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	log.Print("Server started on port 8123")

	<-done
	log.Print("Server stopped")

}

// For Katy.
