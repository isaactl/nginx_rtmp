package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const ListenAdd = "localhost:8081"
const VideoStorePath = "/mnt/vod/"

func handleUpload(w http.ResponseWriter, r *http.Request) {
	log.Printf("got request: %s", r.URL.String())

	// parse args
	r.ParseForm()
	r.ParseMultipartForm(32 << 20)
	log.Printf("request args: %+v", r.Form)

	filePath := r.Form.Get("file.path")
	fileName := r.Form.Get("file.name")
	contentType := r.Form.Get("file.content_type")
	md5 := r.Form.Get("file.md5")
	size := r.Form.Get("file.size")
	log.Printf("file path: %s, name: %s, content type: %s, md5: %s, size: %s",
		filePath, fileName, contentType, md5, size)

	log.Print(os.Rename(filePath, fmt.Sprintf("%s%s", VideoStorePath, fileName)))
}

func main() {
	log.Printf("starting uploader server %s", ListenAdd)
	http.HandleFunc("/upload", handleUpload)

	log.Fatal(http.ListenAndServe(ListenAdd, nil))
}
