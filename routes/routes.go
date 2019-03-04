package routes

import (
	"github.com/gorilla/mux"
	"questionnaire/controllers"
)

/*
Function name:
"Handle"

Function description:
The function provides a list of all available URLs which would use RESTful web service.
*/

func Handle(router *mux.Router) {
	router.HandleFunc("/api/factors", controllers.GetFactors).Methods("GET")
	router.HandleFunc("/api/factor", controllers.CreateFactor).Methods("POST")
	router.HandleFunc("/api/factor/{factor_id:[0-9]+}", controllers.GetFactor).Methods("GET")
	router.HandleFunc("/api/factor/{factor_id:[0-9]+}", controllers.DeleteFactor).Methods("DELETE")
	router.HandleFunc("/api/factor/{factor_id:[0-9]+}", controllers.UpdateFactor).Methods("PUT")
}
