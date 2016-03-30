// Example of code deployable to google app engine, 
// and importing a third party library (util.Reverse)
package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/ObjectIsAdvantag/go-samples/util"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, util.Reverse("Hello Gopher !"))
	fmt.Fprint(w, ", running go version " + runtime.Version())
}

// not executed if running on google app engine
func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
