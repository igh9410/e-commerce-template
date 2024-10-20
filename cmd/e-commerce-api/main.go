package main

import (
	"log/slog"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil { // Running in local, must be run on go run . in ./cmd directory
		slog.Info("No .env file found. Using OS environment variables.")
	}
	/*
	   dbConn, err := db.NewDatabase()

	   	if err != nil {
	   		log.Fatalf("Could not initialize database connection: %s", err)
	   	}

	   	routerConfig := &server.RouterConfig{
	   		UserHandler:   userHandler,
	   		ChatHandler:   chatHandler,
	   		ChatWsHandler: chatWsHandler,
	   		MatchHandler:  matchHandler,
	   		// Future handlers can be added here without changing the InitRouter signature
	   	}

	   r := server.InitRouter(routerConfig)

	   	srv := &http.Server{
	   		Addr:    ":8080",
	   		Handler: r,
	   	}

	   	go func() {
	   		// service connections
	   		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	   			log.Fatalf("listen: %s\n", err)
	   		}
	   	}()

	   log.Println("Server listening on http://localhost:8080")

	   // Wait for interrupt signal to gracefully shutdown the server with
	   // a timeout of 5 seconds.
	   quit := make(chan os.Signal, 1)
	   signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	   <-quit
	   log.Println("Shutting down server...")

	   ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	   defer cancel()

	   	if err := srv.Shutdown(ctx); err != nil {
	   		log.Fatal("Server Shutdown:", err)
	   	}

	   <-ctx.Done()
	   log.Println("Server exiting")
	*/
}
