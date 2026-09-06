package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlePage)
	const port = "8010"
	srv := http.Server{
		Handler:      mux,
		Addr:         ":" + port,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	fmt.Println("Server starts on port:", port)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	const page = `<html>
<head></head>
<body>
	<p> Hello from Docker! I'm a Go server. </p>
</body>
</html>
`
	w.Write([]byte(page))
}
