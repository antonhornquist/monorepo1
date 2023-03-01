package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var (
		blobstorePort = flag.Int("port", 5000, "Blob store port.")
		filename      = flag.String("file", "", "File to upload.")
	)

	flag.Parse()

	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	res, err := http.Post(fmt.Sprintf("http://localhost:%d/", *blobstorePort), "binary/octet-stream", file)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	message, _ := ioutil.ReadAll(res.Body)
	fmt.Printf(string(message))
}
