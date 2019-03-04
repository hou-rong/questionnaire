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

	router.HandleFunc("/api/widgets_types", controllers.GetWidgetsTypes).Methods("GET")
	router.HandleFunc("/api/widget_type", controllers.CreateWidgetType).Methods("POST")
	router.HandleFunc("/api/widget_type/{widget_type:[0-9]+}", controllers.GetWidgetType).Methods("GET")
	router.HandleFunc("/api/widget_type/{widget_type:[0-9]+}", controllers.DeleteWidgetType).Methods("DELETE")
	router.HandleFunc("/api/widget_type/{widget_type:[0-9]+}", controllers.UpdateWidgetType).Methods("PUT")
}
