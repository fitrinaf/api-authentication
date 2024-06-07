package Router

import (
	"apotek-paten/Controller"
	"apotek-paten/Middleware"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ServeApps() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	authRoutes := router.Group("/auth")
	{
		AuthRoutes(authRoutes)
	}

	userRoutes := router.Group("/user")
	{
		UserRoutes(userRoutes)
	}

	router.Run(":8080")
	fmt.Println("Server is running on port 8080")
}

func AuthRoutes(router *gin.RouterGroup) {
	router.POST("/register", Controller.Register)
	router.POST("/login", Controller.Login)
}

func UserRoutes(router *gin.RouterGroup) {
	//add protection middleware
	router.POST("/change-user", Middleware.AdminMiddleware(), Controller.ChangeUser)
	router.GET("/get-user", Middleware.AdminMiddleware(), Controller.GetUser)
}
