package models

var Bookstore = make(map[string]*Book)

type Book struct {
	// The main identifier for the Book. This will be unique.
	ISDN   string `json:"isdn"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}



