//page web welcome
package main
import (
	"fmt"
	"net/http"
)
func handlerW(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome in golang")
}
func main() {
	http.HandleFunc("/welcome", handlerW)
	http.ListenAndServe(":7777", nil)
}

