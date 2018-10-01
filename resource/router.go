package resource

import (
	"github.com/gorilla/mux"
	"github.com/yizhiheng/golang-rest-service/logger"
)

var controller = NewController(Dao{})

//NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	logger.Logger(controller.ListResources, "ListResources")
	router.Methods("GET").Path("/resources").Name("ListResources").HandlerFunc(controller.ListResources)

	logger.Logger(controller.CreateResource, "CreateResource")
	router.Methods("POST").Path("/resources").Name("CreateResource").HandlerFunc(controller.CreateResource)

	logger.Logger(controller.UpdateResource, "UpdateResource")
	router.Methods("PUT").Path("/resources").Name("UpdateResource").HandlerFunc(controller.UpdateResource)

	logger.Logger(controller.DeleteResource, "DeleteResource")
	router.Methods("DELETE").Path("/resources/{id}").Name("DeleteResource").HandlerFunc(controller.DeleteResource)

	return router
}
