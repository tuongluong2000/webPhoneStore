package main

import (
	"fmt"
	"net/http"

	"github.com/tuongluong2000/webPhoneStore/back-end/config"
	"github.com/tuongluong2000/webPhoneStore/back-end/routes"
)

func main() {
	cfg := config.LoadConfig()
	routes.SetupRoutes()

	fmt.Println("Server running :" + cfg.Port)
	http.ListenAndServe(":"+cfg.Port, nil)
}
