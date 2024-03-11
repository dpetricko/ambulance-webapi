package main

import (
    "log"
    "os"
    "strings"
    
    "github.com/gin-gonic/gin"
    "github.com/dpetricko/ambulance-webapi/api"
    "github.com/dpetricko/ambulance-webapi/internal/ambulance_wl"
)

func main() {
    log.Printf("Server started")
    port := os.Getenv("AMBULANCE_API_PORT")
    if port == "" {
        port = "8080"
    }
    environment := os.Getenv("AMBULANCE_API_ENVIRONMENT")
    if !strings.EqualFold(environment, "production") { // case insensitive comparison
        gin.SetMode(gin.DebugMode)
    }
    engine := gin.New()
    engine.Use(gin.Recovery())
    // request routings
    ambulance_wl.AddRoutes(engine)
    engine.GET("/openapi", api.HandleOpenApi)
    engine.Run(":" + port)
}