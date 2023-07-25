package routes

import (
	controller "Practice/Go-Projects/jwt/controllers"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.Engine) {
	g := r.Group("/api")

	g.POST("/signup", controller.SignUp)
	g.POST("/login", controller.Login)

}
