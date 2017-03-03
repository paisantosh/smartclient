package main

import (
        "fmt"
	"net/http"
        "io/ioutil"
)

func main() {
        resp, _ := http.Get("http://localhost:8000/")
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)
        fmt.Println(body);
      
}
