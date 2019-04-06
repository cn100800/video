package cmd

import (
	"log"
	"net/http"
	"os"
	"time"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	video, err := os.Open("test.mp4")
	if err != nil {
		log.Fatal(err)
	}
	defer video.Close()

	http.ServeContent(w, r, "test.mp4", time.Now(), video)
}

func Video1() {
	http.HandleFunc("/", ServeHTTP)
	http.ListenAndServe(":8080", nil)
}
