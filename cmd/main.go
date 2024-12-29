package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Dostonlv/http3-go/controllers"
	_ "github.com/Dostonlv/http3-go/docs"
	"github.com/quic-go/quic-go/http3"
	httpSwagger "github.com/swaggo/http-swagger"
)

// swag init -g controllers/user_controller.go -o docs
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/register", controllers.RegisterUser)
	mux.HandleFunc("/users", controllers.GetUsers)
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	fmt.Println("Starting HTTP/3 server on https://localhost:4433")
	err := http3.ListenAndServeTLS(":4433", "certs/server.crt", "certs/server.key", mux)

	if err != nil {
		log.Fatalf("Failed to configure HTTP/3: %v", err)
	}

}
