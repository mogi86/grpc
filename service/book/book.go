package book

import "regexp"

type Book struct {
	Id int32
	Title string
	Price int32
}

var books = []Book{
	{
		Id: 1,
		Title: "Golang入門",
		Price: 2000,
	},
	{
		Id: 2,
		Title: "PHP入門",
		Price: 2500,
	},
	{
		Id: 3,
		Title: "Python入門",
		Price: 2500,
	},
}

func Find(keyword string) Book {
	r := regexp.MustCompile(keyword)

	for _, v := range books {
		if r.MatchString(v.Title) == true {
			return v
		}
	}
	return Book{}
}

func (book *Book) isEmpty() bool {
	if book.Id == 0 {
		return true
	}
	return false
}
