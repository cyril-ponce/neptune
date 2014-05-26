package bkz

import (
	
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

)

type Book struct {

	Title string
	Author string
	ISBN string
	Genre string
	Id string

}

// Creates an account and adds it to the Database
func CreateBook(book *Book) bool {

	// Dial up a mongoDB session
	session, err := mgo.Dial("127.0.0.1:27017/")
	if err != nil {
		fmt.Println(err)
		return false
	}

	// Opens the "library" databases, "books" collection
	c := session.DB("library").C("books")
	result := Book{}
	
	// Search for the bookID, place in result.Id
	err = c.Find(bson.M{"id": book.Id}).One(&result)
	
	if result.Id != "" {
		// return true because book is present in the database
		// and we can say, "it's been added" without causing errors
		return true
	}

	// insert the book if it is not already in the database
	err = c.Insert(*book)

	if err != nil {
		return false
	}
	return true
}

// Finds book already in database
func FindBook(bookid string) (book *Book) {

	session, err := mgo.Dial("127.0.0.1:27017/")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	
	// Opens the "library" databases, "books" collection
	c := session.DB("library").C("books")
	
	// Finds the bookid and fills the book struct with the data
	err = c.Find(bson.M{"id": bookid}).One(&book)

	return book
}
