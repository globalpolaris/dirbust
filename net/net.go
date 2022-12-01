package net

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
)

type Web struct {
	statusCode int
	directory  string
}

var web Web
var client = &http.Client{}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func CheckDir(dir, url string, wg *sync.WaitGroup) {
	fullUrl := fmt.Sprintf("%s/%s", url, dir)
	req, err := http.NewRequest("GET", fullUrl, nil)
	check(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// log.Println(string(body))
	if resp.StatusCode != 404 {
		web.directory = dir
		web.statusCode = resp.StatusCode
		fmt.Printf("%s/%s\t\tStatus: %d\n", url, web.directory, web.statusCode)
	}
	wg.Done()

}

func NormalizeUrl(url string) string {
	validUrl := AddProtocol(url)
	if ValidateDomain(validUrl) != nil {
		panic("error URL is not valid")
	}
	return validUrl
}

func ValidateProtocol(url string) string {
	if strings.Contains(url, "http") {
		return "http"
	} else if strings.Contains(url, "https") {
		return "https"
	}
	return ""
}

func AddProtocol(url string) string {
	if ValidateProtocol(url) != "" {
		return url
	}
	return fmt.Sprintf("https://%s", url)
}

func GetDomain(url string) string {
	if ValidateProtocol(url) == "https" {
		return url[9:]
	} else if ValidateProtocol(url) == "http" {
		return url[8:]
	}
	return ""
}

func ValidateDomain(url string) error {
	domain := GetDomain(url)
	res, err := net.LookupHost(domain)
	check(err)
	if len(res) == 0 {
		return errors.New("invalid domain")
	}
	return nil
}
