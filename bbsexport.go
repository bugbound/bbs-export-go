package main

import (
	"fmt"
	"os"
    "net/http"
    "io/ioutil"
    "time"
    "encoding/json"
)


type PagedRecords struct {
    Num_results int  
    Page int
    Objects []DomainRecord 
}

type DomainRecord struct {
    Domain string
    Id int
}

func main() {
    if os.Args[1] == "domain" {
        //export_domains()
        foo1 := new(PagedRecords) // or &Foo{}
        getJson("http://bbsstore-service:7002/api/dns_store?page=10", foo1)
        //println(foo1.Objects[0].Domain)
        for i, s := range foo1.Objects {
            fmt.Println(i, s.domain)
        }
    }
}


var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
    r, err := myClient.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}

func export_domains() {
    url := "http://bbsstore-service:7002/api/dns_store?page=1"
    
    
    req, err := http.NewRequest("GET", url, nil)
    //req.Header.Set("Content-Type", "application/json")

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
