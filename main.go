package main

import (
	"gin-quick-start/router"
	"log"
)

func main() {

	r := router.InitRouter()

	err := r.Run(":2333")
	if err != nil {
		log.Fatal(err)
	}
}
