package main

import (
	"Saham-BE/authentication"
	"Saham-BE/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=alvin dbname=saham-db port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}

	// db.AutoMigrate(authentication.User{})

	authRepo := authentication.NewRepository(db)
	authService := authentication.NewService(authRepo)
	authHandler := handler.NewHandler(authService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.POST("/user/login", authHandler.LoginHandler)
	v1.POST("/user/create", authHandler.RegisterHandler)
	v1.POST("/user/forget-password", authHandler.ForgetPasswordHandler)
	v1.POST("/user/reset-password", authHandler.ResetPasswordHandler)

	v1.GET("/admin/list", authHandler.FindAll)
	v1.GET("/admin/list/:id", authHandler.FindById)
	v1.POST("/admin/create-user", authHandler.RegisterHandler)
	v1.POST("/admin/update-user", authHandler.UpdateHandler)
	v1.POST("/admin/delete-user", authHandler.DeleteHandler)

	router.Run()
}
