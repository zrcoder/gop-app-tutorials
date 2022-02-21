package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("a port needed")
		return
	}
	port := os.Args[1]
	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Printf("serving on [http://localhost:%s\n", port)
	fmt.Println(http.ListenAndServe(":"+port, nil))
}
