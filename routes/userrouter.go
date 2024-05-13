package routes

import(
	"github.com/gin-gonic/gin"
	"golang-restaurant-management/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine){
     incomingRoutes.GET("/users",controllers.GetUsers())
	 incomingRoutes.GET("/users/:user_id",controllers.GetUser())
	 incomingRoutes.GET("/users/signup",controllers.SignUsers())
     incomingRoutes.GET("/users/login",controllers.login())

}