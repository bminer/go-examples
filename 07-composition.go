package main

import "fmt"

type Author struct {
	FirstName string
	LastName  string
}

func (a Author) FullName() string {
	return a.FirstName + " " + a.LastName
}

type Post struct {
	Title   string
	Content string
	Author
}

func (p Post) String() string {
	return p.Title + "\nBy " + p.FullName() + "\n\n" + p.Content
}

func main() {
	me := Author{"Blake", "Miner"}
	p := Post{
		Title:   "Composition vs. Inheritance",
		Content: "Inheritance is really not what you want.",
		Author:  me,
	}
	fmt.Println(p)
}
