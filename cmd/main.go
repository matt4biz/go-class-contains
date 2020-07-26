package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"contains"
)

const (
	url   = "https://jsonplaceholder.typicode.com"
	known = `{"id": 1}`
)

func main() {
	id := "1"

	if len(os.Args) > 1 {
		id = os.Args[1]
	}

	resp, err := http.Get(url + "/todos/" + id)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	if err := contains.CheckData(known, body); err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
