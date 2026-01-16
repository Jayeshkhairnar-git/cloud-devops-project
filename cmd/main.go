package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	gorillahandlers "github.com/gorilla/handlers"
	"github.com/hs-heilbronn-devsecops/acetlisto/handlers"
	"github.com/hs-heilbronn-devsecops/acetlisto/stores"
	"github.com/spf13/viper"
)

func main() {

	ctx := context.Background()

	shutdown := initTracer(ctx)

	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Printf("failed to shutdown tracer: %v", err)
		}
	}()

	viper.AutomaticEnv()
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("READ_TIMEOUT", "30")

	store := stores.NewMemoryItemStore()

	r := handlers.New(store)

	port := viper.GetString("PORT")
	log.Printf("Server starting on :%s", port)

	server := &http.Server{
		Addr:        fmt.Sprintf(":%s", port),
		Handler:     gorillahandlers.LoggingHandler(os.Stdout, r),
		ReadTimeout: viper.GetDuration("READ_TIMEOUT") * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
