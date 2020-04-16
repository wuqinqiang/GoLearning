package routes

import (
	"github.com/wuqinqiang/goblog/goblog/handlers"
	"net/http"
)

//定义一个webRoute结构体存放单个路由
type WebRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type WebRoutes []WebRoute

var webRoutes = WebRoutes{
	WebRoute{
		"Home",
		"GET",
		"/",
		handlers.Home,
	},
	WebRoute{
		"Posts",
		"GET",
		"/posts",
		handlers.GetPosts,
	},
	WebRoute{
		"User",
		"GET",
		"/user/{id}",
		handlers.GetUser,
	},
}
