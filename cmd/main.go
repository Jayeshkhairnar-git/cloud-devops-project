package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	gorillahandlers "github.com/gorilla/handlers"
	"github.com/hs-heilbronn-devsecops/acetlisto/handlers"
	"github.com/hs-heilbronn-devsecops/acetlisto/stores"
	"github.com/spf13/viper"

	"github.com/hs-heilbronn-devsecops/acetlisto/internal/telemetry"
)

func main() {
    ctx := context.Background()

	shutdown, err := telemetry.InitTracer("acetlisto-service-cloudcommanders")
	if err != nil {
		log.Fatalf("failed to init tracer: %v", err)
	}
	defer shutdown(ctx)

	viper.AutomaticEnv()
	viper.SetDefault("PORT", "8080")

	store := stores.NewMemoryItemStore()
	r := handlers.New(store)

	port := viper.GetString("PORT")
	log.Printf("Server starting on :%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), gorillahandlers.LoggingHandler(os.Stdout, r)))
}
