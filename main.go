package main

import (
    "github.com/gin-gonic/gin"
    "go-api-spreadsheet/controllers"
)

func main() {
    r:= gin.Default()

    r.POST("api/pengusahaan", controllers.PostData)

	r.Run(":8888")
}
