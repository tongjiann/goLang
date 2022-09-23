package main

import "fmt"

func main() {
	//stru()
	//struPoint()
	str()
}

func str() {
	s := "123"
	fmt.Println(s)
}

func struPoint() {
	var person = Person{1, "xiw", 25}
	var personP = &person
	fmt.Println(person)
	fmt.Println(personP.name)
}

type Person struct {
	id   int
	name string
	age  int
}

func stru() {
	var p = Person{1, "xiw", 25}
	var p2 Person
	p2.id = 2
	p2.name = "tong"
	p2.age = 24
	fmt.Println(p)
	fmt.Println(p2)
}

func arrPoint() {
	arr := [...]int{1, 2, 3, 4, 5}
	var arrP = &arr
	for i, p := range arrP {
		fmt.Println(i, p)
	}
}

func point() {
	var ip *int
	var i = 2113
	ip = &i
	fmt.Println(ip)
	fmt.Println(*ip)
}

func array() {
	var intArr = [10]int{1, 2, 3}
	fmt.Println(intArr)
}

func helloGo() {
	fmt.Println("hello, go...")

}
