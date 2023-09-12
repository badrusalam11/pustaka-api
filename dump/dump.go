package dump

import (
	"fmt"
	"log"
	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DumpDb(c *gin.Context) {
	dsn :=
		"root:@tcp(127.0.0.1:3306)/pustakaapi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("db connection error")
	}
	db.AutoMigrate(book.Book{})

	book := book.Book{}
	book.Title = "Atomic habits"
	book.Price = 12000
	book.Description = "Buku self development tentang membangun kebiasaan baik"
	book.Rating = 4
	err = db.Create(&book).Error
	if err != nil {
		fmt.Println("Error creating book record")
	}

}
