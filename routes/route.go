package routes

import (
	"github.com/RichSvK/Go_WebServer/services"
	"github.com/julienschmidt/httprouter"
)

func SetupRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", services.RootHandler)
	router.GET("/student/:NIM", services.GetStudentInfo)
	router.GET("/students", services.GetStudents)
	router.DELETE("/students", services.DeleteStudents)
	router.POST("/students", services.PostStudents)
	return router
}
