package main

import (
    "fmt"
    "log"
    "net/http"

    "go_api_destya/config"
    "go_api_destya/routes"

    "github.com/gorilla/mux"
)

func main() {
    config.LoadConfig()
    config.ConnectDB()

    r := mux.NewRouter()

    
    routes.RouteIndex(r)
    routes.BookRoutes(r)
    routes.AuthorRoutes(r)

    log.Println("Server running on port", config.ENV.PORT)
    http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), r)
}
