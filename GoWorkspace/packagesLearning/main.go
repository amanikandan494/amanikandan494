package main

import (
	"log"

	"github.com/amanikandan494/packagesLearning/helpers"
)

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "Ha stala vista"
	log.Println(myVar.TypeName)
}
