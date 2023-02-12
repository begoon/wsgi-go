package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func MakeRequest(url string, i int, ch chan<- string) {
	start := time.Now()
	fmt.Printf("- calling %s\n", url)

	resp, _ := http.Get(url)
	secs := time.Since(start).Seconds()
	body, _ := ioutil.ReadAll(resp.Body)
	ch <- fmt.Sprintf("%d %.2f elapsed with response length: %d %s\n%s", i, secs, len(body), url, string(body))
}

func main() {
	start := time.Now()
	ch := make(chan string)
	url := os.Args[2]
	fmt.Println(url, "\n")
	n, _ := strconv.Atoi(os.Args[1])
	for i := 0; i < n; i++ {
		duration := rand.Intn(5)
		u := fmt.Sprintf("%s/%d/%d", strings.TrimSuffix(url, "/"), i, duration)
		go MakeRequest(u, i, ch)
	}
	for i := 0; i < n; i++ {
		fmt.Printf("- finished %s\n", <-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
