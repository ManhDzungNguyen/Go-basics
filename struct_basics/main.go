package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	email     string
	phone     string
}

func (p person) getInfo(isPrint bool) (fullName string, contactInfo string) {
	fullName = p.firstName + " " + p.lastName
	contactInfo = fmt.Sprintf("Phone: %v\nEmail: %v", p.phone, p.email)
	if isPrint {
		fmt.Printf("Fullname: %v\n", fullName)
		fmt.Println(contactInfo)
	}
	return

}

func main() {
	p := person{
		firstName: "Manh Dung",
		lastName:  "Nguyen",
		email:     "kuuhaku.work@gmail.com",
		phone:     "+84 365336468",
	}
	_, _ = p.getInfo(true)

}
