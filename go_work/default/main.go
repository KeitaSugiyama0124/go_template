package main

import (
    "github.com/joho/godotenv"
    
    "default/config/routes"
)

func main() {

    // Load .env
    godotenv.Load()

    // Routing
    routes.Init()
}

