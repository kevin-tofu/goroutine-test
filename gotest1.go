package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {

	const arg1 = ""
	fmt.Println(greeting("Hello World"))
	fmt.Println(greeting2("Hello", "World"))
	fmt.Println(convert_order("Hello", "World"))

	a, b := convert_order("Hello", "World")
	fmt.Println(a, b)

	var v_int int = 1
	var v_float64 float64 = 13131311.0
	var v_bool bool = false
	var v_string string = "news"
	fmt.Printf("%v %v %v %q\n", v_int, v_float64, v_bool, v_string)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
	for sum < 200 {
		sum += sum
	}
	fmt.Println(sum)

	fmt.Println("if_statement", IfStatement(v_string))

	fmt.Println("IsContained", IsContained(v_string))

	defer fmt.Println("test1 with defer")
	defer fmt.Println("test2 with defer")
	fmt.Println("test3 without defer")

	var p1, p2, p3 = testStruct()
	testArraySlice(p1, p2, p3)
}

func testStruct() (Person, Person, Person) {

	var mike Person
	fmt.Println(mike, "without initialization")
	mike.name = "Mike"
	mike.age = 23

	var bob = Person{"Bob", 35}

	var sam = Person{age: 89, name: "Sam"}

	fmt.Println(mike)
	fmt.Println(bob)
	fmt.Println(sam)

	return mike, bob, sam
}

func testArraySlice(a Person, b Person, c Person) {

	var myarr [3]Person = [3]Person{a, b, c}
	// var arr3 = [...]Person{a, b, c}

	var myslice []Person = []Person{a, b, c}

	fmt.Println(len(myslice)) //=> 2
	fmt.Println(cap(myslice)) //=> 2

	var PointerPerson *Person = &myarr[0]
	var myslice2 = make([]string, 2, 2)
	fmt.Println(PointerPerson)
	fmt.Println(myslice2)

	myslice = append(myslice, a)
	myslice = append(myslice, a)
	myslice = append(myslice, a)
	for index, value := range myslice[0:3] {
		fmt.Println(index, value.name, value.age)
	}

}

func testMap() {
	// map1 := make(map[string]string)
	map1 := make(map[string]string, 2)
	map1["Name"] = "KK"
	map1["Gender"] = "Male"

	map2 := map[string]int{"Age": 25, "UserId": 97}
	delete(map2, "Age")

	fmt.Println(map1)
	fmt.Println(map2)

}

func IsContained(category string) bool {
	switch category {
	case
		"auto",
		"news",
		"sport",
		"music":
		return true
	}
	return false
}

func IfStatement(arg string) string {
	if v := "sport"; arg == v {
		return "This is sport"
	} else {
		return "This is not sport"
	}
}

func greeting(arg string) string {
	return arg
}

func greeting2(arg1, arg2 string) string {
	return arg1 + " " + arg2
}

func convert_order(arg1, arg2 string) (string, string) {
	return arg2, arg1
}
