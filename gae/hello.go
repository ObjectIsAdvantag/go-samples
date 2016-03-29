// Example of module deployable to google app engine, and importing a third party library (util.Reverse)
// Tested ok on 1/18/2016
package gae

import (
	"fmt"
	"net/http"

	"github.com/ObjectIsAdvantag/go-samples/util"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, util.Reverse("Hello Gopher !"))
//	fmt.Fprint(w, "Hello Gopher you're welcome here !")
}

