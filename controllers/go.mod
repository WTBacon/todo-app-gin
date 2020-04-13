module github.com/WTBacon/todo-app-gin/controllers

go 1.14

require (
	github.com/WTBacon/todo-app-gin/models v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.2
)

replace (
    github.com/WTBacon/todo-app-gin/models => ../models
    github.com/WTBacon/todo-app-gin/config => ../config
)