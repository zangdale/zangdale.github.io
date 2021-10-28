package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// var port = flag.Int("p", 80, "port")

func main() {

	// flag.Parse()
	port := 80

	dir, err := os.Getwd()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("http://127.0.0.1:%d/index.html \nDir: %s\n", port, dir)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(dir)))
	if err != nil {
		log.Panicln(err)
	}
}
