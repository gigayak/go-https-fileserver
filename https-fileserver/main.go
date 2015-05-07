/* https-fileserver is an HTTPS webserver for static files.

Example usage:
  $ https-fileserver \
    --certificate=/opt/ssl/ssl.crt \
    --key=/opt/ssl/ssl.key \
    --dir=/opt/http
*/
package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	cert = flag.String("certificate", "", "Location of HTTPS .crt file.")
	key = flag.String("key", "", "Location of HTTPS .key private key file.")
	dir = flag.String("dir", "", "Directory to serve from.")
	addr = flag.String("address", ":443", "Address to listen on.")
)

func main() {
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*dir)))
	log.Printf("Listening on %s", *addr)
	err := http.ListenAndServeTLS(*addr, *cert, *key, nil)
	log.Fatal(err)
}
