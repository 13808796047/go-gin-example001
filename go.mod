module github.com/13808796047/go-gin-example

go 1.13

replace (
	github.com/13808796047/go-gin-example/conf => ./conf
	github.com/13808796047/go-gin-example/handlers => ./handlers
	github.com/13808796047/go-gin-example/middleware => ./middleware
	github.com/13808796047/go-gin-example/models => ./models
	github.com/13808796047/go-gin-example/pkg/e => ./pkg/e
	github.com/13808796047/go-gin-example/pkg/setting => ./pkg/setting
	github.com/13808796047/go-gin-example/pkg/util => ./pkg/util
	github.com/13808796047/go-gin-example/routers => ./routers
)

require (
	github.com/astaxie/beego v1.12.3 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/jinzhu/gorm v1.9.16
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/unknwon/com v1.0.1
)
