package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "net/http"
    "bytes"
    "io/ioutil"
)


func main() {
    if os.Args[1] == "domain" {
        export_domains()
    }
}

func export_domains() {
    url := "http://bbsstore-service:7002/api/dns_store"
    var pageNumber = 1
    var json = []byte(`{"page":"1"}`)
    
    
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    //fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
    
}
