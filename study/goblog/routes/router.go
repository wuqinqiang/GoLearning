package routes

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router:=mux.NewRouter().StrictSlash(true)
	//遍历web.go 中定义的所有 webRouters
	for _,route:=range webRoutes{
		//将每个web路由应用到路由器上
		router.Methods(route.Method).Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
