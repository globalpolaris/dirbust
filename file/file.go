package file

import (
	"bufio"
	"log"
	"os"
	"readfile/net"
	"sync"
)

func OpenFile(fileName, url string) {
	var wg sync.WaitGroup
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Failed opening file1: ", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		wg.Add(1)

		dir := scanner.Text()
		go net.CheckDir(dir, url, &wg)

	}
	wg.Wait()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
