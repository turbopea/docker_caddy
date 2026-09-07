package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlePage)
	//err := godotenv.Load()
	//if err != nil {
	//log.Fatal("Error loading .env file")
	//}
	port := os.Getenv("PORT")
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
	<p> Hi Docker, I pushed a new version. </p>
</body>
</html>
`
	w.Write([]byte(page))
}
