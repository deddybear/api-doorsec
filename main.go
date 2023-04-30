package main

import (
	"api-iotdoor/app/config"
	exception "api-iotdoor/exceptions"
	"api-iotdoor/middleware"
	"api-iotdoor/routes"
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
		Addr:    "localhost:80",
		Handler: middleware.NewLogMiddleware(handler),
		//Handler: middleware.NewAuthMiddleware(handler),
	}

	fmt.Println("Server dijalankan pada port 80")
	log.Fatal(server.ListenAndServe())
}
