package main

import (
	"pustaka-api/dump"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	// dsn :=
	// 	"root:@tcp(127.0.0.1:3306)/pustakaapi?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("db connection error")
	// }
	// db.AutoMigrate(book.Book{})

	// book := book.Book{}
	// // Create
	// // book.Title = "Atomic habits"
	// // book.Price = 12000
	// // book.Description = "Buku self development tentang membangun kebiasaan baik"
	// // book.Rating = 4
	// // err = db.Create(&book).Error
	// // if err != nil {
	// // 	fmt.Println("Error creating book record")
	// // }

	// // Select
	// // err = db.Debug().First(&book).Error
	// err = db.Debug().Last(&book).Error
	// // var books []book.Book
	// // err = db.Debug().Where("title=?", "Atomic Habits").Find(&book).Error
	// if err != nil {
	// 	fmt.Println("error finding book record")
	// }
	// fmt.Println("book object %v", book)

	// //Update
	// book.Title = "Atomic habit v2 (revised version)"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("Error update book record")
	// }

	// //Delete
	// book.ID = 1
	// err = db.Delete(&book).Error

	//route
	router := gin.Default()
	routerV1 := router.Group("/v1")
	routerV1.GET("/", handler.RootHandler)

	routerV1.GET("/hello", handler.HelloHandler)

	routerV1.GET("/books/:id/:title", handler.BooksHandler)
	routerV1.GET("/query", handler.QueryHandler)
	routerV1.POST("/books", handler.PostBooksHandler)
	routerV1.GET("/dump", dump.DumpDb)

	router.Run(":3030")

}
