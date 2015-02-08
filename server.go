package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {

	curDir, _ := os.Getwd()

	var port = flag.Int("port", 3030, "specify port to listen on.")
	var dir = flag.String("dir", curDir, "directory to serve. if omitted, current directory is used.")

	flag.Parse()

	fs := http.FileServer(http.Dir(*dir))

	fmt.Printf("Listening on :%d, serving %q\n", *port, *dir)

	http.Handle("/", fs)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
