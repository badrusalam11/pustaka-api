package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/dump"
	"pustaka-api/handler"
	"pustaka-api/initializers"
	"pustaka-api/middleware"
	"pustaka-api/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	initializers.LoadEnvVariables()
	db, err = initializers.ConnectToDatabase()
	err = initializers.SyncDatabase(db)
	if err != nil {
		log.Fatal("Connection to database failed", err.Error())
	}
}

func main() {
	//book
	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)
	//user
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

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
	routerV1Books := routerV1.Group("/books", middleware.RequireAuth)
	routerV1Books.POST("", bookHandler.PostBooksHandler)
	routerV1Books.GET("", bookHandler.GetBooks)
	routerV1Books.GET("/:id", bookHandler.GetBook)
	routerV1Books.PUT("/:id", bookHandler.UpdateBookHandler)
	routerV1Books.DELETE("/:id", bookHandler.DeleteBook)

	routerV1.POST("/signup", userHandler.Signup)
	routerV1.POST("/login", userHandler.Login)
	routerV1.GET("/dump", dump.DumpDb)

	router.Run(":3030")

}
