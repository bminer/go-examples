package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Birthday  time.Time
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}
func (p Person) String() string {
	return p.FullName() + " was born on " + p.Birthday.Format("Jan 2, 2006")
}
func main() {
	me := Person{"Blake", "Miner",
		time.Date(1988, time.June, 7, 12, 30, 0, 0, time.UTC)}
	fmt.Println(me.FullName())
	fmt.Println(me)
}
