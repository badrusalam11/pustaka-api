package book

type BookInput struct {
	Title    string `binding:"required"`
	Price    int    `json:"price" binding:"required,number"`
	SubTitle string `json:"sub_title", binding:"required"`
}
