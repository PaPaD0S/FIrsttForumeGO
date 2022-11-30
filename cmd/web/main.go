// main
package main

import (
	"fmt"

	"github/PaPa_D0S/FIrsttForumeGO/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", hanldlers.Home)
	http.HandleFunc("/about", hanldlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
