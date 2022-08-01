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

	// tokenList := token.GetDataFromAPI()
	// fmt.Println(tokenList)

	dsn := "host=localhost user=postgres password=postgres dbname=magic-eden port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection Error")
	}

	fmt.Println("DATABASE CONNECTION SUCCESS")

	db.AutoMigrate(&token.Token{})

	tokenRepository := token.NewRepository(db)
	tokenService := token.NewService(tokenRepository)
	tokens, err := tokenService.FindAll()
	fmt.Println(tokens)

	// CREATE
	// for _, t := range tokenList {
	// 	t.CreatedAt = time.Now()
	// 	t.UpdatedAt = time.Now()
	// 	err = db.Create(&t).Error
	// 	if err != nil {
	// 		log.Fatal("Error creating token")
	// 	}
	// }

	// READ
	// var tokens []token.Token
	// err = db.Find(&tokens).Error
	// if err != nil {
	// 	fmt.Println("Error finding token")
	// }

	// DELETE
	// var deleteToken token.Token
	// err = db.Debug().Where("mint_address = ?", "8ESNUs5p8hu67byo971piyWHaXGNVvvJsoYKK2iA5JgT").Find(&deleteToken).Error
	// if err != nil {
	// 	fmt.Println("Error finding token")
	// }
	// fmt.Println(deleteToken)
	// err = db.Delete(&deleteToken).Error

	router := gin.Default()

	router.GET("/", handler.RootHandler)

	router.Run()

}

// func PostData(tokenList []Token){

// }
