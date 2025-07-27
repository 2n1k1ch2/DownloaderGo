package main

import (
	"DownloaderGo/internal/server"
	"fmt"
	"net/http"
	"os"
)

func main() {
	os.RemoveAll("files")
	serv := server.NewServer()
	wd, _ := os.Getwd()
	fmt.Println("WORKING DIR:", wd)
	http.ListenAndServe(":8080", serv)

}
