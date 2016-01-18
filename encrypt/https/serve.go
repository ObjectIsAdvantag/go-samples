// Extract from feyeleanor talk at dotGo 2015 Paris
// http://fr.slideshare.net/feyeleanor/encrypt-all-transports?ref=http://www.thedotpost.com/2015/11/eleanor-mchugh-encrypt-all-transports
package main

import . "fmt"
import . "net/http"

const ADDRESS = ":443"

func main() {
	message := "hello world"
	HandleFunc("/hello", func(w ResponseWriter, r *Request) {
		w.Header().Set("Content-Type", "text/plain")
		Fprintf(w, message)
	})

	ListenAndServeTLS(ADDRESS, "cert.pem", "key.pem", nil)
}