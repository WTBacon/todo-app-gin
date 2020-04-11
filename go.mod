module github.com/WTBacon/todo-app-gin

go 1.14

require (
	github.com/WTBacon/todo-app-gin/config v0.0.0-00010101000000-000000000000
	github.com/WTBacon/todo-app-gin/models v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/jinzhu/gorm v1.9.12
)

replace (
	github.com/WTBacon/todo-app-gin/config => ./config
	github.com/WTBacon/todo-app-gin/models => ./models
	github.com/WTBacon/todo-app-gin/routers => ./routers
)
