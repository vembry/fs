package main

import (
	"api/internal/app"
	"api/pkg/db"
	_ "embed"
	"log"
	"os"
	"os/signal"
	"syscall"
)

//go:embed sqlc/schema.sql
var sqliteDdlSchema string

func main() {
	// initiate dependencies
	appDb := app.NewDb(sqliteDdlSchema)
	defer appDb.Close()

	sqlcClient := db.New(appDb)

	// construct http-server
	hs := newHttpServer(sqlcClient)

	log.Printf("starting up server...")
	hs.Start()

	// awaits for termination
	WatchForExitSignal()

	log.Printf("shutting down server...")
	hs.Stop()

}

// WatchForExitSignal is to awaits incoming interrupt signal
// sent to the service
func WatchForExitSignal() os.Signal {
	log.Printf("awaiting sigterm...")
	ch := make(chan os.Signal, 4)
	signal.Notify(
		ch,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
		syscall.SIGTSTP,
	)

	return <-ch
}
