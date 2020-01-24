package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
	"simple_test_project/app"
)

func main() {
	var c app.Config

	if err := envconfig.Process("", &c); err != nil {
		log.Fatal(err)
	}

	srv := fmt.Sprintf(":%s", c.Port)

	log.Fatal(http.ListenAndServe(srv, nil))

}
