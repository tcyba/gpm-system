package main

import (
	"github.com/gin-gonic/gin"
	srv "test.com/project-common"
)

func main() {
	r := gin.Default()
	srv.Run(r, "webcenter", ":80")
}
