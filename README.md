# project

create new repository named project

check add README.md

add .gitignore file

create repository

~$ cd Desktop/

~/Desktop$ git clone https://github.com/fularani/project.git

~/Desktop$ cd project/

~/Desktop/project$ go mod init project
go: creating new go.mod: module project

create a file called example.go:

# program

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

~/Desktop/project$ go get -u github.com/gin-gonic/gin 

~/Desktop/project$ go run example.go

- [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

-[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env: export GIN_MODE=release
 - using code: gin.SetMode(gin.ReleaseMode)

-[GIN-debug] GET    /ping                     --> main.main.func1 (3 handlers)
-[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
-[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
-[GIN-debug] Listening and serving HTTP on :8080
-[GIN] 2022/09/27 - 15:45:51 | 200 |      49.107µs |             ::1 | GET      "/ping"
