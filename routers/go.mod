module github.com/WTBacon/todo-app-gin/routers

go 1.14

require (
	github.com/WTBacon/todo-app-gin/controllers v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.2
)

replace (
	github.com/WTBacon/todo-app-gin/config => ../config
	github.com/WTBacon/todo-app-gin/controllers => ../controllers
    github.com/WTBacon/todo-app-gin/models => ../models
)
