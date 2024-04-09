package main

import (
	"fmt"
	"hw1/internal/app"
	"log"
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = a.Run()
	if err != nil {
		fmt.Println(err)
		log.Fatalf(err.Error())
	}
}
