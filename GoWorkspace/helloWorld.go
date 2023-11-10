package main

import "fmt"

func main() {
	fmt.Println("Hello world !!!!!")
	var i int
	i = 7
	fmt.Println("i is set to", i)
	whatWasSaid := saySomething()
	fmt.Println("said thing ==", whatWasSaid)
}
func saySomething() string {
	return "something"
}
