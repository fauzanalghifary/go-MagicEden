package main

import (
	"fmt"
	"goMagicEden/handler"
	"goMagicEden/token"
	"log"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=postgres dbname=magic-eden port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection Error")
	}

	fmt.Println("DATABASE CONNECTION SUCCESS")

	db.AutoMigrate(&token.Token{})

	tokenRepository := token.NewRepository(db)
	tokenService := token.NewService(tokenRepository)
	tokenHandler := handler.NewTokenHandler(tokenService)

	router := gin.Default()

	router.GET("/", tokenHandler.RootHandler)
	router.GET("/public/:wallet_address/contents", tokenHandler.GetWalletContents)
	router.POST("public/fetch_tokens", tokenHandler.PostWalletContents)
	router.DELETE("/public/token/delete", tokenHandler.DeleteToken)

	router.Run()

}
