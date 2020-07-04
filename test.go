package main

import "fmt"

type TestStr struct {
	id      int
	another int
	some    string
}

func (te *TestStr) Add() {
	te = TestStr{
		id:      10,
		another: 10,
		some:    "Whatever",
	}
}

func main() {
	test := new(TestStr)

	fmt.Println("Main")
	fmt.Println(test)
	fmt.Println(&test)
	fmt.Println(*test)

	testPass1(*test)
	fmt.Println("Main after test 1")
	fmt.Println(test)
	fmt.Println(&test)
	fmt.Println(*test)

	testPass2(test)
	fmt.Println("Main after test 2")
	fmt.Println(test)
	fmt.Println(&test)
	fmt.Println(*test)

}

func testPass1(te TestStr) {
	fmt.Println("Test pass 1")
	fmt.Println(te)
	fmt.Println(&te)
	te.Add()
}

func testPass2(te *TestStr) {
	fmt.Println("Test pass 2")
	fmt.Println(te)
	fmt.Println(&te)
	fmt.Println(*te)
	te.Add()
}
