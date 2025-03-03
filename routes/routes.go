package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leosantosw/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/students", controllers.AllStudents)
	r.Run()
}
