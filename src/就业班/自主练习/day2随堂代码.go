package main

import "fmt"

type person struct {
	name   string
	gender bool
	age    int
	bag    []string
}

func main() {
	var p person
	setValue(&p)
	fmt.Println("p", p)
}
func setValue(p *person) {
	p.name = "harry"
	p.age = 22
	p.gender = false
	p.bag = append(p.bag, "phone", "book", "power bank")
}
