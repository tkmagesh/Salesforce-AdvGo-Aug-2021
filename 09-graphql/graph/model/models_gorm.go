package model

type AuthorGorm struct {
	AuthorID  string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName";gorm:"column:author_id"`
}

type BookGorm struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"Author";gorm:"ForeignKey:AuthorID"`
}
