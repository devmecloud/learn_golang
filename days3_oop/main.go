package main

import ("fmt")

type Book struct {
	name string
	author string
	pages int
}

func (book Book) print_details() {
	fmt.Printf("Book %s was written by %s", book.name, book.author)
	fmt.Printf("\nIt contains %d pages\n", book.pages)
}


func main()  {
	book1 := Book{"Monster Blood", "Khai TK", 143}
	book1.print_details()

	book1.name = "Deploy React to S3 use AWS"
	book1.author = "DevMeCloud"

	book1.print_details()
}
