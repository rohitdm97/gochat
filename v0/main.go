package v0

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, time.Now().Format("2006-01-02 15:04:05"))
}

func Main() {
	http.HandleFunc("/", MainHandler)

	fmt.Println("Listening on port 5050...")

	_ = http.ListenAndServe(":5050", nil)
}
