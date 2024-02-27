package router

import (
    "github.com/gin-gonic/gin"
)

func Initialize () *gin.Engine {
    r := gin.Default()
    return r
}
