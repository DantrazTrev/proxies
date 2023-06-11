// Go tunnel server
package main
import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)


func main() {
	// Define the proxy server's address and port
	proxyHost := "localhost"
	proxyPort := "8000"

	// Create a reverse proxy handler
	 target, _ := url.Parse("http://localhost:443")

	proxy := httputil.NewSingleHostReverseProxy(target)

	// Create a custom handler that handles incoming requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Modify the request headers if necessary
		r.Header.Set("X-Custom-Header", "Value")

		// Serve the request through the reverse proxy
		proxy.ServeHTTP(w, r)
	})

	// Start the proxy server
	log.Printf("Proxy server is running on %s:%s", proxyHost, proxyPort)
	log.Fatal(http.ListenAndServe(proxyHost+":"+proxyPort, nil))
}





