package main

import "fmt"

func main() {
   	/* local variable definition */
   	var greeting = "Hi How are you"
	fmt.Printf("%s", greeting)
	fmt.Printf("\n jjjj")
	var n [10] int 	

	type books struct {
	title string
	author string
	book_id int }
	var book1 books
	book1.title = "Go programming"
	book1.author = "Sanjay"
	book1.book_id = 12345

	for i := 0; i<10; i++ {
	n[i] = i+10
	fmt.Printf("Test %d array value is %d\n",i, n[i])
		if (i == 7) {
		fmt.Printf("Test -----> %d array value is %d\n",i, n[i])
		}

	}
   fmt.Printf( "Book 1 title : %s\n", book1.title)
   fmt.Printf( "Book 1 author : %s\n", book1.author)
   fmt.Printf( "Book 1 book_id : %d\n", book1.book_id)
}
