package main

import (
	"fmt"
	"log"
	"net/http"

	restHandler "github.com/vitormatzembacher/desafio_clean_architecture_fullcycle/internal/interfaces/rest"
)

func main() {
	http.HandleFunc("/order", restHandler.ListOrdersHandler)
	fmt.Println("REST server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
