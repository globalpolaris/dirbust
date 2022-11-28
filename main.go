package main

import (
	"flag"
	"fmt"
	"log"
	"readfile/file"
	"strings"
	"time"
)

func main() {

	start := time.Now()
	fileName := flag.String("file", "", "File to read")
	url := flag.String("url", "", "URL to check")
	flag.Parse()
	if *fileName == "" || strings.ToLower(*url) == "" {
		log.Fatal("Please specify file and URL!")
	}
	fmt.Printf("Host: %s\n", *url)
	fmt.Printf("Wordlist: %s\n", *fileName)
	file.OpenFile(*fileName, *url)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
