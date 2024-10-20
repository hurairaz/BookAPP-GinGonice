package main

import (
	"gin-gonic/auth"
	"gin-gonic/config"
	"gin-gonic/controllers"
	"gin-gonic/models"
	"gin-gonic/services"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db := config.InitDB()
	er := db.AutoMigrate(&models.User{}, &models.Book{})
	if er != nil {
		log.Fatal("Migration problem")
	}
	router := gin.Default()

	userService := services.NewUserService(db)
	userController := controllers.NewUserController(userService)
	bookService := services.NewBookService(db)
	bookController := controllers.NewBookController(bookService)

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/signup", userController.CreateUser)
		userRoutes.POST("/login", userController.LoginUser)
	}
	bookRoutes := router.Group("/books")
	{
		bookRoutes.Use(auth.AuthenticateJWT())
		bookRoutes.POST("/", bookController.CreateBook)
		bookRoutes.PATCH("/", bookController.UpdateBook)
		bookRoutes.DELETE("/:book_id", bookController.DeleteBook)
		bookRoutes.GET("/:book_id", bookController.GetBookByID)
		bookRoutes.GET("/author", bookController.GetBooksByAuthor)
		bookRoutes.GET("/title", bookController.GetBooksByTitle)
	}
	_ = router.Run(":8080")
}
