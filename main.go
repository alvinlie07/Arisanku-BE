package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"Saham-BE/authentication"
	"Saham-BE/handler"
	"Saham-BE/profile"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	dsn := "host=localhost user=postgres password=alvin dbname=arisan-db port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}

	// db.AutoMigrate(authentication.User{})

	authRepo := authentication.NewRepository(db)
	authService := authentication.NewService(authRepo)
	authHandler := handler.NewHandler(authService)

	profileRepo := profile.NewRepository(db)
	profileService := profile.NewService(profileRepo)
	profileHandler := handler.ProfileHandler(profileService)
	// router.use(handleErrors)
	router := gin.Default()

	// router.SetTrustedProxies([]string{"127.0.0.1", "localhost:8080"})

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

	v1.GET("/admin/get-bank-list", profileHandler.GetBankList)
	router.Run(":8080")
}
