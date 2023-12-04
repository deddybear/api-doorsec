package main

import (
	"api-doorsec/app/config"
	exception "api-doorsec/exceptions"
	"api-doorsec/middleware"
	"api-doorsec/routes"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	db := config.NewDB()
	validate := validator.New()
	router := httprouter.New()

	routes.UsersRoutes(router, db, validate)
	routes.WebRoutes(router)
	router.PanicHandler = exception.ErrorHandler

	handler := cors.Default().Handler(router)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewLogMiddleware(handler),
		//Handler: middleware.NewAuthMiddleware(handler),
	}

	fmt.Println("Server dijalankan pada port 8080")
	log.Fatal(server.ListenAndServe())
}
