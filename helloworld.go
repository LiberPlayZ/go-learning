package main

import (
	"fmt"
	"test_project/temp"

)

func myMessage() {
	fmt.Println("I just got executed!")
}

func myFunction(x int, y int) int {
	return x + y
}

func myFunction2(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

type Person struct {
	name string
	age int
	job string
	salary int
  }

var t int

func main() {
	var student string = "dada"
	var stduent2 = "dad"
	x := 2
	var a int
	var z, j int = 2, 6
	a = 7
	var arr1 = [3]int{1, 2, 3}
	var arr2 = [...]int{1, 2, 3} // infered length
	var arr3 = [3]int{}
	arr5 := [5]int{1: 10, 2: 40} // initial some specific index
	myslice := arr1[2:3]         // create slice from array using start and end
	fmt.Println(student)
	fmt.Println(x)
	fmt.Println(a)
	fmt.Println(stduent2)
	fmt.Println(t)
	fmt.Println(z, j)
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr5)
	fmt.Println(len(arr1))
	fmt.Println(myslice)

	myslice1 := make([]int, 5, 10) // initial slice with make function . if capacity not provide it equal to len .
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
	fmt.Println("Hello World!")

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15} // create slice with specfic values from array and resize capacity .
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)
	fmt.Println(neededNumbers)
	fmt.Println(cap(neededNumbers))
	fmt.Println(cap(numbersCopy))
	neededNumbers[0] = 5
	fmt.Println(numbers)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for idx, val := range arr5 { // for in range
		fmt.Printf("%v\t%v\n", idx, val)
	}

	for _, val := range arr1 { // only value
		fmt.Printf("%v\n", val)
	}

	myMessage()
	fmt.Println(myFunction(1, 2))
	a, b := myFunction2(5, "Hello")
	fmt.Println(a, b)

	var pers1 Person
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000
	fmt.Println(pers1)

	bs := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}
	fmt.Println("dadada",len(bs))
	fmt.Println(bs)
	var at = make(map[string]string) // The map is empty now
	at["brand"] = "Ford"
	at["model"] = "Mustang"
	at["year"] = "1964"
	fmt.Println(at)
	val1, ok1 := at["brand"] // Checking for existing key and its value and return his value
	fmt.Println(val1,ok1)
	delete(at,"brand")
	fmt.Println(at)
	val2, ok2 := at["brand"] // Checking for existing key and its value and return his value
	fmt.Println(val2,ok2)
	for k, v := range at {
		fmt.Printf("%v : %v, ", k, v)
	  }
	temp.Hello()
	fmt.Println(string(student[0]))


}
