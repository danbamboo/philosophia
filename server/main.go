package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/philosophia/server/api"
)

const port = ":5070"

func main() {
	fmt.Println("hi")

	router := http.NewServeMux()
	router.HandleFunc("/healthcheck", api.Healthcheck)
	log.Printf("Starting simple-go-server on port %s", port)
	a

	handler := cors.Default().Handler(router)

	http.ListenAndServe(port, handler)
	return
}
