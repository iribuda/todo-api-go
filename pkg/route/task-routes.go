package route

import(
	"github.com/gorilla/mux"
)

var RegisterTaskStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/task/", controllers.C)
}