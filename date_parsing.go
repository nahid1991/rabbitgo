package main

import (
	"fmt"
	"time"
)

func main() {
	var dateByte = []byte("2014-11-12T11:45:26.371Z")
	dateString := string(dateByte[:])
	fmt.Println(dateString)
	timeGo, err := time.Parse(time.RFC3339, dateString)
	if err!=nil {
		fmt.Println("Error while parsing date :", err)
	}
	fmt.Println(timeGo)
}