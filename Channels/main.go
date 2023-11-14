package main

import (
	"log"

	"github.com/amanikandan494/Channels/helpers"
)

const numPool = 30

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumbers(numPool)
	intChan <- randomNumber
}
func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)
	num := <-intChan
	log.Println(num)
}
