package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/begoon/wsgi-go/wsgi"
)

func PostUserClient() {
	name, host := os.Args[2], os.Args[3]
	url := strings.TrimSuffix(host, "/") + "/api/user"
	fmt.Printf("- post user %s to %s\n", name, url)

	request := wsgi.User{Name: name}
	fmt.Printf("- request [%v]\n", request)

	raw := Must(json.Marshal(request))
	fmt.Printf("- request [raw] %v\n", string(raw))

	r := Must(http.NewRequest("POST", url, bytes.NewBuffer(raw)))

	r.Header.Add("X-Api-Key", "secret")
	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp := Must(client.Do(r))
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	fmt.Println("headers", string(Must(json.MarshalIndent(resp.Header, "", "\t"))))

	body := Must(ioutil.ReadAll(resp.Body))
	fmt.Println("response:", string(body))
}
