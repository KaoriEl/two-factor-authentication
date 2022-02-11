package server

import (
	"context"
	"github.com/fatih/color"
	"log"
	"main/internal/server/routes"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Serve() {
	var wait time.Duration
	color.New(color.FgHiWhite).Add(color.Underline).Println("Server Tuning... ")
	router := routes.NewRouter()
	color.New(color.FgHiWhite).Add(color.Underline).Println("Start server. Port:3003 ")

	srv := &http.Server{
		Handler: router,
		Addr:    ":3003",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("RIP Server Shutdown")
	os.Exit(0)

}
