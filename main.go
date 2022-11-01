package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arturgumerov/monapi/router"
)

func main() {
	fmt.Println("Start monAPI..")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 ..")
}
