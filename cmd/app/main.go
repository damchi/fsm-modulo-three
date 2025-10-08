package main

import (
	_ "fsm-modulo-three/doc"
	"fsm-modulo-three/internal/adapters/http"
	"log"

	"github.com/gin-gonic/gin"
)

// @title fsm_modulo API
// @version 1.0
// @description API for computing binary modulo using FSM
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()
	http.RegisterRoutes(r)

	log.Printf("Starting server on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}

}
