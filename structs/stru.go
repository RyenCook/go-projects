package main

import "fmt"

type Book struct {
    title string
    author string
    genre string
    id int
}

func main() {
    var Book1 Book
    var Book2 Book

    Book1.title = "CS102 is Dumb"
    Book1.author = "Ryan Cook"
    Book1.genre = "Nonfiction"
    Book1.id = 420

    Book2.title = "CS102 is Great"
    Book2.author = "Lila Holt"
    Book2.genre = "Fiction"
    Book2.id = 69

    /* print Book1 info */
    fmt.Printf( "Book 1 title: %s\n", Book1.title)
    fmt.Printf( "Book 1 author: %s\n", Book1.author)
    fmt.Printf( "Book 1 genre: %s\n", Book1.genre)
    fmt.Printf( "Book 1 id: %d\n", Book1.id)

    /* print Book2 info */
    fmt.Printf( "Book 2 title: %s\n", Book2.title)
    fmt.Printf( "Book 2 author: %s\n", Book2.author)
    fmt.Printf( "Book 2 genre: %s\n", Book2.genre)
    fmt.Printf( "Book 2 id: %d\n", Book2.id)
}
