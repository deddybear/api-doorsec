package routes

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func WebRoutes(router *httprouter.Router) {

	router.GET("/", func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		fmt.Fprint(res, "Built by golang\n")
	})

}
