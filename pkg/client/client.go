package client

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func MakeRequest(url string, i int, ch chan<- string) {
	start := time.Now()
	fmt.Printf("- calling %s\n", url)

	secs := time.Since(start).Seconds()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%d %.2f elapsed failed: %s", i, secs, url)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	ch <- fmt.Sprintf(
		"%d %.2f elapsed with response length: %d %s\n%s", i, secs, len(body),
		url, string(body),
	)
}

func Must[T any](r T, err error) T {
	Success(err)
	return r
}

func Success(err error) {
	if err != nil {
		panic(err)
	}
}

func Client() {
	runtime.GOMAXPROCS(1)

	if len(os.Args) > 3 && os.Args[1] == "user" {
		PostUserClient()
	}

	if len(os.Args) > 3 && os.Args[1] == "ws" {
		WebSocketClient()
	}

	start := time.Now()
	ch := make(chan string)
	url := os.Args[2]
	fmt.Printf("%s\n", url)
	n, _ := strconv.Atoi(os.Args[1])
	for i := 0; i < n; i++ {
		duration := rand.Intn(5)
		u := fmt.Sprintf("%s/%d/%d", strings.TrimSuffix(url, "/"), i, duration)
		go MakeRequest(u, i, ch)
	}
	failed := 0
	for i := 0; i < n; i++ {
		r := <-ch
		if strings.Contains(r, "failed") {
			failed += 1
		}
		fmt.Printf("- finished %s\n", r)
	}
	fmt.Printf("%.2fs elapsed, failed %d\n", time.Since(start).Seconds(), failed)
}
